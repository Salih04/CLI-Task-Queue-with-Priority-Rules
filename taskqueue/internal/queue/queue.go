package queue

import (
	"container/heap"

	"github.com/Salih04/CLI-Task-Queue-with-Priority-Rules/internal/task"
)

// Queue aslında içinde []Task yani Task listesi tutacak

// Queue = yapılacak görevlerin beklediği kutu
// items []Task = kutunun içindeki görevler

type Queue struct {
	items []*task.Task // Pointer to Task, from the task package
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Less(i, j int) bool {
	// Higher priority number = comes first
	// So we use > not
	return q.items[i].Priority > q.items[j].Priority
}

func (q *Queue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *Queue) Push(x any) {
	q.items = append(q.items, x.(*task.Task))
	// The x.(*task.Task) part is a type assertion — you're telling Go "I know this any is actually a *task.Task, unwrap it."
}

func (q *Queue) Pop() any {
	n := len(q.items)
	item := q.items[n-1]
	q.items = q.items[:n-1]
	return item
}

func New() *Queue {
	q := &Queue{}
	heap.Init(q)
	return q
}

func (q *Queue) AddTask(t *task.Task) {
	heap.Push(q, t)
}

func (q *Queue) Next() *task.Task {
	return heap.Pop(q).(*task.Task)
}
