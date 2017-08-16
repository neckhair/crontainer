package crontainer_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/neckhair/crontainer/crontainer"
)

type mockConfigurationManager struct {
	Logfile  string
	Command  string
	Schedule string
}

func (cm *mockConfigurationManager) GetString(key string) string {
	switch key {
	case "logfile":
		return cm.Logfile
	case "command":
		return cm.Command
	case "schedule":
		return cm.Schedule
	}
	return ""
}

func TestInitializeConfig(t *testing.T) {
	const logfile = "sample_file.log"
	const command = "test -v"
	const schedule = "* * *"

	cm := mockConfigurationManager{
		Logfile:  logfile,
		Command:  command,
		Schedule: schedule}

	crontainer.InitializeConfig(&cm)

	assert.Equal(t, crontainer.Configuration.Logfile, logfile)
	assert.Equal(t, crontainer.Configuration.Task.Command, command)
	assert.Equal(t, crontainer.Configuration.Task.Schedule, schedule)
}
