package main

import (
	"fmt"
	"time"

	"github.com/andr-ik/go-workers/pkg/workers"
)

func main() {
	fmt.Println("Start")

	managerWorker1 := workers.NewManager(func() {
		time.Sleep(1 * time.Second)
		fmt.Println()
	})
	managerWorker1.Start()
	managerWorker1.Add()

	managerWorker2 := workers.NewManager(func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	})

	managerWorker2.Start()

	managerWorker2.AddBy(4)
	time.Sleep(5 * time.Second)
	managerWorker2.Add()
	time.Sleep(5 * time.Second)
	managerWorker2.Remove()
	managerWorker2.Remove()
	managerWorker2.Remove()

	time.Sleep(10 * time.Second)
	managerWorker2.Stop()
	managerWorker1.Stop()

	fmt.Println("")
	fmt.Println("Stop")
}
