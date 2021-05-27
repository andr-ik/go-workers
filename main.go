package main

import (
	"fmt"
	"time"

	"github.com/andr-ik/go-workers/pkg/config"
	"github.com/andr-ik/go-workers/pkg/worker"
)

func main() {
	conf := config.NewConfig()
	err := conf.Update()
	if err != nil {
		return
	}

	ticker := time.NewTicker(time.Second)
	updateConfigWorker := worker.NewTickerWorker(ticker)
	updateConfigWorker.Start(func() {
		_ = conf.Update()
		fmt.Println("Update conf")
	})

	fmt.Println("Start")

	groupWorker := worker.NewGroupWorker(conf)
	groupWorker.Start(func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	})

	time.Sleep(20 * time.Second)
	groupWorker.Stop()
	updateConfigWorker.Stop()

	fmt.Println("")
	fmt.Println("Stop")
}
