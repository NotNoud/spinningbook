package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/NotNoud/spinningbook/internal/tui"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Fprintln(os.Stderr, "subcommands komen later (run/list/init)")
		os.Exit(1)
	}

	if _, err := tea.NewProgram(tui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
