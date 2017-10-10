package tasks

import "os/exec"

// Command is a task which runs a cli command on the host. You can use it for example
// to run curl or some file system commands.
type Command struct {
	Command string
}

// Run the task on the command line
func (t *Command) Run(log func(text string)) {
	cmd := exec.Command("/bin/sh", "-c", t.Command)

	out, err := cmd.Output()
	if err != nil {
		log(err.Error())
	}

	// TODO Provide an io.Writer as Stdout and Stderr to capture the whole output
	// The following line only returns the last line of the output
	log(string(out))
}
