package cmd

import "container/heap"

/*
Worker heap data structure

Heap for managing Workers
*/
type WorkerHeap []int

/*
Basic heap operations

These are value receiver methods since we are not making changed to the underlying data
*/
func (h WorkerHeap) Len() int           { return len(h) }
func (h WorkerHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h WorkerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

/*
Push interface to heap
*/
func (h *WorkerHeap) Push(x interface{}) {
	// Append to heap
	*h = append(*h, x.(int))
}

/*
Pop heap
*/
func (h *WorkerHeap) Pop() interface{} {
	// Old heap instance and length
	old := *h
	n := len(old)
	// Get top element from heap
	x := old[n-1]
	// Slice off top element
	*h = old[0 : n-1]
	// Return element
	return x
}

// Global Worker pool
var (
	WorkerPool map[int]*Worker = make(map[int]*Worker)
	FreeIds    WorkerHeap
)

/*
Add Worker to pool

Returns -> Worker Id in heap
*/
func AddWorker(w *Worker) int {
	var i int
	// If our heap is empty we dont need to pop
	if len(FreeIds) <= 0 {
		i = len(WorkerPool)
	} else { // Heap is not empty, we need to pop
		i = heap.Pop(&FreeIds).(int)
	}
	// Worker to pool
	WorkerPool[i] = w
	// TODO: Add logic to run Worker?
	// go w.Start()
	return i
}
