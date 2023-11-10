package link

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"sync"
	"xShell/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
Link singleton

Conn -> gRPC connection

Host -> TeamServer host address

Port -> TeamServer port (defailt: 1991)

Cert -> Client certificate

Secure -> true: verify teamserver cert (recommended), false: disable all checks (not secure)
*/
type Link struct {
	Conn   *grpc.ClientConn
	Client protobuf.ControllerServiceClient
	Host   string
	Port   string `default:"1991"`
	Cert   []byte
	Secure bool
}

var (
	linkInstance *Link
	once         sync.Once
)

func GetLinkInstance() *Link {
	once.Do(func() {
		linkInstance = &Link{
			Port:   "1991",
			Secure: true,
		}
	})
	return linkInstance
}

/*
Starts connection to TeamServer
*/
func (l *Link) Connect() {
	// Decode all PEM blocks
	var clientCertPEM []byte
	var clientKeyPEM []byte
	var caCertPEM []byte

	for {
		block, rest := pem.Decode(l.Cert)
		if block == nil {
			break
		}
		switch block.Type {
		case "CERTIFICATE":
			// If we haven't set the client cert, set it, else append to CA pool
			if clientCertPEM == nil {
				clientCertPEM = pem.EncodeToMemory(block)
			} else {
				caCertPEM = append(caCertPEM, pem.EncodeToMemory(block)...)
			}
		case "PRIVATE KEY", "RSA PRIVATE KEY", "EC PRIVATE KEY":
			if clientKeyPEM == nil {
				clientKeyPEM = pem.EncodeToMemory(block)
			}
		default:
			log.Fatalf("unknown PEM block type: %s", block.Type)
		}
		l.Cert = rest
	}

	if clientCertPEM == nil || clientKeyPEM == nil {
		log.Fatalf("client certificate and/or private key not found")
	}

	// Create a certificate pool with the CA's certificates
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCertPEM) {
		log.Fatalf("failed to add CA certificates to pool")
	}

	// Load the client certificate and key
	cert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		log.Fatalf("failed to load client key pair: %v", err)
	}

	// Create the TLS credentials for transport
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	// Disable all verification checks
	if !l.Secure {
		config.InsecureSkipVerify = true
	}

	creds := credentials.NewTLS(config)

	// Dial the gRPC server
	l.Conn, err = grpc.Dial(net.JoinHostPort(l.Host, l.Port), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to teamserver: %v", err)
	}

	// Initialize the gRPC client
	l.Client = protobuf.NewControllerServiceClient(l.Conn)
}

/*
Close connection to TeamServer

Should be defered until program exit
*/
func (l *Link) Close() error {
	return l.Conn.Close()
}
