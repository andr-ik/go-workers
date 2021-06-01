package workers

import (
	"fmt"
)

type Manager struct {
	controller Worker
	handler    func()
	pool       []Worker

	add    chan uint
	remove chan uint
}

func NewManager(handler func()) Manager {
	manager := Manager{
		controller: NewWorker(),
		handler:    handler,
		pool:       []Worker{},

		add:    make(chan uint),
		remove: make(chan uint),
	}

	return manager
}

func (d *Manager) Start() {
	fmt.Println("Start manager worker.")
	d.controller.Start(func() {
		for {
			select {
			case count := <-d.add:
				for i := 0; i < int(count); i++ {
					d.addWorker()
				}
				return
			case count := <-d.remove:
				for i := 0; i < int(count); i++ {
					d.removeWorker()
				}
				return
			default:
			}
		}
	})
}

func (d *Manager) Add() {
	d.add <- 1
}

func (d *Manager) AddBy(count uint) {
	d.add <- count
}

func (d *Manager) addWorker() {
	job := NewWorker()
	d.pool = append(d.pool, job)
	job.Start(d.handler)

	fmt.Println("Add worker. Now workers", len(d.pool))
}

func (d *Manager) Remove() {
	d.remove <- 1
}

func (d *Manager) RemoveBy(count uint) {
	d.remove <- count
}

func (d *Manager) removeWorker() {
	removeIndex := len(d.pool) - 1
	d.pool[removeIndex].Stop()
	d.pool = d.pool[:removeIndex]

	fmt.Println("Remove worker. Now workers", len(d.pool))
}

func (d *Manager) Count() int {
	return len(d.pool)
}

func (d *Manager) Stop() {
	for i := range d.pool {
		fmt.Println("Remove worker.")
		d.pool[i].Stop()
	}
	fmt.Println("Stop manager worker.")
	d.controller.Stop()
}
