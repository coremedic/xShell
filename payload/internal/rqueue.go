package internal

import (
	"net/http"
	"sync"
)

type SafeRequestQueue struct {
	mtx   sync.Mutex
	Queue []*http.Request
}

var RequestQueue SafeRequestQueue = SafeRequestQueue{Queue: make([]*http.Request, 0)}

func (q *SafeRequestQueue) Add(request *http.Request) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.Queue = append(q.Queue, request)
}

func (q *SafeRequestQueue) GetNext() *http.Request {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.Queue) <= 0 {
		return nil
	}
	return q.Queue[0]
}

func (q *SafeRequestQueue) ShiftUp() {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.Queue) > 0 {
		q.Queue = q.Queue[1:]
	}
}
