package gcron

var Configuration ConfigurationData

type ConfigurationData struct {
	Logfile string
	Job     *Job
}
