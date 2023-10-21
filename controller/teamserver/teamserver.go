package teamserver

import "sync"

type TeamServer struct {
}

var (
	tsInstance *TeamServer
	once       sync.Once
)

func GetTeamServerInstance() *TeamServer {
	once.Do(func() {
		tsInstance = &TeamServer{}
	})
	return tsInstance
}

func (ts *TeamServer) Start() {

}
