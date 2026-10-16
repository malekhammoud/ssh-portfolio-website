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
			switch m.state {
				case titleView:
					m.state = listView
				case listView:
					switch msg.String(){
					case "q":
						return m, tea.Quit
					}
				case bodyView:
					switch msg.String(){
					case "q":
						return m, tea.Quit
					}
			}
			
		//TODO: rest of cases(up, down, esc, etc)	
	}
	return m, nil
}
