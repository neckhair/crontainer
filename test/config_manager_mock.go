package test

import "github.com/neckhair/crontainer/crontainer"

type ConfigManagerMock struct {
	Tasks       []map[string]interface{}
	SingleValue map[string]string
}

func (cm *ConfigManagerMock) GetString(key string) string {
	return cm.SingleValue[key]
}

func (cm *ConfigManagerMock) Get(key string) interface{} {
	return cm.Tasks
}

func (cm *ConfigManagerMock) UnmarshalKey(key string, rawVal interface{}) error {
	switch rawVal := rawVal.(type) {
	case *[]crontainer.Task:
		for _, task := range cm.Tasks {
			*rawVal = append(*rawVal, crontainer.Task{Name: task["name"].(string), Schedule: task["schedule"].(string)})
		}
	default:
		panic("invalid type")
	}
	return nil
}
