package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 10
const listWidth = 30
const tableColumnWidth = 30
const commandTableColumnName = "Command names"

// item implements list.Item interface
type item string

func (i item) FilterValue() string { return string(i) }

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.Copy().PaddingLeft(4).PaddingBottom(1)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("150"))
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	cursorStyle       = focusedStyle.Copy()
	noStyle           = lipgloss.NewStyle()
	blurredStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	focusedButton     = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton     = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	helpStyleRender   = helpStyle.Copy().Width(100).Bold(true).Render
	docStyle          = lipgloss.NewStyle().Margin(1, 2)
)

// View returns the string representation of the main view
func (m Model) View() string {
	if m.showPopup {
		return fmt.Sprintf("\n\n%s\n\n%s\n\n[Enter] OK", m.popupMsg, strings.Repeat("-", 15))
	}
	table := baseStyle.Render(m.table.View())
	input := m.input.View()
	list := m.list.View()
	button := m.submitButton

	helpString := helpStyleRender(`
        Use the arrow keys to navigate.
        Press ctrl+e or space to switch to the input form and set a value manually.
        Press tab to switch between a table form and a submit button.
        Press enter on the submit button to run selected command.
        Press enter on the table form for selecting the available value from drop down list.
        Press esc to return to the previous form or to exit
	`)

	// if textArea is enabled and it's a flag mode
	// then textArea form of the flags description will be shown
	if m.isFlagMode && m.textAreaEnabled {
		tArea := createTextArea(m.textAreaMsg)
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			table,
			lipgloss.JoinVertical(lipgloss.Top, "Flag description", tArea.View()),
			docStyle.Render(list),
		) + "\n\n" + input + "\n" + button + "\n\n" + helpString
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		table,
		docStyle.Render(list),
	) + "\n\n" + input + "\n" + button + "\n\n" + helpString
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

func createInputView() textinput.Model {
	txt := textinput.New()
	txt.Cursor.Style = cursorStyle
	txt.CharLimit = 40

	txt.Placeholder = ""
	txt.PromptStyle = noStyle
	txt.TextStyle = noStyle
	return txt
}

// createTableView creates a table for flags of a cobra command
func createFlagTableView(title string, width int, arguments []argument) table.Model {
	columns := []table.Column{
		{Title: title, Width: 30},
		{Title: "Current value", Width: width + 10},
	}
	rows := []table.Row{}
	for _, val := range arguments {
		rows = append(rows, table.Row{val.name, val.value.String()})
	}
	return createTableTemplate(columns, rows)
}
