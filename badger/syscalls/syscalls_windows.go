package syscalls

import (
	"fmt"
	"sort"
)

func init() {
	getCleanTrampolines(ntSyscalls)
}

var (
	cleanTrampolines []uintptr
	ntdllBase        uintptr    = getNtdllBase()
	ntSyscalls       []*Syscall = parseNtSyscalls()
)

// readCStringAt reads a null-terminated ANSI string from memory.
func readCStringAt(start uintptr, offset uint32) []byte {
	var buf []byte
	for {
		ch := readByteAtOffset(start, offset)
		if ch == 0 {
			break
		}
		buf = append(buf, ch)
		offset++
	}
	return buf
}

/*
Get clean (unhooked) trampolines from ntdll
*/
func getCleanTrampolines(syscalls []*Syscall) {
	// sort syscalls by RVA
	sort.Slice(syscalls, func(i, j int) bool {
		return syscalls[i].RVA < syscalls[j].RVA
	})

	// find clean trampolines
	for _, st := range syscalls {
		if trampoline := getTrampoline(st.VA); trampoline != uintptr(0) {
			st.TrampolinePtr = trampoline
			cleanTrampolines = append(cleanTrampolines, trampoline)
		}
	}

	// get SSNs
	for i, st := range syscalls {
		st.SSN = uint16(i)

		if st.TrampolinePtr == uintptr(0) {
			syscalls[i].TrampolinePtr = cleanTrampolines[0]
		}
	}
}

/*
Syscall structure
*/
type Syscall struct {
	// Name of syscall i.e. "NtVirtualProtect"
	Name string
	// Relative Virtual Address of syscall
	RVA DWORD
	// Virtual Address of syscall (pointer)
	VA uintptr
	// System Service Number (Syscall ID)
	SSN uint16
	// Pointer to clean trampoline
	TrampolinePtr uintptr
}

// adapted from github.com/f1zm0/acheron/
// Parses syscalls in Ntdll
func parseNtSyscalls() []*Syscall {
	modExpDirAddr := getModuleExportsDirAddr(ntdllBase)
	expNumNames := getExportsNumberOfNames(modExpDirAddr)
	expAddrNames := getExportsAddressOfNames(ntdllBase, modExpDirAddr)
	expAddrFunc := getExportsAddressOfFunctions(ntdllBase, modExpDirAddr)
	expAddrOrd := getExportsAddressOfOrdinals(ntdllBase, modExpDirAddr)

	syscallStubs := make([]*Syscall, 0)
	for i := uint32(0); i < expNumNames; i++ {
		fn := readCStringAt(ntdllBase, uint32(readDwordAtOffset(expAddrNames, i*4)))
		if fn[0] == 'Z' && fn[1] == 'w' {
			fn[0] = 'N'
			fn[1] = 't'
			nameOrd := readWordAtOffset(expAddrOrd, i*2)
			rva := readDwordAtOffset(expAddrFunc, uint32(nameOrd*4))

			syscallStubs = append(syscallStubs, &Syscall{
				Name: string(fn),
				RVA:  rva,
				VA:   rVA2VA(ntdllBase, uint32(rva)),
			})
		}
	}
	return syscallStubs
}

/*
IndirectSyscaller struct object
*/
type IndirectSyscaller struct{}

/*
Executre indirect syscall with clean trampoline
*/
func (i IndirectSyscaller) Syscall(fnName string, args ...uintptr) (uint32, error) {
	var syscall *Syscall
	// Loop through found syscalls to find match
	for _, sc := range ntSyscalls {
		if sc.Name == fnName {
			syscall = sc
			break
		}
	}
	// Syscall not found
	if syscall.Name == "" {
		return 1, fmt.Errorf("failed to find syscall")
	}
	// Execute indirect syscall
	ret := execIndirectSyscall(syscall.SSN, syscall.TrampolinePtr, args...)
	// If NT_SUCCESS is false
	if !NT_SUCCESS(ret) {
		return ret, fmt.Errorf("failed with code: 0x%x", ret)
	}
	return ret, nil
}

/*
Debug indirect syscalls
*/
func Debug() {
	modExpDirAddr := getModuleExportsDirAddr(ntdllBase)
	expNumNames := getExportsNumberOfNames(modExpDirAddr)
	expAddrNames := getExportsAddressOfNames(ntdllBase, modExpDirAddr)
	expAddrFunc := getExportsAddressOfFunctions(ntdllBase, modExpDirAddr)
	expAddrOrd := getExportsAddressOfOrdinals(ntdllBase, modExpDirAddr)

	fmt.Printf("Ntdll: 0x%x\nModuleExportsDirAddr: 0x%x\nExportsNumberOfNames: %d\nExportsAddressOfNames: 0x%x\nExportsAddressOfFunctions: 0x%x\nExportsAddressOfOrdinals: 0x%x\n", ntdllBase, modExpDirAddr, expNumNames, expAddrNames, expAddrFunc, expAddrOrd)

	found := parseNtSyscalls()
	getCleanTrampolines(found)
	for _, sc := range found {
		fmt.Printf("Found syscall '%s'\n\tSSN: %d\n\tAddr: 0x%x\n\tTrampoline: 0x%x\n", sc.Name, sc.SSN, sc.VA, sc.TrampolinePtr)
	}
}
