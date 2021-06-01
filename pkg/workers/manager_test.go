package workers

import (
	"testing"
	"time"
)

func TestCreateManager(t *testing.T) {
	managerWorker := NewManager(func() {
		time.Sleep(1 * time.Second)
	})
	managerWorker.Start()

	managerWorker.Add()
	if managerWorker.Count() != 1 {
		t.Errorf("Manager workers count is not equal 1")
	}

	managerWorker.AddBy(3)
	if managerWorker.Count() != 4 {
		t.Errorf("Manager workers count is not equal 4")
	}

	managerWorker.RemoveBy(3)
	if managerWorker.Count() != 1 {
		t.Errorf("Manager workers count is not equal 1")
	}

	managerWorker.Remove()
	if managerWorker.Count() != 0 {
		t.Errorf("Manager workers count is not equal 0")
	}

	managerWorker.Add()
	if managerWorker.Count() != 1 {
		t.Errorf("Manager workers count is not equal 1")
	}
}
