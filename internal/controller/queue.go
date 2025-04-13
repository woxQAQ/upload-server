package controller

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type DtQueue[T any] struct {
	queue       chan T
	checkMemory bool
	maxBytes    int64
	curBytes    int64
	mu          sync.Mutex
}

func NewDtQueue[T any](capacity int, maxBytes int64) *DtQueue[T] {
	return &DtQueue[T]{
		queue:       make(chan T, capacity),
		maxBytes:    maxBytes,
		checkMemory: maxBytes > 0,
	}
}

func (q *DtQueue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *DtQueue[T]) IsFull() bool {
	return len(q.queue) == cap(q.queue)
}

func (q *DtQueue[T]) Len() int {
	return len(q.queue)
}

func (q *DtQueue[T]) Push(item T) error {
	for q.IsFull() {
		time.Sleep(1 * time.Millisecond)
	}

	if q.checkMemory {
		for atomic.LoadInt64(&q.curBytes) > q.maxBytes {
			time.Sleep(1 * time.Millisecond)
		}
		atomic.AddInt64(&q.curBytes, int64(unsafe.Sizeof(item)))
	}

	q.queue <- item
	return nil
}

func (q *DtQueue[T]) Pop() (T, error) {
	item := <-q.queue

	if q.checkMemory {
		if q.IsEmpty() {
			atomic.StoreInt64(&q.curBytes, 0)
		} else {
			atomic.AddInt64(&q.curBytes, -int64(unsafe.Sizeof(item)))
		}
	}

	return item, nil
}
