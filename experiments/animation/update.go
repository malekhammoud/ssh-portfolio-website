package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
		//TODO: case window sizing
		case tea.KeyMsg:
			if  msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		case tickMsg:
			if m.counter == 9 {
				m.direction = -1
			}else if m.counter == 0{
				m.direction  = 1
			}
			m.counter+= m.direction
			return m, tickCmd()

	}
	return m, nil
}

