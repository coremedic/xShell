package console

import (
	"os/exec"
	"runtime"
)

type Command func([]string) error

var CommandMap = map[string]Command{
	"clear": clear,
}

func clear([]string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	return cmd.Run()
}
