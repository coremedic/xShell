package internal

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func SpoofPPID(ppid uint32, cmd *exec.Cmd) error {
	pHandle, err := windows.OpenProcess(windows.PROCESS_CREATE_PROCESS|windows.PROCESS_CREATE_THREAD|windows.PROCESS_QUERY_INFORMATION, false, ppid)
	if err != nil {
		return err
	}
	cmd.SysProcAttr.ParentProcess = syscall.Handle(pHandle)
	return nil
}
