package gcron

import (
	"errors"
	"github.com/robfig/cron"
)

type CronService interface {
	Start()
	Stop()
	AddJob(schedule string, cmd cron.Job) error
}

type Scheduler struct {
	CronService CronService
}

func NewScheduler() *Scheduler {
	return &Scheduler{cron.New()}
}

func (s Scheduler) Start(task *Task) error {
	if task.Schedule == "" {
		return errors.New("No schedule defined.")
	}

	if s.CronService != nil {
		s.CronService.Stop()
	}

	s.AddTask(task)
	s.CronService.Start()

	return nil
}

func (s Scheduler) Stop() {
	s.CronService.Stop()
}

func (s Scheduler) AddTask(task *Task) {
	s.CronService.AddJob(task.Schedule, task)
}
