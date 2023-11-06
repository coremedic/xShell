package console

import (
	"context"

	"github.com/rivo/tview"
)

var (
	app       *tview.Application
	shellList *tview.List
	logView   *tview.TextView
	c2Conf    *tview.Form
	cmdIn     *tview.Form
	cIO       consoleIO
)

/*
consoleIO

Handles IO to/from console UI
*/
type consoleIO struct {
	in        chan map[string][]byte
	out       chan string
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (io *consoleIO) GetIn() chan map[string][]byte {
	return io.in
}

func (io *consoleIO) GetOut() chan string {
	return io.out
}

func (io *consoleIO) GetCtx() context.Context {
	return io.ctx
}

/*
Refresh the UI
*/
func updater() {

}

/*
CreateListenerForm

Creates a new form for adding C2 listeners
*/
func CreateListenerForm() *tview.Form {
	// Form for creating a new listener
	listenerForm := tview.NewForm().
		AddInputField("Listener Name", "", 20, nil, nil).
		AddInputField("Host", "0.0.0.0", 20, nil, nil).
		AddInputField("Port", "443", 20, nil, nil).
		AddButton("Create", func() {
			// Logic to create listener
		}).
		AddButton("Cancel", func() {
			app.SetFocus(shellList) // Return focus to shell list or another appropriate component
		})
	listenerForm.SetBorder(true).SetTitle("Create Listener").SetTitleAlign(tview.AlignLeft)
	return listenerForm
}

/*
EditListenerForm

Creates a form to edit existing C2 listener configurations
*/
// func EditListenerForm(listenerName string) *tview.Form {
// 	// Assume getListenerConfig returns the current configuration for the listener
// 	ip, port := getListenerConfig(listenerName)

// 	listenerForm := tview.NewForm().
// 		AddInputField("IP", ip, 20, nil, nil).
// 		AddInputField("Port", port, 20, nil, nil).
// 		AddButton("Save", func() {
// 			// Logic to save listener configuration
// 		}).
// 		AddButton("Cancel", func() {
// 			app.SetFocus(shellList) // Return focus to shell list or another appropriate component
// 		})
// 	listenerForm.SetBorder(true).SetTitle("Edit Listener Config").SetTitleAlign(tview.AlignLeft)
// 	return listenerForm
// }

/*
Start console UI
*/
func Start() {
	go func() {
		// Create new tview application
		app = tview.NewApplication()
		// Create new tview list for active shells
		shellList = tview.NewList()
		shellList.SetTitle("Shells").SetBorder(true)
		// Create new tview textview for shell log
		logView = tview.NewTextView()
		logView.SetTitle("Log").SetBorder(true)
		// Create new tview form for command input
		cmdIn = tview.NewForm()
		cmdIn.SetTitle("Command").SetBorder(true)
		// Create channels and context for console IO
		cIO.in = make(chan map[string][]byte, 1)
		cIO.out = make(chan string, 1)
		cIO.ctx, cIO.ctxCancel = context.WithCancel(context.Background())
		// Add an input field for cmdIn
		cmdIn.AddInputField("", "", 20, nil, func(text string) {
			cIO.out <- text
		})
		// Create C2 config UI
		c2Conf = tview.NewForm()
		c2Conf.SetTitle("C2 Config").SetBorder(true)
		c2Conf.
			AddInputField("Host", "0.0.0.0", 0, nil, nil).
			AddInputField("Port", "443", 0, nil, nil)
		grid := tview.NewGrid().
			SetRows(-5, -1, -2).
			SetColumns(0, -3).
			SetBorders(false)
		grid.AddItem(shellList, 0, 0, 1, 1, 0, 100, false).
			AddItem(c2Conf, 1, 0, 3, 1, 0, 100, true).
			AddItem(logView, 0, 1, 2, 1, 0, 100, false).
			AddItem(cmdIn, 2, 1, 1, 1, 0, 100, false)

		if err := app.SetRoot(grid, true).SetFocus(grid).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	}()
}
