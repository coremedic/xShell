package main

import (
	"crypto/x509"
	"log"
	"os"
	"xShell/controller/teamserver"
	"xShell/internal/logger"
)

const (
	certDir        = ".xshell/"
	caCertPath     = certDir + "ca-cert.pem"
	caKeyPath      = certDir + "ca-key.pem"
	serverCertPath = certDir + "teamserver-cert.pem"
	serverKeyPath  = certDir + "teamserver-key.pem"
	rootCertPath   = certDir + "root-cert.pem"
)

func init() {
	// Check if .xshell directory exists, if not create it
	if _, err := os.Stat(certDir); os.IsNotExist(err) {
		os.Mkdir(certDir, 0700)
	}

	// Check if CA cert exists, if not create it
	if _, err := os.Stat(caCertPath); os.IsNotExist(err) {
		cert, key, err := teamserver.GenCACert()
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(caCertPath, cert, 0644)
		os.WriteFile(caKeyPath, key, 0644)
	}
	// Check if server cert exists, if not create it
	if _, err := os.Stat(serverCertPath); os.IsNotExist(err) {
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
		serverCert, serverKey, err := teamserver.GenTeamServerCert(caCert, caKey)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(serverCertPath, serverCert, 0644)
		os.WriteFile(serverKeyPath, serverKey, 0644)
	}
	// Check if root user cert exits, if not create it
	if _, err := os.Stat(rootCertPath); os.IsNotExist(err) {
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
		clientCert, err := teamserver.GenClientCert(caCert, caKey, "root")
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(rootCertPath, clientCert, 0644)
	}
}

func main() {
	// Create new logger
	logger.NewLogger(".xshell/controller.log")
	// Close at program exit
	defer logger.Close()
	// Start teamserver
	ts := teamserver.TeamServer{}
}
