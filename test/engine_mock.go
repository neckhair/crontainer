package test

import "github.com/neckhair/crontainer/crontainer"

type EngineMock struct {
	Tasks []*crontainer.Task
}

func (e *EngineMock) Start() {}
func (e *EngineMock) Stop()  {}
func (e *EngineMock) AddTask(task *crontainer.Task) error {
	e.Tasks = append(e.Tasks, task)
	return nil
}
