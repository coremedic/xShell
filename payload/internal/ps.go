package internal

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Process struct {
	pid  uint32
	ppid uint32
	exe  string
	arch string
}

func GetPid(proc string) (uint32, error) {
	ps, err := processes()
	if err != nil {
		return 0, err
	}
	for _, p := range ps {
		if p.exe == proc {
			return p.pid, nil
		}
	}
	return 0, fmt.Errorf("no process found")
}

func newProc(entry *syscall.ProcessEntry32) *Process {
	term := 0
	for {
		if entry.ExeFile[term] == 0 {
			break
		}
		term++
	}
	return &Process{
		pid:  entry.ProcessID,
		ppid: entry.ParentProcessID,
		exe:  syscall.UTF16ToString(entry.ExeFile[:term]),
	}
}

func processes() ([]Process, error) {
	hProc, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(hProc)

	var procEntry32 syscall.ProcessEntry32
	procEntry32.Size = uint32(unsafe.Sizeof(procEntry32))
	if err = syscall.Process32First(hProc, &procEntry32); err != nil {
		return nil, err
	}

	procs := make([]Process, 0)
	for {
		procs = append(procs, *newProc(&procEntry32))

		if err := syscall.Process32Next(hProc, &procEntry32); err != nil {
			break
		}
	}
	return procs, nil
}
