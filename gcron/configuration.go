package gcron

var Configuration struct {
	Logfile string
	Task    *Task
}

type ConfigurationManager interface {
	GetString(key string) string
}

// InitializeConfig stores configuration data in internal structure
func InitializeConfig(cm ConfigurationManager) {
	Configuration.Logfile = cm.GetString("logfile")

	task := &Task{}
	task.Command = cm.GetString("command")
	task.Schedule = cm.GetString("schedule")
	Configuration.Task = task
}
