package tasks

type Queue struct {
	TaskQueue chan Runnable
}

func NewQueue(size int) *Queue {
	return &Queue{
		TaskQueue: make(chan Runnable, size),
	}
}

func (q *Queue) Push(r Runnable) {
	q.TaskQueue <- r
}

func (q *Queue) Pop() Runnable {
	return <-q.TaskQueue
}
