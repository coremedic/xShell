linux_amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o controller-linux_amd64  controller/controller.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o client-linux_amd64 client/client.go
windows_amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o client-windows_amd64.exe client/client.go
darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o client-darwin_amd64 client/client.go
darwin_arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o client-darwin_arm64 client/client.go
all:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o controller-linux_amd64  controller/controller.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o client-linux_amd64 client/client.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o client-windows_amd64.exe client/client.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o client-darwin_amd64 client/client.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o client-darwin_arm64 client/client.go
clean:
	rm controller-linux_amd64 && rm client-linux_amd64 && rm client-windows_amd64.exe && rm client-darwin_amd64 && rm client-darwin_arm64