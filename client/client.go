package main

import (
	"xShell/client/console"
)

var err error

func main() {
	// // Check if proper flags have been passed
	// if len(os.Args) < 3 {
	// 	log.Fatalf("Usage: './%s [teamserver_ip] [path_to_client_cert]'", os.Args[0])
	// }
	// // Fetch Link singleton instance
	// linkInstance := link.GetLinkInstance()
	// // Read the PEM encoded certificate and key from os.Args[2]
	// linkInstance.Cert, err = os.ReadFile(os.Args[2])
	// if err != nil {
	// 	log.Fatalf("Failed to read the certificate file: %v", err)
	// }
	// // Set Link host address, fetch from first arguement
	// linkInstance.Host = os.Args[1]
	console.Start()
	select {}

}
