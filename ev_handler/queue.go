package evhandler

import (
	"sync"
)

// Queue implementation
type Queue struct {
	items []uint32
	mu    sync.Mutex
}

func (q *Queue) Enqueue(item uint32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (uint32, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) Peek() (uint32, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}
