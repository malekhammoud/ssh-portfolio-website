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
	return titleStyle.Render("I am Builder.") + "\n\n" + 
			hintStyle.Render("Press any key to continue")
}

func (m model) renderList() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("I am Malek Hammoud.")+"\n")

	for i, item := range menueItems {
		if i == m.cursor{
			b.WriteString(selectedStyle.Render("> " + item) + "\n")
		} else {
			b.WriteString(itemStyle.Render(" " + item) + "\n")
		}
	}

	b.WriteString("\n" + hintStyle.Render("up/down move - enter select - q to quit"))
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
	return titleStyle.Render(menueItems[m.cursor]) + 
		   "\n\n" + 
		   bodyStyle.Render(body)+
		   "\n\n" + 
		   hintStyle.Render("esc back - q quit ")
}
