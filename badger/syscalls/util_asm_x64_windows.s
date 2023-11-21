#include "textflag.h"

// func rVA2VA(moduleBase uintptr, rva uint32) uintptr
TEXT ·rVA2VA(SB), NOSPLIT, $0-24
    MOVQ moduleBase+0(FP), AX  // Load moduleBase into AX
    MOVL rva+8(FP), CX         // Load rva into CX (using CX instead of RDI)
    ADDQ CX, AX                // Add rva to moduleBase
    MOVQ AX, ret+16(FP)        // Store result in return value
    RET

// func ReadDwordAtOffset(start uintptr, offset uint32) DWORD
TEXT ·readDwordAtOffset(SB), NOSPLIT, $0-24
    MOVQ start+0(FP), AX       // Load start into AX
    MOVL offset+8(FP), CX      // Load offset into CX
    ADDQ CX, AX                // Add offset to start
    MOVL (AX), CX              // Read DWORD from the address in AX
    MOVL CX, ret+16(FP)        // Store the DWORD in return value
    RET

// func ReadWordAtOffset(start uintptr, offset uint32) WORD
TEXT ·readWordAtOffset(SB), NOSPLIT, $0-24
    MOVQ start+0(FP), AX       // Load start into AX
    MOVL offset+8(FP), CX      // Load offset into CX
    ADDQ CX, AX                // Add offset to start
    MOVW (AX), CX              // Read WORD from the address in AX
    MOVW CX, ret+16(FP)        // Store the WORD in return value
    RET

// func ReadByteAtOffset(start uintptr, offset uint32) uint8
TEXT ·readByteAtOffset(SB), NOSPLIT, $0-24
    MOVQ start+0(FP), AX       // Load start into AX
    MOVL offset+8(FP), CX      // Load offset into CX
    ADDQ CX, AX                // Add offset to start
    MOVB (AX), CL              // Read byte from the address in AX
    MOVB CL, ret+16(FP)        // Store the byte in return value
    RET
