package gcron

import (
	"fmt"
	"os/exec"
)

type JobFunction func()

type Job struct {
	Command  string
	Schedule string
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
