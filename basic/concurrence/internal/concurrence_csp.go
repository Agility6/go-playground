package internal

import (
	"context"
	"sync"
)

var (
	once  = sync.Once{}
	queue *AsyncQueue
)

type AsyncQueueJob interface {
	Execute(ctx context.Context) error
	String() string
}

type AsyncQueue struct {
	Queue chan AsyncQueueJob
}

func NewAsyncQueue() *AsyncQueue {
	if queue != nil {
		return queue
	}
	once.Do(func() {
		queue = &AsyncQueue{
			Queue: make(chan AsyncQueueJob, 1024),
		}
		go func() {
			for {
				job := <-queue.Queue
				job.Execute(context.Background())
			}
		}()
	})
	return queue
}
