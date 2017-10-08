package crontainer

type Engine interface {
	Initialize(ConfigurationManager) error
	Start()
	Stop()
	AddTask(task *Task) error
}
