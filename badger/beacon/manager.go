package beacon

import (
	"net/http"
	"sync"
)

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
Mutex protected queue of Tasks
*/
type safeTaskQueue struct {
	mtx   sync.Mutex
	Tasks []*Task
}

// Global Task queue
var TaskQueue safeTaskQueue = safeTaskQueue{Tasks: make([]*Task, 0)}

/*
Check if Task queue is empty
*/
func (tq *safeTaskQueue) IsEmpty() bool {
	tq.mtx.Lock()
	defer tq.mtx.Unlock()
	return len(tq.Tasks) == 0
}

/*
Enqueue Task
*/
func (tq *safeTaskQueue) Enqueue(task *Task) {
	tq.mtx.Lock()
	defer tq.mtx.Unlock()
	tq.Tasks = append(tq.Tasks, task)
}

/*
Return and dequeue top Task from queue
*/
func (tq *safeTaskQueue) Dequeue() *Task {
	// Check if queue is empty before locking mutex
	if tq.IsEmpty() {
		return nil
	}
	tq.mtx.Lock()
	defer tq.mtx.Unlock()
	task := tq.Tasks[0]
	// Write null value to top of queue
	tq.Tasks[0] = nil
	// Dequeue
	tq.Tasks = tq.Tasks[1:]
	return task
}

/*
Mutex protected request queue
*/
type safeRequestQueue struct {
	mtx   sync.Mutex
	Queue []*http.Request
}

// Global Request queue
var RequestQueue safeRequestQueue = safeRequestQueue{Queue: make([]*http.Request, 0)}

/*
Check if Request queue is empty
*/
func (rq *safeRequestQueue) IsEmpty() bool {
	rq.mtx.Lock()
	defer rq.mtx.Unlock()
	return len(rq.Queue) == 0
}

/*
Enqueue Request
*/
func (rq *safeRequestQueue) Enqueue(request *http.Request) {
	rq.mtx.Lock()
	defer rq.mtx.Unlock()
	rq.Queue = append(rq.Queue, request)
}

/*
Return and dequeue top Request from queue
*/
func (rq *safeRequestQueue) Dequeue() *http.Request {
	// Check if queue is empty before locking mutex
	if rq.IsEmpty() {
		return nil
	}
	rq.mtx.Lock()
	defer rq.mtx.Unlock()
	request := rq.Queue[0]
	// Write null value to top of queue
	rq.Queue[0] = nil
	// Dequeue
	rq.Queue = rq.Queue[1:]
	return request
}
