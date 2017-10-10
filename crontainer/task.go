package crontainer

import (
	"log"

	"github.com/satori/go.uuid"

	"github.com/neckhair/crontainer/crontainer/tasks"
)

type Task struct {
	Schedule string
	Name     string
	Type     string
	Command  *tasks.Command
}

type TaskType interface {
	Run(log func(text string))
}

func (t *Task) Run() {
	var task TaskType

	switch t.Type {
	case "command":
		task = t.Command
	}

	t.log("Started")
	task.Run(t.log)
	t.log("Finished") // TODO: Log time
}

func (t *Task) GetName() string {
	if t.Name == "" {
		t.Name = uuid.NewV4().String()[:8]
		return t.Name
	}
	return t.Name
}

func (t *Task) log(text string) {
	log.Printf("[%8s] %s", t.GetName(), text)
}
