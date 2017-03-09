package gcron

import (
	"errors"
	c "github.com/robfig/cron"
)

type Scheduler struct {
	cron *c.Cron
}

func (s Scheduler) Start(job *Job) error {
	if job.Schedule == "" {
		return errors.New("No schedule defined.")
	}

	if s.cron != nil {
		s.cron.Stop()
	}

	s.cron = c.New()
	s.AddJob(job)
	s.cron.Start()

	return nil
}

func (s Scheduler) Stop() {
	s.cron.Stop()
}

func (s Scheduler) AddJob(job *Job) {
	s.cron.AddFunc(job.Schedule, job.Function())
}
