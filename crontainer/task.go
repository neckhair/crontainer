package crontainer

import (
	"log"
	"os/exec"
)

type Task struct {
	Command  string
	Schedule string
}

// Run the task on the command line
func (t *Task) Run() {
	log.Printf("-> %s\n", t.Command)
	cmd := exec.Command("/bin/sh", "-c", t.Command)

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	// TODO Provide an io.Writer as Stdout and Stderr to capture the whole output
	// The following line only returns the last line of the output
	log.Println(string(out))

	log.Println("-> Task done.")
}
