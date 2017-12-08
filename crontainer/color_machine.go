package crontainer

import (
	"sync"

	"github.com/fatih/color"
)

// ColorizeString colorizes a string with the given color and returns it
func ColorizeString(str string, clr int) string {
	colorFunc := color.New(color.Attribute(clr)).SprintfFunc()
	return colorFunc(str)
}

// ColorMachine provides the interface for a color picker.
type ColorMachine interface {
	Get(key interface{}) int
}

// RoundRobinColorMachine keeps a set of colors and cycles through them when requested with Pluck()
type RoundRobinColorMachine struct {
	availableColors []color.Attribute
	colorMemory     sync.Map
}

var colors = []color.Attribute{
	color.FgBlue, color.FgGreen, color.FgMagenta, color.FgRed, color.FgYellow, color.FgHiCyan}

// NewRoundRobinColorMachine initialises a new round robin color machine
func NewRoundRobinColorMachine() *RoundRobinColorMachine {
	return &RoundRobinColorMachine{colors, sync.Map{}}
}

// Pluck returns the first available color
func (m *RoundRobinColorMachine) Pluck() int {
	var clr color.Attribute

	if len(m.availableColors) == 0 {
		m.availableColors = colors
	}

	clr, m.availableColors = m.availableColors[0], m.availableColors[1:]

	return int(clr)
}

// Get loads the color for a given key. It gets a new color from the list
// if the key is unknown. The color is then remembered for the key.
func (m *RoundRobinColorMachine) Get(key interface{}) int {
	color, ok := m.colorMemory.Load(key)
	if !ok {
		color = colorMachine.Pluck()
		m.colorMemory.Store(key, color)
	}
	return color.(int)
}
