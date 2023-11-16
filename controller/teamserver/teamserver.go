package teamserver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"xShell/controller/c2"
	"xShell/internal/logger"
	"xShell/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

/*
TeamServer struct

Port -> Listening port for TeamServer (default is 1991)

ServerCert -> TeamServer TLS cert for mTLS connections

CACert -> Certificate Authority cert

CAKey -> Certificate Authority key

Listener -> C2 listener object

protobuf -> Protobuf service struct
*/
type TeamServer struct {
	Port       string `default:"1991"`
	ServerCert *tls.Certificate
	CACert     []byte
	CAKey      []byte
	Listener   *c2.C2
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
				logger.Log(logger.PANIC, fmt.Sprintf("%v", r))
			}
		}()
		// Create new x509 cert pool
		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM(ts.CACert); !ok {
			logger.Log(logger.CRITICAL, "Failed to append client CA cert")
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
			// Add client logging and failed auth interceptors
			grpc.UnaryInterceptor(chainUnaryServerInterceptors(clientCertInterceptor, authErrorInterceptor)),
		)
		// Register gRPC service
		protobuf.RegisterControllerServiceServer(grpcServer, ts)
		// Listen for incoming connections
		listener, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", ts.Port))
		if err != nil {
			log.Panic(err)
		}
		// Pass listener to gRPC service
		if err := grpcServer.Serve(listener); err != nil {
			log.Panic(err)
		}
	}()
}

/*
gRPC server interceptor for logging client gRPC requests.
*/
func clientCertInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	methodName := info.FullMethod // Get the full RPC method name
	logger.Log(logger.DEBUG, fmt.Sprintf("Intercepting method: %s", methodName))

	peer, ok := peer.FromContext(ctx)
	if !ok || peer.AuthInfo == nil {
		logger.Log(logger.ERROR, "No peer information available")
		return handler(ctx, req)
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		logger.Log(logger.ERROR, "No TLS information available")
		return handler(ctx, req)
	}

	for _, cert := range tlsInfo.State.PeerCertificates {
		username := cert.Subject.CommonName

		if pb, ok := req.(proto.Message); ok {
			logger.Log(logger.DEBUG, fmt.Sprintf("Marshaling request for method %s", methodName))
			marshalledPb, err := proto.Marshal(pb)
			if err != nil {
				logger.Log(logger.ERROR, fmt.Sprintf("Failed to marshal proto message for method %s: %v", methodName, err))
				continue
			}
			if len(marshalledPb) == 0 {
				logger.Log(logger.AUDIT, fmt.Sprintf("%s called an empty request for method %s", username, methodName))
			} else {
				logger.Log(logger.AUDIT, fmt.Sprintf("%s called method %s with request: %v", username, methodName, marshalledPb))
			}
		} else {
			logger.Log(logger.AUDIT, fmt.Sprintf("%s called method %s but request could not be logged (not a proto message)", username, methodName))
		}
	}

	return handler(ctx, req)
}

/*
gRPC server interceptor for logging failed mTLS authentications
*/
func authErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		peer, ok := peer.FromContext(ctx)
		if ok && peer.AuthInfo != nil {
			if tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo); ok {
				for _, cert := range tlsInfo.State.PeerCertificates {
					username := cert.Subject.CommonName
					logger.Log(logger.WARNING, fmt.Sprintf("mTLS auth failed for user %s: %v", username, err))
				}
			}
		}
	}
	return resp, err
}

/*
Chain multiple interceptors into one handler
*/
func chainUnaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for i := len(interceptors) - 1; i >= 0; i-- {
			handler = createHandler(interceptors[i], handler)
		}
		return handler(ctx, req)
	}
}

/*
Create handler with interceptor
*/
func createHandler(interceptor grpc.UnaryServerInterceptor, handler grpc.UnaryHandler) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, &grpc.UnaryServerInfo{}, handler)
	}
}
