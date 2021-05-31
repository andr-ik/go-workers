package worker

import (
	"fmt"

	"github.com/andr-ik/go-workers/pkg/config"
)

type Manager struct {
	controller Worker
	pool       []Worker
	conf       *config.Config

	add    chan func()
	remove chan bool
}

func NewManager(conf *config.Config) Manager {
	return Manager{
		controller: NewWorker(),
		pool:       []Worker{},
		conf:       conf,

		add:    make(chan func()),
		remove: make(chan bool),
	}
}

func (d *Manager) Start() {
	fmt.Println("Start manager worker.")
	d.controller.Start(func() {
		for {
			select {
			case h := <-d.add:
				d.addWorker(h)
				return
			case <-d.remove:
				d.removeWorker()
				return
			default:
			}
		}
	})
}

func (d *Manager) Add(handler func()) {
	d.add <- handler
}

func (d *Manager) addWorker(handler func()) {
	job := NewWorker()
	d.pool = append(d.pool, job)
	job.Start(handler)

	fmt.Println("Add worker. Now workers", len(d.pool))
}

func (d *Manager) Remove() {
	d.remove <- true
}

func (d *Manager) removeWorker() {
	removeIndex := len(d.pool) - 1
	d.pool[removeIndex].Stop()
	d.pool = d.pool[:removeIndex]

	fmt.Println("Remove worker. Now workers", len(d.pool))
}

func (d *Manager) Stop() {
	for i := range d.pool {
		fmt.Println("Remove worker.")
		d.pool[i].Stop()
	}
	fmt.Println("Stop manager worker.")
	d.controller.Stop()
}
