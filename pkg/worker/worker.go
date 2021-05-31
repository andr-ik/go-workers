package worker

import "context"

type Worker struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewWorker() Worker {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return Worker{
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (w *Worker) Start(handler func()) {
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
				handler()
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.cancelFunc()
}
