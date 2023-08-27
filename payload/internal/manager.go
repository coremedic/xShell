package internal

import (
	"container/heap"
	"fmt"
	"time"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var Workers map[int]*Worker = make(map[int]*Worker)
var OpenIds IntHeap

func AddWorker(w *Worker) int {
	var index int
	if len(OpenIds) > 0 {
		index = heap.Pop(&OpenIds).(int)
	} else {
		index = len(Workers)
	}
	Workers[index] = w
	go w.Run()
	return index
}

func StopWorker(index int) {
	if worker, exists := Workers[index]; exists {
		worker.Status = "stopping"
	}
	delete(Workers, index)
	heap.Push(&OpenIds, index)
}

func GetWorkerstatus(index int) (string, error) {
	worker, exists := Workers[index]
	if !exists {
		return "", fmt.Errorf("Worker not found at index %d", index)
	}
	return worker.Status, nil
}

func GetWorkerTime(index int) (time.Time, error) {
	worker, exists := Workers[index]
	if !exists {
		return time.Time{}, fmt.Errorf("Worker not found at index %d", index)
	}
	return worker.Time, nil
}

func GetAllStatus() map[int]string {
	statuses := make(map[int]string)
	for index, worker := range Workers {
		statuses[index] = worker.Status
	}
	return statuses
}
