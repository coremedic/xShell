package teamserver

import (
	"context"
	"xShell/protobuf"
)

func (ts *TeamServer) ListShells(ctx context.Context, req *protobuf.ListShellsRequest) (*protobuf.ListShellsResponse, error) {
	var shells []*protobuf.ShellInfo

	return &protobuf.ListShellsResponse{Shells: shells}, nil
}
