package teamserver

import (
	"os"
	"sync"
)

type TeamServer struct {
	SshServer *sshServer
	Port      string
}

var (
	sshInstance *sshServer
	tsInstance  *TeamServer
	once        sync.Once
)

func init() {
	if _, err := os.Stat("./.xshell"); err != nil {
		os.Mkdir("./xhsell", 0744)
	}
}

func GetTeamServerInstance() *TeamServer {
	once.Do(func() {
		sshInstance = &sshServer{}
		tsInstance = &TeamServer{
			SshServer: sshInstance,
		}
	})
	return tsInstance
}

func (ts *TeamServer) Start() {

}
