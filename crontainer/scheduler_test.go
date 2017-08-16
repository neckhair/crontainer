package crontainer_test

import (
	"testing"

	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"

	"github.com/neckhair/crontainer/crontainer"
)

var cronService = mockCronService{}
var task = crontainer.Task{Schedule: "* *"}

// START Mocking the cron service

type mockCronService struct {
	crontainer.CronService
	Started  bool
	Stopped  bool
	AddedJob *cron.Job
}

func (c *mockCronService) Start() {
	c.Started = true
}

func (c *mockCronService) Stop() {
	c.Stopped = true
}

func (c *mockCronService) AddJob(schedule string, cmd cron.Job) error {
	c.AddedJob = &cmd
	return nil
}

// END Mocking

func TestStart(t *testing.T) {
	scheduler := crontainer.NewScheduler()
	scheduler.CronService = &cronService

	scheduler.Start(&task)

	assert.Equal(t, true, cronService.Started)
	assert.NotNil(t, cronService.AddedJob)
}

func TestStartNoScheduleDefined(t *testing.T) {
	job := &crontainer.Task{Schedule: ""}
	scheduler := &crontainer.Scheduler{}

	err := scheduler.Start(job)

	assert.Equal(t, err.Error(), "No schedule defined.")
}
