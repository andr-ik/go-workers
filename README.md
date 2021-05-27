Simple implementation create dinamic workers

```golang
groupWorker := worker.NewGroupWorker(conf)
groupWorker.Start(func() {
    time.Sleep(100 * time.Millisecond)
    fmt.Print(".")
})
```