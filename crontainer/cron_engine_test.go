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

			task := cron.Jobs[0].(*crontainer.Task)
			assert.Equal(t, "@daily", task.Schedule)
			assert.Equal(t, "test", task.Command.Command)
		})

		t.Run("Adds multiple tasks to the list", func(t *testing.T) {
			config := &test.ConfigManagerMock{}
			cron := &test.CronMock{}
			engine := crontainer.CronEngine{cron}

			config.Tasks = []map[string]interface{}{
				{"name": "cmd 1", "schedule": "@daily"},
				{"name": "cmd 2", "schedule": "@monthly"},
			}

			engine.Initialize(config)

			assert.Equal(t, 2, cap(cron.Jobs))
			firstTask := cron.Jobs[0].(*crontainer.Task)
			assert.Equal(t, "cmd 1", firstTask.Name)
			assert.Equal(t, "@daily", firstTask.Schedule)

			secondTask := cron.Jobs[1].(*crontainer.Task)
			assert.Equal(t, "cmd 2", secondTask.Name)
			assert.Equal(t, "@monthly", secondTask.Schedule)
		})
	})

	t.Run("AddTask()", func(t *testing.T) {
		cron := &test.CronMock{}
		engine := &crontainer.CronEngine{cron}

		task := crontainer.Task{Schedule: "@daily"}
		engine.AddTask(task)

		assert.Equal(t, cap(cron.Jobs), 1, "has 1 task")
		assert.Equal(t, cron.Jobs[0], &task)
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
