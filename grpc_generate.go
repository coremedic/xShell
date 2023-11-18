package main

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
//go:generate tar -czvf controller/badger.tar.gz badger/
//go:generate chmod 744 ./grpc_generate.sh
//go:generate ./grpc_generate.sh
