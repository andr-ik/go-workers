package worker

import (
	"fmt"

	"github.com/andr-ik/go-workers/pkg/config"
)

type GroupWorker struct {
	controller Worker
	workers    []Worker
	conf       *config.Config
}

func NewGroupWorker(conf *config.Config) GroupWorker {
	return GroupWorker{
		controller: NewWorker(),
		workers:    []Worker{},
		conf:       conf,
	}
}

func (d *GroupWorker) Start(handler func()) {
	fmt.Println("Start group worker.")
	d.controller.Start(func() {
		need := d.conf.GetNumWorker()
		remove := len(d.workers) - need
		if remove == 0 {
			return
		}

		if remove > 0 {
			removeJobs := d.workers[need:]
			d.workers = d.workers[:need]
			for i := range removeJobs {
				fmt.Println("Remove worker.")
				removeJobs[i].Stop()
			}
		}

		if remove < 0 {
			for i := 0; i < remove*-1; i++ {
				job := NewWorker()
				d.workers = append(d.workers, job)
				fmt.Println("Add worker.")
				job.Start(handler)
			}
		}

		fmt.Println("Now workers ", len(d.workers))
	})
}

func (d *GroupWorker) Stop() {
	for i := range d.workers {
		fmt.Println("Remove worker.")
		d.workers[i].Stop()
	}
	fmt.Println("Stop group worker.")
	d.controller.Stop()
}
