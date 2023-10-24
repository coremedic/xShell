package teamserver_test

import (
	"net"
	"testing"
	"time"
	"xShell/controller/teamserver"

	"golang.org/x/crypto/ssh"
)

var err error

func TestSSHServer(t *testing.T) {
	// Get singleton
	ts := teamserver.GetTeamServerInstance()
	// Generate host keys
	ts.SshServer.PrivKey, ts.SshServer.PubKey, _, _, err = teamserver.GenHostKeys()
	if err != nil {
		t.Fatal(err.Error())
	}
	// Start ssh server
	ts.SshServer.Start()
	time.Sleep(1000)
	// Create ssh client config
	config := &ssh.ClientConfig{
		User:            "test",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to server
	client, err := ssh.Dial("tcp", net.JoinHostPort("localhost", "2222"), config)
	if err != nil {
		t.Fatal(err.Error())
	}
	sshchan, _, err := client.OpenChannel("session", []byte("session"))
	if err != nil {
		t.Fatal(err.Error())
	}
	sshchan.Write([]byte("clear"))
	sshchan.Close()
	time.Sleep(1000)
	// Close connection
	client.Close()
}
