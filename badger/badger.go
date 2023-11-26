package main

import (
	"badger/evasion"
	"fmt"
)

//garble:controlflow
func main() {
	err := evasion.UnhookKernel32()
	if err != nil {
		fmt.Println(err)
	}
}
