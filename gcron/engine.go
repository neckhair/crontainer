package crontainer

import (
	"github.com/spf13/cast"
)

type EngineInterface interface {
	Start()
	Stop()
	AddTask(task *Task) error
}

var Engine EngineInterface = NewCronEngineAdapter()

// Read tasks from config into internal structure
func InitializeFromConfig() {
	// Add single task (from command line)
	if command := Config.GetString("command"); command != "" {
		task := &Task{
			Command:  command,
			Schedule: Config.GetString("schedule"),
		}
		Engine.AddTask(task)
	}

	// Add tasks from list in config file
	for _, taskFromConfig := range cast.ToSlice(Config.Get("tasks")) {
		taskMap := cast.ToStringMapString(taskFromConfig)
		task := &Task{
			Command:  taskMap["command"],
			Schedule: taskMap["schedule"],
		}
		Engine.AddTask(task)
	}
}
