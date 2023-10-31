package teamserver

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"sync"

	"golang.org/x/crypto/ssh"
)

/*
TeamServer singleton
*/
type TeamServer struct {
	SshConfig *ssh.ServerConfig   // ssh server config
	PubKey    []byte              // Public host key
	PrivKey   []byte              // Private host key
	AuthKeys  map[string][][]byte // Authorized keys map key: username, value: array of public keys
	Port      string              // TeamServer listener port
}

var (
	tsInstance *TeamServer
	once       sync.Once
)

// func init() {
// 	// Check if .xshell directory exists if not create it
// 	if _, err := os.Stat(".xshell"); err != nil {
// 		os.Mkdir(".xshell", 0744)
// 	}
// }

/*
Gets singleton instace of TeamServer

Return -> Pointer to TeamServer instance
*/
func GetTeamServerInstance() *TeamServer {
	once.Do(func() {
		tsInstance = &TeamServer{
			AuthKeys: make(map[string][][]byte),
			Port:     "1991",
		}
	})
	return tsInstance
}

/*
Starts TeamServer

Both the public and private host key fields
must be set. TeamServer.Start() will start a
TCP listener and listen for incoming ssh connections.
If TeamServer.SshConfig is not set, a default
config will be used.
*/
func (ts *TeamServer) Start() {
	// Run as goroutine
	go func() {
		// Panic recover function
		defer func() {
			if r := recover(); r != nil {
				// Log recovery
				log.Print("Panic in (s *sshServer) Start(): Recovered")
				log.Printf("r: %v\n", r)
			}
		}()
		// Start tcp listener
		listener, err := net.Listen("tcp", net.JoinHostPort("localhost", "2222"))
		if err != nil {
			// Panic
			log.Panicf("Error starting listener: %s", err.Error())
		}
		parsedKey, err := ssh.ParsePrivateKey(ts.PrivKey)
		if err != nil {
			// Panic
			log.Panicf("Error parsing private key: %s", err.Error())
		}
		// If SshConfig is not set
		if ts.SshConfig == nil {
			// Set default config
			ts.SshConfig = &ssh.ServerConfig{
				//PublicKeyCallback: ts.clientAuth,
				NoClientAuth: true,
			}
		}
		// Add host key to config
		ts.SshConfig.AddHostKey(parsedKey)
		// Await incoming connections
		for {
			// Accept connection
			conn, err := listener.Accept()
			if err != nil {
				log.Panicf("Error accepting connection: %s", err.Error())
			}
			// Start new clientHandler goroutine
			go ts.clientHandler(conn, ts.SshConfig)
		}
	}()
}

/*
TeamServer client handler

Handles post-auth client connections.
Accepts channels from client.
Channel is then passed to a new console goroutine.
*/
func (ts *TeamServer) clientHandler(conn net.Conn, config *ssh.ServerConfig) {
	// Panic recover function
	defer func() {
		if r := recover(); r != nil {
			log.Print("Panic in clientHandler(): Recovered")
		}
	}()
	_, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Panicf("Error accepting ssh connection: %s", err.Error())
	}
	go ssh.DiscardRequests(reqs)

	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChan.Accept()
		if err != nil {
			log.Panicf("Error accepting shh channel: %s", err.Error())
		}

		go func(in <-chan *ssh.Request) {
			for req := range in {
				log.Printf("Received SSH request of type: %s", req.Type)
				if req.Type == "pty-req" {
					log.Printf("Handling pty-req.")
					termLen := req.Payload[3]
					modeList := req.Payload[4+termLen:]
					modeList[1] = 1 // ECHO mode. Set to 1
					req.Reply(true, nil)
				}
			}
		}(requests)

		//		go consoleSession(channel)
	}
}

/*
TeamServer client authenticator

Authentications users based on username and public key
*/
func (ts *TeamServer) clientAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	// Marshal public key
	keyBytes := ssh.MarshalAuthorizedKey(key)
	userKeys, ok := ts.AuthKeys[conn.User()]
	if !ok {
		return nil, fmt.Errorf("authentication failed")
	}
	for _, userKey := range userKeys {
		if bytes.Equal(keyBytes, userKey) {
			return nil, nil
		}
	}
	return nil, fmt.Errorf("authentication failed")
}
