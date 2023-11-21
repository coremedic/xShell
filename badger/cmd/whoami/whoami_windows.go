package whoami

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
Emulate "whoami" command on windows by getting host name and users home directory name
*/
//garble:controlflow flatten_passes=1 junk_jumps=2 block_splits=0
func Whoami() (string, error) {
	// Fetch hostname
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	// fetch home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// Extract the base of the path, should be username
	user := filepath.Base(homeDir)
	return fmt.Sprintf("%s\\%s", hostname, user), nil
}
