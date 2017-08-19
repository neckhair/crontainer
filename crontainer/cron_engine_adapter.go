package crontainer

import (
	"github.com/robfig/cron"
)

type CronEngineAdapter struct {
	service *cron.Cron
}

func NewCronEngineAdapter() *CronEngineAdapter {
	return &CronEngineAdapter{cron.New()}
}

func (a *CronEngineAdapter) AddTask(task *Task) error {
	return a.service.AddJob(task.Schedule, task)
}

func (a *CronEngineAdapter) Start() {
	a.service.Start()
}

func (a *CronEngineAdapter) Stop() {
	a.service.Stop()
}
