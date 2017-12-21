package crontainer_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/neckhair/crontainer/crontainer"
	"github.com/stretchr/testify/assert"
)

func TestRoundRobinColorMachine(t *testing.T) {
	t.Run("Pluck()", func(t *testing.T) {
		machine := crontainer.NewRoundRobinColorMachine()
		machine.AvailableColors = []color.Attribute{color.FgBlue, color.FgGreen}

		assert.Equal(t, machine.Pluck(), int(color.FgBlue))
		assert.Equal(t, machine.Pluck(), int(color.FgGreen))
		assert.Equal(t, machine.Pluck(), int(color.FgBlue))
	})

	t.Run("Get()", func(t *testing.T) {
		machine := crontainer.NewRoundRobinColorMachine()
		machine.AvailableColors = []color.Attribute{color.FgBlue, color.FgGreen}

		assert.Equal(t, machine.Get("task-blue"), int(color.FgBlue))
		assert.Equal(t, machine.Get("task-green"), int(color.FgGreen))
		assert.Equal(t, machine.Get("task-blue"), int(color.FgBlue))
	})
}
