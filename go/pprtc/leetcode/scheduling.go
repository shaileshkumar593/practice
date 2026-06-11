/*
Scheduling System (Task Scheduler)
🧠 Idea
Execute tasks based on priority (earliest time / highest priority)
*/
package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	name     string
	priority int
}

type TaskHeap []*Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(*Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	task := old[n-1]
	*h = old[:n-1]
	return task
}

func main() {
	tasks := []*Task{
		{"Task A", 3},
		{"Task B", 1},
		{"Task C", 2},
	}

	h := &TaskHeap{}
	heap.Init(h)

	for _, t := range tasks {
		heap.Push(h, t)
	}

	fmt.Println("Task execution order:")
	for h.Len() > 0 {
		task := heap.Pop(h).(*Task)
		fmt.Println(task.name)
	}
}