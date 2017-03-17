package gcron

import (
	"fmt"
	"os/exec"
)

type Task struct {
	Command  string
	Schedule string
}

func (t *Task) Run() {
	cmd := exec.Command("/bin/sh", "-c", t.Command)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
