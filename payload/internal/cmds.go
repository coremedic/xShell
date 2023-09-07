//go:build windows
// +build windows

package internal

import (
	"os/exec"

	"syscall"
)

type Op func([]string) ([]byte, error)

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
