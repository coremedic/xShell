package console

import (
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

var helpMenu string = `
xShell v0.3.1 (2023-10-31)
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

func Start() {
	defer func() {
		if r := recover(); r != nil {
			log.Print("[ERR] Panic in console session")
		}
	}()
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
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "xShell > ",
		AutoComplete:    autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Panic(err.Error())
		return
	}
	defer rl.Close()
	rl.HistoryEnable()
	for {
		cmd, err := rl.Readline()
		if err != nil {
			log.Panic(err.Error())
			return
		}
		parts := strings.Fields(cmd)
		if c, e := CommandMap[parts[0]]; e {
			fmt.Println(c(parts))
		} else {
			fmt.Printf("Command '%s' not found", parts[0])
			continue
		}
	}
}
