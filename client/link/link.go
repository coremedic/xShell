package link

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"sync"

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
	Conn *grpc.ClientConn
	Host string
	Port string `default:"1991"`
	Cert []byte
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
	// Decode PEM encoded certificate and key
	block, _ := pem.Decode(l.Cert)
	// Parse x506 cerrificate
	parsedCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse key file: %v", err)
	}

	// Set up credentials for the gRPC connection
	certPool := x509.NewCertPool()
	certPool.AddCert(parsedCert)
	creds := credentials.NewClientTLSFromCert(certPool, "")

	// Set up gRPC connection with mTLS credentials
	l.Conn, err = grpc.Dial(net.JoinHostPort(l.Host, l.Port), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to teamserver: %v", err)
	}
}

/*
Close connection to TeamServer

Should be defered until program exit
*/
func (l *Link) Close() error {
	return l.Conn.Close()
}
