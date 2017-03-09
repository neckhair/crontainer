package gcron

import "fmt"

import "github.com/robfig/cron"

type Scheduler struct {
}

func (s Scheduler) Run() {
	fmt.Println("---> Begin scheduling <---")

	job := Job{Command: "echo hello world"}

	c := cron.New()
	c.AddFunc("* * * * * *", job.Function())
	c.Start()
}
