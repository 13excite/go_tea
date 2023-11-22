package prompt

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

// Model contains the state and forms of UI
type Model struct {
	table table.Model
}

func (m Model) Init() tea.Cmd { return nil }
