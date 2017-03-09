package gcron

import (
	"fmt"
	"os/exec"
)

type JobFunction func()

type Job struct {
	Schedule string
	Command  string
}

func (j Job) Function() JobFunction {
	return func() {
		cmd := exec.Command("/bin/sh", "-c", j.Command)
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
}
