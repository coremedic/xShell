package syscalls

// Assembly function stubs

// parser_asm_x64_windows.s
func getNtdllBase() uintptr

// parser_asm_x64_windows.s
func getKernel32Base() uintptr

// parser_asm_x64_windows.s
func getModuleExportsDirAddr(modAddr uintptr) uintptr

// parser_asm_x64_windows.s
func getExportsNumberOfNames(exportsAddr uintptr) uint32

// parser_asm_x64_windows.s
func getExportsAddressOfFunctions(modAddr uintptr, exportsAddr uintptr) uintptr

// parser_asm_x64_windows.s
func getExportsAddressOfNames(modAddr uintptr, exportsAddr uintptr) uintptr

// parser_asm_x64_windows.s
func getExportsAddressOfOrdinals(modAddr uintptr, exportsAddr uintptr) uintptr

// trampoline_asm_x64_windows.s
func getTrampoline(stubAddr uintptr) uintptr

// syscall_asm_x64_windows.s
func execIndirectSyscall(ssn uint16, trampoline uintptr, argh ...uintptr) uint32

// util_asm_x64_windows.s
func rVA2VA(moduleBase uintptr, rva uint32) uintptr

// util_asm_x64_windows.s
func readDwordAtOffset(start uintptr, offset uint32) DWORD

// util_asm_x64_windows.s
func readWordAtOffset(start uintptr, offset uint32) WORD

// util_asm_x64_windows.s
func readByteAtOffset(start uintptr, offset uint32) uint8
