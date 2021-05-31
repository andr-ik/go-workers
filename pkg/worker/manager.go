package worker

import (
	"fmt"
)

type Manager struct {
	controller Worker
	handler    func()
	pool       []Worker

	add    chan bool
	remove chan bool
}

func NewManager(handler func()) Manager {
	manager := Manager{
		controller: NewWorker(),
		handler:    handler,
		pool:       []Worker{},

		add:    make(chan bool),
		remove: make(chan bool),
	}

	return manager
}

func (d *Manager) Start() {
	fmt.Println("Start manager worker.")
	d.controller.Start(func() {
		for {
			select {
			case <-d.add:
				d.addWorker()
				return
			case <-d.remove:
				d.removeWorker()
				return
			default:
			}
		}
	})
}

func (d *Manager) Add() {
	d.add <- true
}

func (d *Manager) addWorker() {
	job := NewWorker()
	d.pool = append(d.pool, job)
	job.Start(d.handler)

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
