package prompt

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyUp.String():
			m.keyUpUpdateTextArea()
		case tea.KeyDown.String():
			m.keyDownUpdateTextArea()
		case "ctrl+e", tea.KeySpace.String():
			// enable the input form only when table is active
			if !m.isSkipEditor() && m.table.Focused() {
				m.switchToInputForm()
			}
			return m, nil
		case "tab":
			if m.isFlagMode {
				if m.table.Focused() {
					m.table.Blur()
					m.submitButton = focusedButton
				} else if !m.table.Focused() && !m.input.Focused() {
					m.table.Focus()
					m.submitButton = blurredButton
				}
			}
		case tea.KeyEscape.String():
			// blur the button after each ESC pressing
			m.submitButton = blurredButton

			// if popup view is active switch to a table form
			if m.showPopup {
				m.switchFromPopUpToTable()
				return m, nil
			}

			// switch to a table form if list or input form selected
			if m.isListFocused() || m.input.Focused() {
				m.switchToTableForm()
				return m, nil
			}
			// switch to a command form if flag screen is active
			if m.isFlagMode {
				m.switchFromFlagToCommandForm()
				return m, nil
			}
			// switch to a parent command if command exists
			if m.doesHaveParentCommand() {
				m.switchToParentCommand()
				return m, nil
			}
			// or exit
			return m, tea.Quit
		case tea.KeyCtrlC.String():
			// clear a user choice before exit
			m.userChoice.Command = nil
			return m, tea.Quit
		case "enter":
			if m.showPopup {
				m.switchFromPopUpToTable()
				return m, nil
			}
			cmdIndex := m.table.Cursor()
			if m.table.Focused() {
				// if command doesn't have flags just skip rendering after "enter" pressed
				if m.isSkipEditor() {
					return m, nil
				}
				// switch to a list form after press "enter" on a table form,
				// or render a form for subcommands
				if m.isFlagMode {
					m.switchToListForm()
				} else {
					m.switchToCommandOrFlagForm(cmdIndex)
				}
				return m, nil
			}
			if m.isSubmitButtonActive() {
				// if user input is incorrect - show popup
				if !m.isUserInputCorrect() {
					return m, nil
				}
				m.userChoice.Command = m.command
				return m, tea.Quit
			}
			// switch to a table form after press "enter" on a list form
			if m.isListFocused() {
				m.updateTableFromList(cmdIndex)
				return m, nil
			}
			// switch to a table form after press "enter" on a input form
			if m.input.Focused() {
				m.updateTableFromInput(cmdIndex)
				return m, nil
			}
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
