package prompt

import (
	"fmt"
)

// keyUpUpdateTextArea updates textArea for the previous flag if up button was pressed
func (m *Model) keyUpUpdateTextArea() {
	if m.table.Focused() && m.command.HasFlags() && m.isFlagMode {
		cmdIndex := m.table.Cursor()
		args := getArgs(m.command)

		// decrease the table cursor index before it's not equal 0
		switch cmdIndex {
		case 0:
			cmdIndex = 0
		default:
			cmdIndex--
		}
		m.textAreaMsg = args[cmdIndex].description
	}
}

// keyUpUpdateTextArea updates textArea for the next flag if down button was pressed
func (m *Model) keyDownUpdateTextArea() {
	if m.table.Focused() && m.command.HasFlags() && m.isFlagMode {
		cmdIndex := m.table.Cursor()
		args := getArgs(m.command)

		// increase the table cursor index before it's <= len(commandFlags)-1
		switch cmdIndex {
		case len(args) - 1:
			cmdIndex = len(args) - 1
		default:
			cmdIndex++
		}
		m.textAreaMsg = getArgs(m.command)[cmdIndex].description
	}
}

func (m *Model) deactivateInput() {
	m.input.PromptStyle = noStyle
	m.input.TextStyle = noStyle
	m.input.SetValue("")
	m.input.Placeholder = ""
	m.input.Blur()
}

func (m *Model) activateInput(cmdIndex int) {
	m.argument = getArgs(m.command)[cmdIndex]
	m.input.PromptStyle = focusedStyle
	m.input.TextStyle = focusedStyle
	m.input.Placeholder = fmt.Sprintf("Enter value for: %s", m.argument.name)
	m.input.Focus()
}

func (m *Model) switchToInputForm() {
	m.table.Blur()
	cmdIndex := m.table.Cursor()
	m.activateInput(cmdIndex)
}

func (m *Model) switchToListForm() {
	m.table.Blur()
	m.input.Blur()
	cmdIndex := m.table.Cursor()
	m.argument = getArgs(m.command)[cmdIndex]
	m.list = createListView(m.argument.name)
}

func (m *Model) switchToTableForm() {
	m.deactivateInput()
	m.list = createListView("")
	m.table.Focus()
}

// List item doesn't have focus. So if the focus isn't set on a input form
// and a table form, then we consider that it's set on a list form
func (m *Model) isListFocused() bool {
	return !m.table.Focused() && !m.input.Focused()
}

// doesHaveParentCommand checks does command have a parent command and it's not a root command
func (m *Model) doesHaveParentCommand() bool {
	return m.command.HasParent() && m.command.Parent().Name() != "gotea"
}

// isSkipEditor returns true if command doesn't have flags and FlagMode is active
func (m *Model) isSkipEditor() bool {
	return m.isFlagMode && !m.command.HasFlags()
}

func (m *Model) switchFromFlagToCommandForm() {
	m.isFlagMode = false
	// textArea must be disabled after switching from flag to command form
	m.textAreaEnabled = false
	m.textAreaMsg = ""
	m.command = m.subCommands[0].command
	m.subCommands = getCobraSubCommands(m.command.Parent())
	m.table = createCommandsTableView(tableColumnWidth, m.subCommands)
}

func (m *Model) switchToCommandOrFlagForm(cmdIndex int) {
	if m.subCommands[cmdIndex].command.HasSubCommands() {
		m.subCommands = getCobraSubCommands(m.subCommands[cmdIndex].command)
		// always set the command with 0 index from the subcommand slice.
		// it is necessary for the correct transition between the forms of commands when pressing enter/esc
		m.command = m.subCommands[0].command
		m.table = createCommandsTableView(tableColumnWidth, m.subCommands)
	} else {
		m.isFlagMode = true
		m.command = m.subCommands[cmdIndex].command
		annotateArgsAsUIRelated(m.command)
		args := getArgs(m.command)
		// if a command has flags then textArea is enabled with desctioption of the first flag(0 index)
		if len(args) > 0 {
			m.textAreaMsg = args[0].description
			m.textAreaEnabled = true
		}
		m.table = createFlagTableView(m.generateFlagTableColumnName(), tableColumnWidth, args)
	}
}

func (m *Model) switchToParentCommand() {
	m.command = m.command.Parent()
	m.userChoice.Command = nil
	m.subCommands = getCobraSubCommands(m.command.Parent())
	m.table = createCommandsTableView(tableColumnWidth, m.subCommands)
}

func (m *Model) generateFlagTableColumnName() string {
	return fmt.Sprintf("Flag names of %s", m.command.Name())
}

func (m *Model) updateTableFromInput(cmdIndex int) {
	m.command.Flags().Set(m.argument.name, m.input.Value())
	m.table = createFlagTableView(m.generateFlagTableColumnName(),
		tableColumnWidth, getArgs(m.command),
	)
	m.deactivateInput()
	m.table.SetCursor(cmdIndex)
	m.table.Focus()
}

func (m *Model) updateTableFromList(cmdIndex int) {
	var title string
	if i, ok := m.list.SelectedItem().(item); ok {
		title = i.FilterValue()
	}
	m.command.Flags().Set(m.argument.name, title)
	m.list = createListView("")
	m.table = createFlagTableView(m.generateFlagTableColumnName(), tableColumnWidth, getArgs(m.command))
	m.table.SetCursor(cmdIndex)
}

// isSubmitButtonActive returns true if flagMode is enabled, and submitButton is focused
func (m *Model) isSubmitButtonActive() bool {
	return m.isFlagMode && (m.submitButton == focusedButton) && !m.table.Focused()
}
