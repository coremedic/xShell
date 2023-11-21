package main

import (
	"badger/cmd/whoami"
	"badger/evasion"
	"badger/syscalls"
	"fmt"
)

//garble:controlflow
func main() {
	syscalls.Debug()
	err := evasion.UnhookKernel32()
	if err != nil {
		fmt.Println(err)
	}
	syscalls.Debug()
	i, err := whoami.Whoami()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
