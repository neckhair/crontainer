package crontainer

type EngineInterface interface {
	Start()
	Stop()
	AddTask(task *Task) error
}

var Engine EngineInterface = NewCronEngineAdapter()

// Read tasks from config into internal structure
func InitializeFromConfig() {
	task := &Task{
		Command:  Config.GetString("command"),
		Schedule: Config.GetString("schedule"),
	}
	Engine.AddTask(task)
}
