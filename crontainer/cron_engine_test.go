package crontainer_test

import (
	"testing"

	"github.com/neckhair/crontainer/crontainer"
	"github.com/neckhair/crontainer/test"
	"github.com/stretchr/testify/assert"
)

func TestCronEngine(t *testing.T) {
	t.Run("Initialize()", func(t *testing.T) {
		t.Run("Adds the task to the list", func(t *testing.T) {
			config := &test.ConfigManagerMock{}
			cron := &test.CronMock{}
			engine := crontainer.CronEngine{cron}

			config.SingleValue = map[string]string{"command": "test", "schedule": "@daily"}

			engine.Initialize(config)

			task := cron.Jobs[0]
			assert.Equal(t, "test", task.Command)
		})

		t.Run("Adds multiple tasks to the list", func(t *testing.T) {
			config := &test.ConfigManagerMock{}
			cron := &test.CronMock{}
			engine := crontainer.CronEngine{cron}

			config.Values = []map[string]interface{}{
				{"command": "test 1", "schedule": "@daily"},
				{"command": "test 2", "schedule": "@monthly"},
			}

			engine.Initialize(config)

			firstTask := cron.Jobs[0]
			assert.Equal(t, "test 1", firstTask.Command)
			assert.Equal(t, "@daily", firstTask.Schedule)

			secondTask := cron.Jobs[1]
			assert.Equal(t, "test 2", secondTask.Command)
			assert.Equal(t, "@monthly", secondTask.Schedule)
		})
	})

	t.Run("AddTask()", func(t *testing.T) {
		cron := &test.CronMock{}
		engine := &crontainer.CronEngine{cron}

		task := &crontainer.Task{Schedule: "@daily"}
		engine.AddTask(task)

		assert.Equal(t, cap(cron.Jobs), 1, "has 1 task")
		assert.Equal(t, cron.Jobs[0], task)
	})

	t.Run("Start()", func(t *testing.T) {
		cron := &test.CronMock{}
		engine := &crontainer.CronEngine{cron}

		engine.Start()

		assert.True(t, cron.WasStarted)
	})

	t.Run("Stop()", func(t *testing.T) {
		cron := &test.CronMock{}
		engine := &crontainer.CronEngine{cron}

		engine.Stop()

		assert.True(t, cron.WasStopped)
	})
}
