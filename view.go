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
func (m model) renderTitle() string {
	return "I am Builder. \n\n" + 
			"Press any key to continue"
}

func (m model) renderList() string {
	var b strings.Builder
	b.WriteString("I am Malek Hammoud.\n")

	for i, item := range menueItems {
		if i == m.cursor{
			b.WriteString("> " + item + "\n")
		} else {
			b.WriteString(" " + item + "\n")
		}
	}

	b.WriteString("\nup/down move - enter select\n")
	return b.String()

}
func (m model) renderBody() string {
	var body string
	switch m.cursor {
		case 0:
			body = "I build software that has to work"
		case 1:
			body = "- article1\n-article2\n-article3"
		case 2:
			body = "- project1\n-project2\n-project3"
		default:
			body = "email: malek@malekhammoud.com"
	}
	return menueItems[m.cursor] + "\n\n" + body
}
