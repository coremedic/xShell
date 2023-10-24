package teamserver

import (
	"os/exec"
	"runtime"
)

type Command func([]string) string

var CommandMap = map[string]Command{
	"clear": clear,
}

func clear([]string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Run()
	return ""
}
