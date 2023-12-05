package prompt

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
