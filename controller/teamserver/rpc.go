package teamserver

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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
