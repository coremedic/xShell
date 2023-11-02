package main

import (
	"crypto/x509"
	"log"
	"os"
	"xShell/controller/teamserver"
)

const (
	certDir        = ".xshell/"
	caCertPath     = certDir + "ca-cert.pem"
	caKeyPath      = certDir + "ca-key.pem"
	serverCertPath = certDir + "teamserver-cert.pem"
	serverKeyPath  = certDir + "teamserver-key.pem"
)

func init() {
	// Check if .xshell directory exists, if not create it
	if _, err := os.Stat(certDir); os.IsNotExist(err) {
		os.Mkdir(certDir, 0700)
	}

	// Check if CA and server certs exists, if not create them
	if _, err := os.Stat(caCertPath); os.IsNotExist(err) {
		cert, key, err := teamserver.GenCACert()
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(caCertPath, cert, 0644)
		os.WriteFile(caKeyPath, key, 0644)
	} else if _, err := os.Stat(serverCertPath); os.IsNotExist(err) {
		caCertBytes, err := os.ReadFile(caCertPath)
		if err != nil {
			log.Fatal(err)
		}
		caCert, err := x509.ParseCertificate(caCertBytes)
		if err != nil {
			log.Fatal(err)
		}
		caKeyBytes, err := os.ReadFile(caKeyPath)
		if err != nil {
			log.Fatal(err)
		}
		caKey, err := x509.ParseECPrivateKey(caKeyBytes)
		if err != nil {
			log.Fatal(err)
		}
		serverCert, serverKey, err := teamserver.GenClientCert(caCert, caKey)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(serverCertPath, serverCert, 0644)
		os.WriteFile(serverKeyPath, serverKey, 0644)
	}
}

func main() {

}
