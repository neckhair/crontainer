package test

import (
	"log"

	"github.com/neckhair/crontainer/crontainer"
)

type ConfigManagerMock struct {
	Values      interface{}
	SingleValue map[string]string
}

func (cm *ConfigManagerMock) GetString(key string) string {
	return cm.SingleValue[key]
}

func (cm *ConfigManagerMock) Get(key string) interface{} {
	return cm.Values
}

func (cm *ConfigManagerMock) UnmarshalKey(key string, rawVal interface{}) error {
	log.Println(rawVal)
	rawVal = []crontainer.Task{crontainer.Task{Name: "test"}}
	log.Println(rawVal)
	log.Println("that's it")
	return nil
}
