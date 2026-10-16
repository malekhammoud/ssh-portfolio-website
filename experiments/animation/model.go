package main

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	counter int 
	direction int
}
func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return tickCmd() 
}

type tickMsg time.Time
func tickCmd() tea.Cmd{
  	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

