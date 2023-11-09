package console

import (
	"context"
	"fmt"
	"os"
	"strings"
	"xShell/client/link"

	"github.com/chzyer/readline"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ConsoleUI struct {
	app          *tview.Application
	shellList    *tview.List
	logView      map[string]*tview.TextView
	statusView   *tview.TextView
	commandInput *tview.InputField
	cIO          consoleIO
}

type consoleIO struct {
	in  chan Command
	out chan map[string][]byte
	ctx context.Context
}

func NewConsoleUI() *ConsoleUI {
	cui := &ConsoleUI{
		app:          tview.NewApplication(),
		shellList:    tview.NewList(),
		logView:      make(map[string]*tview.TextView),
		statusView:   tview.NewTextView(),
		commandInput: tview.NewInputField(),
		cIO: consoleIO{
			in:  make(chan Command, 1),
			out: make(chan map[string][]byte, 1),
			ctx: context.Background(),
		},
	}
	cui.initializeUIComponents()
	return cui
}

func (cui *ConsoleUI) initializeUIComponents() {
	cui.shellList.SetTitle("Shells").SetBorder(true)
	cui.shellList.AddItem("Teamserver", "", 0, nil)

	cui.logView["Teamserver"] = tview.NewTextView()
	cui.logView["Teamserver"].SetDynamicColors(true)
	cui.logView["Teamserver"].SetRegions(true)
	cui.logView["Teamserver"].SetWrap(true)
	cui.logView["Teamserver"].SetScrollable(true)
	cui.logView["Teamserver"].SetTitle("Teamserver Log").SetBorder(true)

	cui.statusView.SetTitle("Status").SetBorder(true)
	cui.statusView.SetText("C2 Status: Offline\nTeamserver Status: Disconnected")

	cui.commandInput.SetLabel("Command: ")
	cui.commandInput.SetFieldBackgroundColor(tcell.ColorDefault)
	cui.commandInput.SetFieldTextColor(tcell.ColorGreen)
	cui.commandInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			cui.processInput(cui.commandInput)
		}
	})

	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(cui.logView["Teamserver"], 0, 1, false).
		AddItem(cui.commandInput, 3, 1, true)

	sideBarFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(cui.shellList, 0, 1, true).
		AddItem(cui.statusView, 3, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(sideBarFlex, 30, 1, false).
		AddItem(mainFlex, 0, 3, true)

	cui.app.SetRoot(mainLayout, true).SetFocus(cui.commandInput)
}

func (cui *ConsoleUI) AddNewShell(shellID string) {
	cui.app.QueueUpdateDraw(func() {
		if _, exists := cui.logView[shellID]; !exists {
			newLogView := tview.NewTextView()
			newLogView.SetDynamicColors(true)
			newLogView.SetRegions(true)
			newLogView.SetWrap(true)
			newLogView.SetScrollable(true)
			newLogView.SetBorder(true)
			newLogView.SetTitle(shellID + " Log")

			cui.logView[shellID] = newLogView
			cui.shellList.AddItem(shellID, "", 0, nil)
		}
	})
}

func (cui *ConsoleUI) processInput(input *tview.InputField) {
	command := strings.Fields(input.GetText())
	if len(command) == 0 {
		return // No command entered
	}
	// Process command

	input.SetText("")                  // Clear the input after processing
	cui.app.SetFocus(cui.commandInput) // Return focus to the command input
}

func (cui *ConsoleUI) Start() {
	if err := cui.app.Run(); err != nil {
		panic(err)
	}
}

/*
legacy readline based CLI ported from xShell v0.2

This was annoying to port over tbh.
All of this needs to be depricated once the new UI is functional
*/
func StartLegacyUI() {
	// Current shell we are interacting with
	var currentShell string
	// Create new readline autocompleter
	autoCompleter := readline.NewPrefixCompleter(
		readline.PcItem("shells"),
		readline.PcItem("shell"),
		readline.PcItem("clear"),
		readline.PcItem("help"),
		readline.PcItem("exit"),
		readline.PcItem("whoami"),
		readline.PcItem("kill"),
		readline.PcItem("return"),
	)
	// Create readline instance
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "xShell > ",
		AutoComplete:    autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "quit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	// Channel to receive shell log updates
	logUpdates := make(chan string)
	// Context cancel function to stop log stream
	var cancelStream context.CancelFunc
	// Flag to control the stream gorutine
	steaming := false
	for {
		// Check for shell log updates
		if currentShell != "" && steaming {
			select {
			case update, ok := <-logUpdates: // We received a log update
				if !ok {
					// Some sort of error occurred and the channel closed, stop streaming
					steaming = false
					logUpdates = make(chan string)
					continue
				}
				// Print the update
				fmt.Println(update)
				// Redraw the prompt
				l.Refresh()
			default: // We dont care about this
			}
		}
		// Set/Reset prompt based on current shell
		if currentShell != "" {
			l.SetPrompt(fmt.Sprintf("xShell %s >", currentShell))
		} else {
			l.SetPrompt("xShell > ")
		}
		// Read command
		command, err := l.Readline()
		if err != nil {
			break
		}
		// Remove trailing and leading white space
		command = strings.TrimSpace(command)
		// Split command into operation and arguments
		parts := strings.Split(command, " ")
		op := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}
		// Switch on operation
		switch op {
		case "exit": // Exit xShell client
			os.Exit(0)
		case "return": // Reset current shell, stop log stream, return to main menu
			if steaming && cancelStream != nil {
				cancelStream()
				steaming = false
			}
			currentShell = ""
			l.SetPrompt("xShell > ")
			continue
		case "help": // Display legacy help menu
			fmt.Print(legacyHelpMenu)
			continue
		case "clear": // Clear console
			err := legacyClearConsole()
			if err != nil {
				fmt.Println(err)
			}
			continue
		case "shells": // List active shells
			err := legacyFetchShells()
			if err != nil {
				fmt.Println(err)
				continue
			}
			for shellId, shellInfo := range legacyShellMap {
				fmt.Printf("ID: %s, IP: %s Last Call: %.0d seconds ago\n", shellId, shellInfo.Ip, shellInfo.LastCall)
			}
			continue
		case "shell": // Start shell interaction
			err := legacyFetchShells()
			if err != nil {
				fmt.Println(err)
				continue
			}
			// Check if arg is passed
			if len(args) < 1 {
				fmt.Println("Usage: shell <shell_ID>")
				continue
			}
			// Shell ID should be passed as first argument
			if shell, ok := legacyShellMap[args[0]]; ok {
				currentShell = shell.Id
			} else {
				fmt.Printf("Shell '%s' not found\n", args[0])
				continue
			}
			// Check for existing stream, if stream exists cancel it
			if steaming && cancelStream != nil {
				cancelStream()
			}
			// Fetch full shell log
			log, err := legacyFetchShellLog(args[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			// Begin streaming log
			var ctx context.Context
			ctx, cancelStream = context.WithCancel(context.Background())
			steaming = true
			// This will never actually occur, but to avoid a context leak, defer the close
			defer cancelStream()
			// We need to get link instance manually here
			linkInstance := link.GetLinkInstance()
			go linkInstance.StreamShellLog(ctx, args[0], logUpdates)
			// Print shell log to console
			fmt.Println(string(log))
		}
	}
}
