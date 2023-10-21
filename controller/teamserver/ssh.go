package teamserver

import (
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
				log.Print("[ERR] Failed to start ssh server")
			}
		}()
		for {
			listener := *s.Listener
			conn, err := listener.Accept()
			if err != nil {
				log.Panic(err)
			}
			go clientHandler(conn, s.Config)
		}
	}()
}

func clientHandler(conn net.Conn, config *ssh.ServerConfig) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("[ERR] Failed to handle client")
		}
	}()
	_, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Panic(err.Error())
	}
	go ssh.DiscardRequests(reqs)

	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		// WIP: Accept channel, handle commands
	}
}
