package internal

import (
	"context"
	"fmt"
	"sync"
	"time"
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

func (q *AsyncQueue) PushJob(ctx context.Context, job AsyncQueueJob) {
	q.Queue <- job
}

func (q *AsyncQueue) PushDelayJob(ctx context.Context, job AsyncQueueJob, delay time.Duration) {
	time.AfterFunc(delay, func() {
		q.Queue <- job
	})
}

type TestJob struct {
	Msg string
}

func (t *TestJob) Execute(ctx context.Context) error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Execute TestJob")
	return nil
}

func (t *TestJob) String() string {
	return t.Msg
}

func Export() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Export Job Start...")
	ctx := context.Background()

	job := &TestJob{"Hi"}
	NewAsyncQueue().PushJob(ctx, job)
	NewAsyncQueue().PushDelayJob(ctx, job, time.Second*3)
	NewAsyncQueue().PushDelayJob(ctx, job, time.Second*5)

	time.Sleep(7 * time.Second)
}
