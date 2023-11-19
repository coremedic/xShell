package console

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"xShell/client/link"
	"xShell/protobuf"
)

type Command struct {
	Operation string
	Arguments []string
	Shell     string
}

/*
Legacy menu related code
*/

var legacyHelpMenu string = `
xShell v0.3.1 "Red October" (2023-10-31)
-------------------------------------------
Main Menu Commands:
Command                  Description
-------------------------------------------
shells                   List all active shells
shell <shell_name>       Interact with a specific shell
clear                    Clear the console
help                     Show this menu
return                   Exit shell interaction, return to main menu
status                   C2 listener status information
exit                     Exit xShell client
-------------------------------------------
Shell Interaction Commands:
Command                  Description
-------------------------------------------
exec <command>           Run command (defaults to powershell)
whoami                   Run "whoami" (powershell)
kill                     Kills payload process
-------------------------------------------
`

// Shell map for legacy menu
var legacyShellMap map[string]*protobuf.ShellInfo

// Clears console screen, only for legacy menu
func legacyClearConsole() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

/*
Fetch shell list (legacy console)
*/
func legacyFetchShells() error {
	// Fetch link instance
	linkInstance := link.GetLinkInstance()
	// Make gRPC call
	shells, err := linkInstance.ListShells()
	if err != nil {
		return err
	}
	// If no shells are active return error
	if len(shells.Shells) == 0 {
		return fmt.Errorf("No active shells")
	}
	// Init shell map
	legacyShellMap = make(map[string]*protobuf.ShellInfo, len(shells.Shells))
	// Loop though shells and add them to shell map
	for _, shell := range shells.Shells {
		legacyShellMap[shell.Id] = shell
	}
	return nil
}

/*
Fetch shells entire log (legacy console)
*/
func legacyFetchShellLog(shellID string) ([]byte, error) {
	// Fetch link instance
	linkInstance := link.GetLinkInstance()
	// Make gRPC call
	log, err := linkInstance.ShellLog(shellID)
	if err != nil {
		return nil, err
	}
	return log.ShellLog, nil
}

/*
Fetch C2 listener status (legacy console)
*/
func legacyFetchStatus() (map[string]string, error) {
	// Fetch link instance
	linkInstance := link.GetLinkInstance()
	// Make gRPC call
	resp, err := linkInstance.C2Status()
	if err != nil {
		return nil, err
	}
	// Make listener status map
	status := make(map[string]string)
	status["Online"] = strconv.FormatBool(resp.Online)
	status["Uptime"] = fmt.Sprint(resp.Uptime)
	status["Shell Count"] = fmt.Sprint(resp.ShellCount)
	return status, nil
}

/*
Generate new client certificate (legacy console)
*/
func legacyNewClient(username string) error {
	// Fetch link instance
	linkInstance := link.GetLinkInstance()
	// Make gRPC call
	fmt.Printf("Generating mTLS certificate for '%s'...\n", username)
	resp, err := linkInstance.NewClient(username)
	if err != nil {
		return err
	}
	// Write client cert to disk
	cert := resp.GetCert()
	path := filepath.Join(".", fmt.Sprintf("%s.pem", username))
	err = os.WriteFile(path, cert, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Certificate saved to '%s'\n", path)
	return nil
}
