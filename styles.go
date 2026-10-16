package main
import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
				 Bold(true).
				 Foreground(lipgloss.Color("26"))
	selectedStyle = lipgloss.NewStyle().
				 Foreground(lipgloss.Color("230")).
				 Background(lipgloss.Color("71")).
				 Padding(0, 1)
	itemStyle = lipgloss.NewStyle().Padding(0, 1)
	bodyStyle = lipgloss.NewStyle().Padding(1, 2)
	hintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)
