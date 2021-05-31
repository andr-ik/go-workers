package main

import (
	"fmt"
	"time"

	"github.com/andr-ik/go-workers/pkg/worker"
)

func main() {
	fmt.Println("Start")

	managerWorker := worker.NewManager()
	managerWorker.Start()
	handler := func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}

	for i := 0; i < 4; i++ {
		managerWorker.Add(handler)
	}
	time.Sleep(5 * time.Second)
	managerWorker.Add(handler)
	time.Sleep(5 * time.Second)
	managerWorker.Remove()
	managerWorker.Remove()
	managerWorker.Remove()

	time.Sleep(10 * time.Second)
	managerWorker.Stop()

	fmt.Println("")
	fmt.Println("Stop")
}
