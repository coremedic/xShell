package evasion

import (
	"github.com/coremedic/goldr/pkg/syscalls"
)

var (
	Goldr *syscalls.IndirectSyscaller = &syscalls.IndirectSyscaller{}
)

func Debug() {
	syscalls.Debug()
}
