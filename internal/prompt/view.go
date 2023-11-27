package prompt

import (
	"github.com/charmbracelet/lipgloss"
	//tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/bubbles/table"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {

	//table := baseStyle.Render(m.table.View())

	return baseStyle.Render(m.table.View()) + "\n"
}
