package crontainer

import (
	"log"
)

var colorMachine = NewRoundRobinColorMachine()

// Log loggs a message indicating the task's name
func Log(taskName string, message string) {
	color := colorMachine.Get(taskName)
	log.Printf("[%8s] %s", ColorizeString(taskName, color), message)
}
