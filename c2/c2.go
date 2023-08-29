package main

import (
	"fmt"
	"os"
	"xShell/c2/internal"
)

var Port string = "443"
var Key string = "thisismypassword"

func main() {
	if _, err := os.Stat("c2/data/cert.pem"); err != nil {
		fmt.Println("[*] Generating TLS Cert...")
		internal.GenerateCertificate(&internal.X509Cert{
			Orgs:     []string{""},
			Names:    []string{""},
			Ips:      []string{""},
			CertFile: "c2/data/cert.pem",
			KeyFile:  "c2/data/key.pem",
		})
	}
	fmt.Println("[*] Starting HTTPS Listener...")
	listener := internal.Listener{
		Port:     Port,
		CertFile: "c2/data/cert.pem",
		KeyFile:  "c2/data/key.pem",
		Key:      []byte(Key),
	}
	go listener.StartListener()
	internal.StartCLI()
}
