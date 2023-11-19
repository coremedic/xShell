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

/*
Get a shells entire log

Args -> shellID
*/
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

/*
Get C2 listener status
*/
func (l *Link) C2Status() (*protobuf.C2StatusResponse, error) {
	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Make gRPC request
	resp, err := l.Client.C2Status(ctx, &protobuf.C2StatusRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

/*
Generate new client (mTLS) certificate

Args -> Username
*/
func (l *Link) NewClient(username string) (*protobuf.NewClientResponse, error) {
	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Make gRPC request
	resp, err := l.Client.NewClient(ctx, &protobuf.NewClientRequest{Username: username})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, err
}

/*
Task shell to execute operation

Args -> shellID, Operation, Arguments
*/
func (l *Link) ExecuteOperation(shellID string, op string, args []string) error {
	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Make gRPC request
	_, err := l.Client.ExecuteOperation(ctx, &protobuf.ExecuteOperationRequest{
		Shell:     shellID,
		Operation: op,
		Arguments: args,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
