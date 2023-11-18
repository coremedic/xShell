package main

import (
	"crypto/tls"
	"crypto/x509"
	"embed"
	"encoding/pem"
	"log"
	"os"
	"xShell/controller/c2"
	"xShell/controller/logger"
	"xShell/controller/teamserver"
)

// Embedded badger implant
//
//go:embed badger.tar.gz
var badger embed.FS

const (
	confDir        = ".xshell/"
	logDir         = confDir + "log/"
	caCertPath     = confDir + "ca-cert.pem"
	caKeyPath      = confDir + "ca-key.pem"
	serverCertPath = confDir + "teamserver-cert.pem"
	serverKeyPath  = confDir + "teamserver-key.pem"
	rootCertPath   = confDir + "root-cert.pem"
)

var (
	ts       teamserver.TeamServer = teamserver.TeamServer{}
	listener c2.C2                 = c2.C2{}
)

func init() {
	// Create .xshell directory if it does not exist
	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		if err := os.Mkdir(confDir, 0700); err != nil {
			log.Fatalf("Failed to create directory %s: %v", confDir, err)
		}
	}

	// Create log directory if it does not exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0700); err != nil {
			log.Fatalf("Failed to create directory %s: %v", confDir, err)
		}
	}

	// Create CA certificate and key if they do not exist
	if _, err := os.Stat(caCertPath); os.IsNotExist(err) {
		cert, key, err := teamserver.GenCACert()
		if err != nil {
			log.Fatalf("Failed to generate CA cert: %v", err)
		}
		if err := os.WriteFile(caCertPath, cert, 0644); err != nil {
			log.Fatalf("Failed to write CA cert: %v", err)
		}
		if err := os.WriteFile(caKeyPath, key, 0644); err != nil {
			log.Fatalf("Failed to write CA key: %v", err)
		}
	}

	// Create server certificate and key if they do not exist
	if _, err := os.Stat(serverCertPath); os.IsNotExist(err) {
		caCertPEM, err := os.ReadFile(caCertPath)
		if err != nil {
			log.Fatalf("Failed to read CA cert: %v", err)
		}
		caKeyPEM, err := os.ReadFile(caKeyPath)
		if err != nil {
			log.Fatalf("Failed to read CA key: %v", err)
		}

		caCertDER, _ := pem.Decode(caCertPEM)
		if caCertDER == nil {
			log.Fatalf("Failed to decode CA cert PEM")
		}
		caKeyDER, _ := pem.Decode(caKeyPEM)
		if caKeyDER == nil {
			log.Fatalf("Failed to decode CA key PEM")
		}

		caCert, err := x509.ParseCertificate(caCertDER.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse CA cert: %v", err)
		}
		caKey, err := x509.ParseECPrivateKey(caKeyDER.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse CA key: %v", err)
		}

		serverCert, serverKey, err := teamserver.GenTeamServerCert(caCert, caKey)
		if err != nil {
			log.Fatalf("Failed to generate server cert: %v", err)
		}
		if err := os.WriteFile(serverCertPath, serverCert, 0644); err != nil {
			log.Fatalf("Failed to write server cert: %v", err)
		}
		if err := os.WriteFile(serverKeyPath, serverKey, 0644); err != nil {
			log.Fatalf("Failed to write server key: %v", err)
		}
	}

	// Create root user certificate if it does not exist
	if _, err := os.Stat(rootCertPath); os.IsNotExist(err) {
		caCertPEM, err := os.ReadFile(caCertPath)
		if err != nil {
			log.Fatalf("Failed to read CA cert: %v", err)
		}
		caKeyPEM, err := os.ReadFile(caKeyPath)
		if err != nil {
			log.Fatalf("Failed to read CA key: %v", err)
		}

		caCertDER, _ := pem.Decode(caCertPEM)
		if caCertDER == nil {
			log.Fatalf("Failed to decode CA cert PEM")
		}
		caKeyDER, _ := pem.Decode(caKeyPEM)
		if caKeyDER == nil {
			log.Fatalf("Failed to decode CA key PEM")
		}

		caCert, err := x509.ParseCertificate(caCertDER.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse CA cert: %v", err)
		}
		caKey, err := x509.ParseECPrivateKey(caKeyDER.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse CA key: %v", err)
		}

		clientCert, err := teamserver.GenClientCert(caCert, caKey, "root")
		if err != nil {
			log.Fatalf("Failed to generate client cert: %v", err)
		}
		if err := os.WriteFile(rootCertPath, clientCert, 0644); err != nil {
			log.Fatalf("Failed to write client cert: %v", err)
		}
	}

	caCertPEM, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}
	caKeyPEM, err := os.ReadFile(caKeyPath)
	if err != nil {
		log.Fatalf("Failed to read CA key: %v", err)
	}

	// Assign key and cert to the teamserver instance fields
	ts.CACert = caCertPEM
	ts.CAKey = caKeyPEM

	// Decode server certificate and key from PEM to DER for serverCert
	serverCertPEM, err := os.ReadFile(serverCertPath)
	if err != nil {
		log.Fatalf("Failed to read server cert: %v", err)
	}
	serverKeyPEM, err := os.ReadFile(serverKeyPath)
	if err != nil {
		log.Fatalf("Failed to read server key: %v", err)
	}

	serverCert, err := tls.X509KeyPair(serverCertPEM, serverKeyPEM)
	if err != nil {
		log.Fatalf("Failed to load server key pair: %v", err)
	}

	// Set the server certificate in the teamserver instance
	ts.ServerCert = &serverCert
}

func main() {
	// Create new logger
	logger.NewLogger(".xshell/log/controller.log")
	// Set log level
	if len(os.Args) >= 2 && (os.Args[1] == "--debug" || os.Args[1] == "-d") {
		logger.LogLevel = logger.DEBUG
	} else {
		logger.LogLevel = logger.WARNING
	}
	logger.LogLevel = logger.DEBUG
	// Close at program exit
	defer logger.Close()
	ts.Port = "1991"
	// Start TeamServer, runs as goroutine
	ts.Start()
	// Start C2 listener
	listener.CertFile = serverCertPath
	listener.KeyFile = serverKeyPath
	listener.Port = "1848"
	listener.Type = "https"
	// Add listener to TeamServer object
	ts.Listener = &listener
	listener.Start()
}
