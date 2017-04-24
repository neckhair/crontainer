package crontainer

import (
	"github.com/spf13/viper"
)

type ConfigurationManager interface {
	GetString(key string) string
	Get(key string) interface{}
}

var Config ConfigurationManager = viper.GetViper()
