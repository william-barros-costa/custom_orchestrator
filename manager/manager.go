package manager

import (
	"fmt"
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Selecting worker")
}

func (m *Manager) UpdateTask() {
	fmt.Println("Updating Task")
}

func (m *Manager) SendWork() {
	fmt.Println("sending work")
}
