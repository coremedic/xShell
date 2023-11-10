package c2

import (
	"fmt"
	"sync"
)

/*
Shell struct

Each shell is an instace of the
implant calling back to the C2.
*/
type Shell struct {
	// Id -> Human friendly ID
	Id string
	// Ip -> IP address of shell
	Ip string
	// LastCall -> Time represented as 64bit integer
	LastCall int64
	// Tasks -> Slice of Tasks queued for Shell
	Tasks []*Task
}

/*
Task struct

Operation -> Task operation (exec, whoami, ps, etc)

Arguments -> Task operation arguments
*/
type Task struct {
	// Operation -> Task operation (exec, whoami, ps, etc)
	Operation string
	// Arguments -> Task operation arguments
	Arguments []string
}

/*
SafeShellMap

Mutex protected map of Shells
*/
type SafeShellMap struct {
	mtx sync.Mutex
	// Shells -> key: Shell Id, value: Pointer to Shell
	Shells map[string]*Shell
}

// Init global SafeShellMap
var ShellMap SafeShellMap = SafeShellMap{Shells: make(map[string]*Shell)}

/*
Add new Shell to ShellMap
*/
func (s *SafeShellMap) Add(shell *Shell) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.Shells[shell.Id]; !ok {
		s.Shells[shell.Id] = shell
	}
}

/*
Get Shell from map by Shell Id

Return -> Pointer to shell, error
*/
func (s *SafeShellMap) Get(shellId string) (*Shell, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if shell, exists := s.Shells[shellId]; exists {
		return shell, nil
	}
	return nil, fmt.Errorf("shell '%s' doesnt exist", shellId)
}

/*
Delete Shell from map by Shell Id
*/
func (s *SafeShellMap) Delete(shellId string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.Shells, shellId)
}

/*
Add Task to Shell
*/
func (s *SafeShellMap) AddTask(shellId string, task Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if shell, exists := s.Shells[shellId]; exists {
		shell.Tasks = append(shell.Tasks, &task)
	}
}

/*
Clear Task queue for shell
*/
func (s *SafeShellMap) ClearTasks(shellId string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if shell, exists := s.Shells[shellId]; exists {
		shell.Tasks = make([]*Task, 0)
	}
}
