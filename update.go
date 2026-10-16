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
					case "up", "k":
						if m.cursor > 0 {
							m.cursor --
						}
					case "down", "j":
						if m.cursor < len(menueItems) -1 {
							m.cursor++ 
						}
					case "enter":
						m.state = bodyView
					}
				case bodyView:
					switch msg.String(){
						case "q":
							return m, tea.Quit
						case "esc", "backspace":
							m.state = listView
					}
			}
			
		//TODO: rest of cases(up, down, esc, etc)	
	}
	return m, nil
}

