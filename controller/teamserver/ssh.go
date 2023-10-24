package teamserver

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type sshServer struct {
	Config   *ssh.ServerConfig
	Listener *net.Listener
	PubKey   []byte
	PrivKey  []byte
}

func (s *sshServer) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Print("Panic in (s *sshServer) Start(): Recovered")
				fmt.Printf("r: %v\n", r)
			}
		}()
		listener, err := net.Listen("tcp", net.JoinHostPort("localhost", "2222"))
		if err != nil {
			log.Panicf("Error starting listener: %s", err.Error())
		}
		s.Listener = &listener
		parsedKey, err := ssh.ParsePrivateKey(s.PrivKey)
		if err != nil {
			log.Panicf("Error parsing private key: %s", err.Error())
		}
		if s.Config == nil {
			s.Config = &ssh.ServerConfig{
				NoClientAuth: true,
			}
			s.Config.AddHostKey(parsedKey)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Panicf("Error accepting connection: %s", err.Error())
			}
			go clientHandler(conn, s.Config)
		}
	}()
}

func clientHandler(conn net.Conn, config *ssh.ServerConfig) {
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
		// WIP: Accept channel, handle commands
		channel, _, err := newChan.Accept()
		if err != nil {
			log.Panicf("Error accepting shh channel: %s", err.Error())
		}
		go newConsoleSession(channel)
	}
}
