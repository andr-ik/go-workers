package worker

import (
	"fmt"
	"time"
)

type TickerWorker struct {
	ticker *time.Ticker
	done   chan bool
}

func NewTickerWorker(ticker *time.Ticker) TickerWorker {
	return TickerWorker{
		ticker: ticker,
		done:   make(chan bool),
	}
}

func (w *TickerWorker) Start(handler func()) {
	fmt.Println("Start ticker worker.")
	go func() {
		for {
			select {
			case <-w.done:
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
	w.done <- true
}
