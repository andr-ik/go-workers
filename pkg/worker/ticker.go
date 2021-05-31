package worker

import (
	"context"
	"fmt"
	"time"
)

type Ticker struct {
	ticker     *time.Ticker
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTicker(ticker *time.Ticker) Ticker {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return Ticker{
		ticker:     ticker,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (w *Ticker) Start(handler func()) {
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

func (w *Ticker) Stop() {
	w.ticker.Stop()
	fmt.Println("Stop ticker worker.")
	w.cancelFunc()
}
