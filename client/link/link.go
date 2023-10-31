package link

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type Link struct {
	SshChannel *ssh.Channel
	SshClient  *ssh.Client
	User       string
	Host       string
	Port       string
	UserKey    []byte
	HostKey    []byte
}

/*
Connects to teamserver

Starts connection with teamserver.
*/
func (l *Link) Connect() {
	go func() {
		// Parse user private key, get ssh.Signer
		signer, err := ssh.ParsePrivateKey(l.UserKey)
		if err != nil {
			log.Fatalf("[FATAL] %s\n", err.Error())
		}
		// Parse host public key
		parsedHostKey, err := ssh.ParsePublicKey(l.HostKey)
		if err != nil {
			log.Fatalf("[FATAL] %s\n", err.Error())
		}
		// Build Client Config
		clientConfig := &ssh.ClientConfig{
			User: l.User,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.FixedHostKey(parsedHostKey),
		}
		// Connect to teamserver via ssh
		l.SshClient, err = ssh.Dial("tcp", net.JoinHostPort(l.Host, l.Port), clientConfig)
		if err != nil {
			log.Fatalf("[FATAL] %s\n", err.Error())
		}

	}()
}
