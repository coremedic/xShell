package link

import (
	"context"
	"io"
	"log"
	"time"
	"xShell/protobuf"
)

/*
Execute ListShells gRPC request
*/
func (l *Link) ListShells() (*protobuf.ListShellsResponse, error) {
	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Make gRPC request
	resp, err := l.Client.ListShells(ctx, &protobuf.ListShellsRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func (l *Link) ShellLog(shellID string) (*protobuf.ShellLogResponse, error) {
	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Make gRPC request
	resp, err := l.Client.ShellLog(ctx, &protobuf.ShellLogRequest{Shell: shellID})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

/*
Begin streaming shell log

Args -> Context, shellID, log updates channel
*/
func (l *Link) StreamShellLog(ctx context.Context, shellID string, logUpdates chan<- string) {
	// Begin log stream
	stream, err := l.Client.StreamShellLog(ctx, &protobuf.StreamShellLogRequest{Shell: shellID})
	if err != nil {
		log.Println(err)
		close(logUpdates)
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF { // Server closed stream
			close(logUpdates)
			return
		}
		if err != nil { // Other error
			log.Println(err)
			close(logUpdates)
			return
		}
		select {
		case logUpdates <- resp.GetLog(): // New log
		case <-ctx.Done(): // We done here
			// Context is cancelled, exit the function
			return
		}
	}
}
