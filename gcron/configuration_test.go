package gcron_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/neckhair/gcron/gcron"
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

	gcron.InitializeConfig(&cm)

	assert.Equal(t, gcron.Configuration.Logfile, logfile)
	assert.Equal(t, gcron.Configuration.Task.Command, command)
	assert.Equal(t, gcron.Configuration.Task.Schedule, schedule)
}
