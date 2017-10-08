package crontainer

type ConfigurationManager interface {
	GetString(key string) string
	Get(key string) interface{}
}
