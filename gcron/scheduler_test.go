package gcron_test

import (
	"testing"

	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"

	"github.com/neckhair/gcron/gcron"
)

var cronService = mockCronService{}
var task = gcron.Task{Schedule: "* *"}

// START Mocking the cron service

type mockCronService struct {
	gcron.CronService
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
	scheduler := gcron.NewScheduler()
	scheduler.CronService = &cronService

	scheduler.Start(&task)

	assert.Equal(t, true, cronService.Started)
	assert.NotNil(t, cronService.AddedJob)
}

func TestStartNoScheduleDefined(t *testing.T) {
	job := &gcron.Task{Schedule: ""}
	scheduler := &gcron.Scheduler{}

	err := scheduler.Start(job)

	assert.Equal(t, err.Error(), "No schedule defined.")
}
