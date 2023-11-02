package teamserver

import (
	"log"
	"xShell/protobuf"

	"google.golang.org/grpc"
)

type TeamServer struct {
	Port string
	protobuf.UnimplementedControllerServiceServer
}

func (ts *TeamServer) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%v", r)
				log.Println("Panic in TeamServer: Recovered")
			}
		}()
		// Create new gRPC server
		grpcServer := grpc.NewServer(grpc.Creds())
		protobuf.RegisterControllerServiceServer(grpcServer, ts)
	}()
}
