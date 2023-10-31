package internal

// import (
// 	bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"
// )

// type BP struct {
// 	Bp *bananaphone.BananaPhone
// }

// func (bp *BP) virtualAllocEx(hProcess uintptr, lpAddress uintptr, dwSize uint32, flAllocationType int, flProtect int) (uint32, error) {
// 	sysid, err := bp.Bp.GetSysID("VirtualAllocEx")
// 	if err != nil {
// 		return 0x00, err
// 	}
// 	ret, err := bananaphone.Syscall(
// 		sysid,
// 		hProcess,
// 		lpAddress,
// 		uintptr(dwSize),
// 		uintptr(flAllocationType),
// 		uintptr(flProtect),
// 	)
// 	if err != nil {
// 		return 0x00, err
// 	}
// 	return ret, nil
// }
