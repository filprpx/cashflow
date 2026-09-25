package main

import (
	"fmt"
	"os"

	"cashflow/internal/tui"
	tea "charm.land/bubbletea/v2"
)

func main() {
	program := tea.NewProgram(tui.NewModel())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
