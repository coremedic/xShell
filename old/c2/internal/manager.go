package internal

import (
	"fmt"
	"sync"
	"time"
)

type Shell struct {
	Id    string
	Ip    string
	LCall time.Time
	Cmds  []string
}

type SafeShellMap struct {
	mtx    sync.Mutex
	Shells map[string]*Shell
}

var ShellMap SafeShellMap = SafeShellMap{Shells: make(map[string]*Shell)}

func (s *SafeShellMap) Add(shell *Shell) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, exists := s.Shells[shell.Id]; !exists {
		s.Shells[shell.Id] = shell
	}
}

func (s *SafeShellMap) Get(shellId string) (*Shell, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if shell, exists := s.Shells[shellId]; exists {
		return shell, nil
	}
	return nil, fmt.Errorf("shell '%s' doesnt exist", shellId)
}

func (s *SafeShellMap) GetAll() ([]*Shell, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	var shellA []*Shell
	for _, shell := range s.Shells {
		shellA = append(shellA, shell)
	}
	if shellA != nil {
		return shellA, nil
	}
	return nil, fmt.Errorf("no shells in map")
}

func (s *SafeShellMap) ClearCmds(shellId string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if shell, exists := s.Shells[shellId]; exists {
		shell.Cmds = nil
		return
	}
}
