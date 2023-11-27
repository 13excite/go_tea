package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

// Model contains the state and forms of UI
type Model struct {
	table table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func initialModel() Model {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}
	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "São Paulo", "Brazil", "22,429,800"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	return Model{
		table: t,
	}
}

func Interactive() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, there's been an error: %v", err)
		os.Exit(1)
	}
}
