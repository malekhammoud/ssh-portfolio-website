package main

import "strings"

func (m model) View() string {
	switch m.state {
	case titleView:
		return m.renderTitle()
	case listView:
		return m.renderList()
	case bodyView:
		return m.renderBody()
	}
	return "?"
}

