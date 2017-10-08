package crontainer

import (
	"log"
	"math/rand"
	"os/exec"
	"os"
)

type Task struct {
	Command  string
	Schedule string
	Name     string
}

func NewTask(command string, schedule string, name string) *Task {
	if name == "" {
		name = randStringBytes(15)
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

    // TODO Provide an io.Writer as Stdout and Stderr to capture the whole output
    cmd.Stdin  = os.Stdin
    cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		t.log(err)
	}

	t.log("Done")
}

func (t *Task) log(text interface{}) {
	log.Printf("[%15s] %s", t.Name, text)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
