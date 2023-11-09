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
*/
type Link struct {
	Conn   *grpc.ClientConn
	Client protobuf.ControllerServiceClient
	Host   string
	Port   string `default:"1991"`
	Cert   []byte
}

var (
	linkInstance *Link
	once         sync.Once
)

func GetLinkInstance() *Link {
	once.Do(func() {
		linkInstance = &Link{}
	})
	return linkInstance
}

/*
Starts connection to TeamServer
*/
func (l *Link) Connect() {
	// Split the full chain into its components (assuming PEM encoded)
	block, rest := pem.Decode(l.Cert)
	if block == nil {
		log.Fatalf("failed to parse certificate PEM")
	}
	caCertPEM := rest // Rest contains the CA certificate and possibly more certificates

	clientCertPEM := pem.EncodeToMemory(block) // Client cert is the first block
	clientKeyPEM, _ := pem.Decode(rest)        // Assume client key follows the client cert
	if clientKeyPEM == nil {
		log.Fatalf("failed to parse private key PEM")
	}

	// Create a certificate pool with the CA's certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCertPEM) {
		log.Fatalf("failed to add CA certificate to pool")
	}

	// Load the client certificate and key
	cert, err := tls.X509KeyPair(clientCertPEM, pem.EncodeToMemory(clientKeyPEM))
	if err != nil {
		log.Fatalf("failed to load client key pair: %v", err)
	}

	// Create the TLS credentials for transport
	config := &tls.Config{
		Certificates: []tls.Certificate{cert}, // Add client certificate
		RootCAs:      certPool,
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
