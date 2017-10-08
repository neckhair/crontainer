package test

import (
	"github.com/neckhair/crontainer/crontainer"
	"github.com/robfig/cron"
)

// CronMock mocks the behaviour of cron.Cron
type CronMock struct {
	WasStarted bool
	WasStopped bool
	Jobs       []*crontainer.Task
}

var cronMock crontainer.Cron = &CronMock{}

func (m *CronMock) Start() {
	m.WasStarted = true
}

func (m *CronMock) Stop() {
	m.WasStopped = true
}

func (m *CronMock) AddJob(schedule string, job cron.Job) error {
	m.Jobs = append(m.Jobs, job.(*crontainer.Task))
	return nil
}
