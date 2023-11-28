package prompt

import (
	"github.com/charmbracelet/lipgloss"
	//tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/table"
)

const listHeight = 10
const listWidth = 30
const tableColumnWidth = 30
const commandTableColumnName = "Command names"

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {

	//table := baseStyle.Render(m.table.View())

	return baseStyle.Render(m.table.View()) + "\n"
}

// createCommandsTableView creates a table of the commands that are available to the user
func createCommandsTableView(width int, subCommands []subCommand) table.Model {
	columns := []table.Column{
		{Title: commandTableColumnName, Width: width},
	}
	rows := []table.Row{}
	for _, val := range subCommands {
		rows = append(rows, table.Row{val.name})
	}
	return createTableTemplate(columns, rows)
}

// createTableTemplate creates table template and applies styles to it
func createTableTemplate(columns []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}
