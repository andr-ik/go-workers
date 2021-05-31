Simple implementation create dinamic workers
---

`worker.Manager` service for dinamic create and remove workers in runtime

```golang
managerWorker := worker.NewManager(func() {
	time.Sleep(1 * time.Second)
    fmt.Println()
})
managerWorker.Start()
managerWorker.Add()

managerWorker.Stop()
```