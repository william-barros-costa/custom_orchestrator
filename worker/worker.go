package worker

import (
	"fmt"
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Stats")
}

func (w *Worker) RunTask() {
	fmt.Println("Task")
}

func (w *Worker) StartTask() {
	fmt.Println("start")
}

func (w *Worker) StopTask() {
	fmt.Println("stop")
}
