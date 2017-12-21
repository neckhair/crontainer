package crontainer_test

import (
	"testing"

	"github.com/neckhair/crontainer/crontainer"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	t.Run("GetName()", func(t *testing.T) {
		taskWithName := crontainer.Task{Name: "task1"}
		assert.Equal(t, taskWithName.GetName(), "task1")

		taskWithoutName := crontainer.Task{}
		assert.NotEmpty(t, taskWithoutName.GetName())
	})
}
