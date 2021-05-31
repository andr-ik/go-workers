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
	updateConfigWorker := worker.NewTicker(ticker)
	updateConfigWorker.Start(func() {
		_ = conf.Update()
		fmt.Println("Update conf")
	})

	fmt.Println("Start")

	managerWorker := worker.NewManager(conf, func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	})
	managerWorker.Start()

	for i := 0; i < conf.GetNumWorker(); i++ {
		managerWorker.Add()
	}
	time.Sleep(5 * time.Second)
	managerWorker.Add()
	time.Sleep(5 * time.Second)
	managerWorker.Remove()
	managerWorker.Remove()
	managerWorker.Remove()

	time.Sleep(10 * time.Second)
	managerWorker.Stop()

	updateConfigWorker.Stop()

	fmt.Println("")
	fmt.Println("Stop")
}
