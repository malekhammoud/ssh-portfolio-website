package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _,err := p.Run(); err != nil {
		fmt.Println("there was another error: %v", err)
		os.Exit(1)
	}
}
