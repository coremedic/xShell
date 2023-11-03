package teamserver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"xShell/internal/logger"
	"xShell/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

/*
TeamServer struct

Port -> Listening port for TeamServer (default is 1991)

ServerCert -> TeamServer TLS cert for mTLS connections

CACert -> Certificate Authority cert

CAKey -> Certificate Authority key

protobuf -> Protobuf service struct
*/
type TeamServer struct {
	Port       string `default:"1991"`
	ServerCert *tls.Certificate
	CACert     []byte
	CAKey      []byte
	protobuf.UnimplementedControllerServiceServer
}

/*
Starts TeamServer and gRPC service.

Runs as goroutine and listens for incoming gRPC connections.
Clients authenicate and connect via mTLS.
*/
func (ts *TeamServer) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%v", r)
				log.Println("Panic in TeamServer: Recovered")
			}
		}()
		// Create new x509 cert pool
		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM(ts.CACert); !ok {
			log.Fatal("Failed to append client CA cert")
		}
		// Create new TLS creds for mTLS auth
		creds := credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{*ts.ServerCert},
			ClientCAs:    certPool,
		})
		// Create new gRPC server
		grpcServer := grpc.NewServer(
			grpc.Creds(creds),
			// Add logging middleware
			grpc.UnaryInterceptor(clientCertInterceptor),
		)
		// Register gRPC service
		protobuf.RegisterControllerServiceServer(grpcServer, ts)
		// Listen for incoming connections
		listener, err := net.Listen("tcp", ts.Port)
		if err != nil {
			log.Fatal(err)
		}
		// Pass listener to gRPC service
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
}

/*
Middleware for logging client gRPC requests.
*/
func clientCertInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	peer, ok := peer.FromContext(ctx)
	if ok && peer.AuthInfo != nil {
		if tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo); ok {
			for _, cert := range tlsInfo.State.PeerCertificates {
				username := cert.Subject.CommonName
				logger.Log(logger.AUDIT, fmt.Sprintf("%s called: %v", username, req))
			}
		}
	}
	return handler(ctx, req)
}
