package teamserver_test

import (
	"crypto/rand"
	"net"
	"testing"
	"time"
	"xShell/controller/teamserver"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

var err error

func TestTeamServer(t *testing.T) {
	// Get TeamServer singleton instance
	ts := teamserver.GetTeamServerInstance()

	// Generate host keys
	_, _, ts.PrivKey, ts.PubKey, err = teamserver.GenHostKeys()
	if err != nil {
		t.Fatal("Failed to generate host keys:", err)
	}
	// Start the TeamServer
	ts.Start()
	// Wait for the TeamServer to start
	time.Sleep(2 * time.Second)
	select {}
	// Generate ed25519 key pair for the SSH client
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal("Failed to generate client keys:", err)
	}

	// Convert the ed25519 public key to SSH public key
	sshPubKey, err := ssh.NewPublicKey(pubKey)
	if err != nil {
		t.Fatal("Failed to create SSH public key:", err)
	}
	// Convert the SSH public key to authorized_keys format
	authKeyBytes := ssh.MarshalAuthorizedKey(sshPubKey)
	// Add clients public key to the authorized keys of the "test" user
	ts.AuthKeys["test"] = append(ts.AuthKeys["test"], authKeyBytes)
	// Parse the private key for the SSH client
	privKeySigner, err := ssh.NewSignerFromKey(privKey)
	if err != nil {
		t.Fatal("Failed to create SSH signer from private key:", err)
	}
	// Create SSH client config
	clientConfig := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privKeySigner),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Attempt to connect to the SSH server
	client, err := ssh.Dial("tcp", net.JoinHostPort("localhost", "2222"), clientConfig)
	if err != nil {
		t.Fatal("Failed to dial:", err)
	}
	// Open a new session channel
	channel, _, err := client.OpenChannel("session", []byte("session"))
	if err != nil {
		t.Fatal("Failed to open channel:", err)
	}
	// Send a simple command and close the channel
	channel.Write([]byte("clear"))
	channel.Close()

	// Give it a moment and then close the client
	time.Sleep(1 * time.Second)
	client.Close()
}
