package worker

type Worker struct {
	done chan bool
}

func NewWorker() Worker {
	return Worker{
		done: make(chan bool),
	}
}

func (w *Worker) Start(handler func()) {
	go func() {
		for {
			select {
			case <-w.done:
				return
			default:
				handler()
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.done <- true
}
