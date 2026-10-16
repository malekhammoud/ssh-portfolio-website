package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView 
	bodyView
)

type model struct {
	state uint 
	cursor int
	width int 
	height int 
}
var menueItems= []string{"About Me", "Logs", "Projects", "Contact"}

func NewModel() model {
	return model{
		state: titleView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

