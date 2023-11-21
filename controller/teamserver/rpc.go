package teamserver

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"xShell/controller/c2"
	"xShell/controller/logger"
	"xShell/protobuf"
)

func (ts *TeamServer) ListShells(ctx context.Context, req *protobuf.ListShellsRequest) (*protobuf.ListShellsResponse, error) {
	var shells []*protobuf.ShellInfo

	return &protobuf.ListShellsResponse{Shells: shells}, nil
}

func (ts *TeamServer) NewClient(ctx context.Context, req *protobuf.NewClientRequest) (*protobuf.NewClientResponse, error) {
	certDERBlock, _ := pem.Decode(ts.CACert)
	if certDERBlock == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	caCert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, err
	}

	keyDERBlock, _ := pem.Decode(ts.CAKey)
	if keyDERBlock == nil {
		return nil, fmt.Errorf("failed to parse key PEM")
	}
	caKey, err := x509.ParseECPrivateKey(keyDERBlock.Bytes)
	if err != nil {
		return nil, err
	}

	cert, err := GenClientCert(caCert, caKey, req.Username)
	if err != nil {
		return nil, err
	}
	return &protobuf.NewClientResponse{Cert: cert}, nil
}

func (ts *TeamServer) StreamShellLog(req *protobuf.StreamShellLogRequest, stream protobuf.ControllerService_StreamShellLogServer) error {
	// Get shellID from request
	shellID := req.GetShell()
	logFile := filepath.Join(".xshell", "log", fmt.Sprintf("%s.log", shellID))
	// Ensure log file exists
	if _, err := os.Stat(logFile); err != nil {
		return err
	}
	// Open the log file
	log, err := os.Open(logFile)
	if err != nil {
		return err
	}
	// Defer the close until stream is closed
	defer log.Close()

	// Seek end of file, we only care about new entries
	if _, err := log.Seek(0, io.SeekEnd); err != nil {
		return err
	}

	// Init reader begin streaming loop
	reader := bufio.NewReader(log)
	for {
		select {
		case <-stream.Context().Done(): // Context cancelled, were done here
			return stream.Context().Err() // Return error
		default:
			line, err := reader.ReadString('\n') // Read until we encounter a new line
			if err != nil {
				if err == io.EOF {
					// Wait a second before we try to read again
					time.Sleep(time.Second)
					continue
				}
				// TODO: Add more checks for specific errors
				return err
			}
			// Send the log over the stream
			if err := stream.Send(&protobuf.StreamShellLogResposne{Log: line}); err != nil {
				return err
			}
		}
	}
}

/*
C2 status RPC
*/
func (ts *TeamServer) C2Status(ctx context.Context, req *protobuf.C2StatusRequest) (*protobuf.C2StatusResponse, error) {
	// Check if listener exists
	if ts.C2 == nil {
		logger.Log(logger.DEBUG, "Listener object is nil")
		return &protobuf.C2StatusResponse{
			Online:     false,
			Uptime:     0,
			ShellCount: 0,
		}, nil
	}
	// Make https client, set defaults
	client := http.DefaultClient
	// Ignore TLS certificate errors
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// Make https request to listener
	_, err := client.Get(fmt.Sprintf("https://127.0.0.1:%s/ping", ts.C2.Port))
	if err != nil {
		logger.Log(logger.DEBUG, err.Error())
		return &protobuf.C2StatusResponse{
			Online:     false,
			Uptime:     0,
			ShellCount: 0,
		}, nil
	}
	// Listener is online
	return &protobuf.C2StatusResponse{
		Online:     true,
		Uptime:     ts.C2.Uptime,
		ShellCount: int64(len(c2.ShellMap.Shells)),
	}, nil
}
