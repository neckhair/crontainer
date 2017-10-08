package crontainer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neckhair/crontainer/crontainer"
	"github.com/neckhair/crontainer/test"
)

var configurationManager *test.ConfigManagerMock
var testEngine *test.EngineMock

func ResetData() {
	configurationManager = &test.ConfigManagerMock{}
	testEngine = &test.EngineMock{}

	crontainer.Config = configurationManager
	crontainer.Engine = testEngine
}

func TestInitializeFromConfig(t *testing.T) {
	t.Run("Adds the task to the list", func(t *testing.T) {
		ResetData()

		configurationManager.SingleValue = map[string]string{"command": "test", "schedule": "@daily"}

		crontainer.InitializeFromConfig()

		assert.Equal(t, 1, cap(testEngine.Tasks))
		task := testEngine.Tasks[0]
		assert.Equal(t, "test", task.Command)
	})

	t.Run("Adds multiple tasks to the list", func(t *testing.T) {
		ResetData()

		configurationManager.Values = []map[string]interface{}{
			{"command": "test 1", "schedule": "@daily"},
			{"command": "test 2", "schedule": "@monthly"},
		}

		crontainer.InitializeFromConfig()

		assert.Equal(t, 2, cap(testEngine.Tasks))

		assert.Equal(t, "test 1", testEngine.Tasks[0].Command)
		assert.Equal(t, "@daily", testEngine.Tasks[0].Schedule)

		assert.Equal(t, "test 2", testEngine.Tasks[1].Command)
		assert.Equal(t, "@monthly", testEngine.Tasks[1].Schedule)
	})
}
