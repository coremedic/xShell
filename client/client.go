package client

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

//go:embed "config"
var config []byte

func main() {
	if args := os.Args[1]; args != "" {
		parts := strings.Split(args, "@")
		if len(parts) < 2 {
			fmt.Println("Invalid command")
			os.Exit(1)
		}
		user := parts[0]
		host := parts[1]
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password("password"),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		client, err := ssh.Dial("tcp", host+":1991", config)
		if err != nil {
			fmt.Println("Failed to dial: ", err)
			os.Exit(1)
		}
		defer client.Close()
	}
}
