package crontainer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/neckhair/crontainer/crontainer"
)

type mockEngine struct {
	Tasks []*crontainer.Task
}

func (e *mockEngine) Start() {}
func (e *mockEngine) Stop()  {}
func (e *mockEngine) AddTask(task *crontainer.Task) error {
	e.Tasks = append(e.Tasks, task)
	return nil
}

type mockConfigManager struct {
	Values map[string]string
}

func (cm *mockConfigManager) GetString(key string) string {
	return cm.Values[key]
}

func TestInitializeFromConfig(t *testing.T) {
	configurationManager := &mockConfigManager{}
	testEngine := &mockEngine{}

	configurationManager.Values = map[string]string{
		"command":  "test",
		"schedule": "@daily"}
	crontainer.Config = configurationManager
	crontainer.Engine = testEngine

	t.Run("Adds the task to the list", func(t *testing.T) {
		crontainer.InitializeFromConfig()

		assert.Equal(t, 1, cap(testEngine.Tasks))
		task := testEngine.Tasks[0]
		assert.Equal(t, "test", task.Command)
	})
}
