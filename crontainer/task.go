package crontainer

import (
	uuid "github.com/satori/go.uuid"

	"github.com/neckhair/crontainer/crontainer/tasks"
)

// Task defines the structure of one line of tasks from the config file
type Task struct {
	Schedule string
	Name     string
	Type     string
	Command  *tasks.Command
	Color    int
}

// TaskType defines an interface for all kinds of tasks
type TaskType interface {
	// Run is called by the scheduler. It can return log messages via the channel.
	Run(feedbackChannel chan string)
}

// Run is called by the scheduler to run this specific task. This method then
// defines how it executes the actual task.
func (t *Task) Run() {
	var task TaskType
	feedbackChannel := make(chan string, 5)

	switch t.Type {
	case "command":
		task = t.Command
	}

	t.log("Started")

	go task.Run(feedbackChannel)
	for message := range feedbackChannel {
		t.log(message)
	}

	t.log("Finished") // TODO: Log time
}

// GetName returns the task's name. If none is set it creates a random name.
func (t *Task) GetName() string {
	if t.Name == "" {
		t.Name = uuid.NewV4().String()[:8]
	}
	return t.Name
}

func (t *Task) log(text string) {
	Log(t.GetName(), text)
}
