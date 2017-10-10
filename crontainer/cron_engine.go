package crontainer

import (
	"log"

	"github.com/neckhair/crontainer/crontainer/tasks"
	"github.com/robfig/cron"
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
		commandTask := &tasks.Command{Command: command}
		task := Task{Schedule: config.GetString("schedule"), Type: "command", Command: commandTask}
		e.AddTask(task)
	}

	e.loadTasksFromConfig(config)
}

// AddTask adds a task to the job list
func (e *CronEngine) AddTask(task Task) error {
	return e.Cron.AddJob(task.Schedule, &task)
}

// Start starts the cron engine
func (e *CronEngine) Start() {
	e.Cron.Start()
}

// Stop stops the cron engine
func (e *CronEngine) Stop() {
	e.Cron.Stop()
}

func (e *CronEngine) loadTasksFromConfig(config ConfigurationManager) {
	var tasksFromConfig []Task
	if err := config.UnmarshalKey("tasks", &tasksFromConfig); err != nil {
		log.Printf("Could not load tasks from config: %v\n", err.Error())
		return
	}

	for _, taskFromConfig := range tasksFromConfig {
		e.AddTask(taskFromConfig)
	}
}
