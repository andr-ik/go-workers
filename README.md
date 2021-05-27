Simple implementation create dinamic workers
---

`GroupWorker` depends on `*config.Config`, read file `config/config.txt` and dinamic update num of workers in runtime

```golang
groupWorker := worker.NewGroupWorker(conf)
groupWorker.Start(func() {
    time.Sleep(100 * time.Millisecond)
    fmt.Print(".")
})
```