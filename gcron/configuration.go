package gcron

var Configuration struct {
	Logfile string
	Job     *Job
}

type ConfigurationManager interface {
	GetString(key string) string
}

// InitializeConfig stores configuration data in internal structure
func InitializeConfig(cm ConfigurationManager) {
	Configuration.Logfile = cm.GetString("logfile")

	job := &Job{}
	job.Command = cm.GetString("command")
	job.Schedule = cm.GetString("schedule")
	Configuration.Job = job
}
