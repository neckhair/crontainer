package crontainer

import (
	"github.com/robfig/cron"
	"github.com/spf13/cast"
)

// Cron is an interface to a background processor.
type Cron interface {
	Start()
	Stop()
	AddJob(string, cron.Job) error
}

// CronEngine is an adapter for the robfig/cron library
type CronEngine struct {
	Cron Cron
}

// NewCronEngine returns an initialized cron engine
func NewCronEngine() *CronEngine {
	return &CronEngine{cron.New()}
}

// Initialize reads tasks from config into internal structure
func (e *CronEngine) Initialize(config ConfigurationManager) {
	// Add single task (from command line)
	if command := config.GetString("command"); command != "" {
		task := NewTask(command, config.GetString("schedule"), "")
		e.AddTask(task)
	}

	// Add tasks from list in config file
	for _, taskFromConfig := range cast.ToSlice(config.Get("tasks")) {
		taskMap := cast.ToStringMapString(taskFromConfig)
		task := e.newTaskFromMap(taskMap)
		e.AddTask(task)
	}
}

// AddTask adds a task to the job list
func (e *CronEngine) AddTask(task *Task) error {
	return e.Cron.AddJob(task.Schedule, task)
}

// Start starts the cron engine
func (e *CronEngine) Start() {
	e.Cron.Start()
}

// Stop stops the cron engine
func (e *CronEngine) Stop() {
	e.Cron.Stop()
}

func (e *CronEngine) newTaskFromMap(taskMap map[string]string) *Task {
	return NewTask(taskMap["command"], taskMap["schedule"], taskMap["name"])
}
