package prompt

import (
	"github.com/charmbracelet/lipgloss"
	//tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
)

const listHeight = 10
const listWidth = 30
const tableColumnWidth = 30
const commandTableColumnName = "Command names"

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.Copy().PaddingLeft(4).PaddingBottom(1)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("150"))
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
)

// item implements list.Item interface
type item string

func (i item) FilterValue() string { return string(i) }

// View returns the string representation of the main view
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

// createTextArea creates a text area with the given text
func createTextArea(text string) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = "Flag description"
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))
	ta.SetValue(text)
	ta.ShowLineNumbers = false
	ta.Blur()
	return ta
}

func createListView(flagName string) list.Model {
	// items contains results of available autocompletion result
	var items []list.Item

	// flags whose values can be received from fake-api
	switch flagName {
	case "name":
		items = getCityNamesCompletion()
	case "":
		items = []list.Item{}
	default:
		items = []list.Item{item("No items found.")}
	}

	l := list.New(items, itemDelegate{}, listWidth, listHeight)
	l.Title = "Available options"

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.DisableQuitKeybindings()
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return l
}
