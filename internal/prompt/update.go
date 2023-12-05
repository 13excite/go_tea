package prompt

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyUp.String():
			m.keyUpUpdateTextArea()
		case tea.KeyDown.String():
			m.keyDownUpdateTextArea()
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
		case tea.KeyCtrlC.String():
			// clear a user choice before exit
			m.userChoice.Command = nil
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
