package crontainer

import (
	"log"
	"os/exec"
	"github.com/satori/go.uuid"
)

type Task struct {
	Command  string
	Schedule string
	Name     string
}

func NewTask(command string, schedule string, name string) *Task {
	if name == "" {
		name = uuid.NewV4().String()
	}
	return &Task{
		Command:  command,
		Schedule: schedule,
		Name:     name,
	}
}

// Run the task on the command line
func (t *Task) Run() {
	t.log("Started")
	cmd := exec.Command("/bin/sh", "-c", t.Command)

	out, err := cmd.Output()
	if err != nil {
		t.log(err)
	}

	// TODO Provide an io.Writer as Stdout and Stderr to capture the whole output
	// The following line only returns the last line of the output
	t.log(string(out))

	t.log("Done")
}

func (t *Task) log(text interface{}) {
	log.Printf("[%8s] %s", t.Name, text)
}
