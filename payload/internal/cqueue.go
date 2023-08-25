package internal

import "sync"

type SafeCommandQueue struct {
	mtx   sync.Mutex
	Queue []*string
}

var CommandQueue SafeCommandQueue = SafeCommandQueue{Queue: make([]*string, 0)}

func (q *SafeCommandQueue) Add(command *string) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.Queue = append(q.Queue, command)
}

func (q *SafeCommandQueue) GetNext() *string {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.Queue) <= 0 {
		return nil
	}
	ret := q.Queue[0]
	q.Queue = q.Queue[1:]
	return ret
}
