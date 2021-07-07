package tasks

type Executor struct {
	awaitingTasks <-chan Runnable
}

func (e *Executor) Execute(output chan interface{}) {
	for runnable := range e.awaitingTasks {
		result := runnable.Run()
		output <- result
	}
}
