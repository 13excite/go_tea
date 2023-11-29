package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// UserChoice contains pointer to the command that user has chosen in the UI
type UserChoice struct {
	Command *cobra.Command
}

// Model contains the state and forms of UI
type Model struct {
	command     *cobra.Command
	subCommands []subCommand // contains subcommands of the command
	table       table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func initialModel(cmd *cobra.Command, userChoice *UserChoice) Model {

	subCommands := getCobraSubCommands(cmd)
	t := createCommandsTableView(tableColumnWidth, subCommands)

	return Model{
		command:     cmd,
		subCommands: subCommands,
		table:       t,
	}
}

func Interactive(cmd *cobra.Command, userChoice *UserChoice) {
	p := tea.NewProgram(initialModel(cmd, userChoice))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, there's been an error: %v", err)
		os.Exit(1)
	}
}
