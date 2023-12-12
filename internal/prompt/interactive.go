package prompt

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

// UserChoice contains pointer to the command that user has chosen in the UI
type UserChoice struct {
	Command *cobra.Command
}

// Model contains the state and forms of UI
type Model struct {
	isFlagMode      bool // will be true if command doesn't have subcommands
	textAreaEnabled bool // will be true if command has flags
	submitButton    string
	textAreaMsg     string // contains a flag desctiption of command
	argument        argument
	table           table.Model
	list            list.Model
	input           textinput.Model
	userChoice      *UserChoice // set command for CLI which user has chosen in the UI
	command         *cobra.Command
	subCommands     []subCommand // contains subcommands of the command
}

func (m Model) Init() tea.Cmd { return nil }

func initialModel(cmd *cobra.Command, userChoice *UserChoice) Model {

	subCommands := getCobraSubCommands(cmd)
	t := createCommandsTableView(tableColumnWidth, subCommands)
	l := createListView("")
	i := createInputView()

	return Model{
		command:     cmd,
		subCommands: subCommands,
		table:       t,
		list:        l,
		input:       i,
		userChoice:  userChoice,
	}
}

func Interactive(cmd *cobra.Command, userChoice *UserChoice) {
	p := tea.NewProgram(initialModel(cmd, userChoice))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, there's been an error: %v", err)
		os.Exit(1)
	}
}
