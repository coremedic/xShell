package syscalls

/*
* Define Windows types
 */
type (
	// Win32 BOOL
	BOOL uint32
	// Win32 BOOLEAN
	BOOLEAN byte
	// Win32 BYTE
	BYTE byte
	// 32 bit DWORD
	DWORD uint32
	// 64 bit DWORD
	DWORD64 uint64
	// Win32 HANDLE
	HANDLE uintptr
	// Win32 LocalHandle
	HLOCAL uintptr
	// 64 bit integer
	LARGE_INTEGER int64
	// 32 bit integer
	LONG int32
	// LPVOID pointer
	LPVOID uintptr
	// Win32 SIZE_T
	SIZE_T uintptr
	// Unsigned 32 bit integer
	UINT uint32
	// Unsigned integer pointer
	ULONG_PTR uintptr
	// Unsigned 64 bit integer
	ULONGLONG uint64
	// 16 bit WORD
	WORD uint16
)

func NT_SUCCESS(x uint32) bool {
	return (x)&(1<<31) == 0
}

const (
	ImageNumberOfDirectoryEntries = 16
	ImageDirectoryEntryImport     = 1
	ImageDirectoryEntryBaseReloc  = 5
	ErrSyscallNotFound            = ((3 << 30) | (1 << 29) | 1) // Severity: STATUS_SEVERITY_ERROR, Custom: 1, Code:
)

type _ImageDataDirectory struct {
	VirtualAddress DWORD
	Size           DWORD
}

type ImageDataDirectory _ImageDataDirectory

/*
* IMAGE_FILE_HEADER structure (winnt.h)
* Represents the COFF header format
 */
type ImageFileHeader struct {
	Machine              WORD
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      WORD
}

/*
* IMAGE_NT_HEADERS64 structure (winnt.h)
* Represents the PE header format (64bit)
 */
type ImageNTHeaders64 struct {
	Signature      DWORD
	FileHeader     ImageFileHeader
	OptionalHeader ImageOptionalHeader64
}

/*
* IMAGE_DOS_HEADER structure (winnt.h)
* Header at start of PE file
 */
type ImageDosHeader struct {
	EMagic    WORD
	ECblp     WORD
	ECp       WORD
	ECrlc     WORD
	ECparhdr  WORD
	EMinalloc WORD
	EMaxalloc WORD
	ESs       WORD
	ESp       WORD
	ECsum     WORD
	EIp       WORD
	ECs       WORD
	ELfarlc   WORD
	EOvno     WORD
	ERes      [4]WORD
	EOemid    WORD
	EOeminfo  WORD
	ERes2     [10]WORD
	ELfanew   DWORD
}

/*
* IMAGE_EXPORT_DIRECTORY (winnt.h)
* Export directory of PE file
 */
type ImageExportDirectory struct {
	Characteristics       DWORD
	TimeDateStamp         DWORD
	MajorVersion          WORD
	MinorVersion          WORD
	Name                  DWORD
	Base                  DWORD
	NumberOfFunctions     DWORD
	NumberOfNames         DWORD
	AddressOfFunctions    DWORD
	AddressOfNames        DWORD
	AddressOfNameOrdinals DWORD
}

/*
* IMAGE_OPTIONAL_HEADER64 (winnt.h)
* 64 bit optional header
 */
type ImageOptionalHeader64 struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitalizedData        DWORD
	SizeOfUninitalizedData      DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	BaseOfData                  DWORD
	ImageBase                   DWORD
	SectionAllignment           DWORD
	FileAllignment              DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsytem                    WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          DWORD
	SizeOfStackCommit           DWORD
	SizeOfHeapReserve           DWORD
	SizeOfHeapCommit            DWORD
	LoaderFlag                  DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [ImageNumberOfDirectoryEntries]ImageDataDirectory
}

type ImageImportDescriptor struct {
	OriginalFirstThunk DWORD
	TimeDateStamp      DWORD
	ForwarderChain     DWORD
	Name               DWORD
	FirstThunk         DWORD
}

type ImageThunkData struct {
	AddressOfData uintptr
}

type OriginalImageThunkData struct {
	Ordinal uint
}
