package worker

import (
	"context"
	"fmt"
	"time"
)

type TickerWorker struct {
	ticker     *time.Ticker
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTickerWorker(ticker *time.Ticker) TickerWorker {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return TickerWorker{
		ticker:     ticker,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (w *TickerWorker) Start(handler func()) {
	fmt.Println("Start ticker worker.")
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			case <-w.ticker.C:
				handler()
			}
		}
	}()
}

func (w *TickerWorker) Stop() {
	w.ticker.Stop()
	fmt.Println("Stop ticker worker.")
	w.cancelFunc()
}
