package prompt

import (
	"github.com/spf13/cobra"
)

// subCommand represents a single subcommand for the command
type subCommand struct {
	name    string
	command *cobra.Command
}

// ignoredCommands contains a list of commands that should be ignored
var ignoredCommands = []string{"help", "version", "completion", "tui"}

// getCobraSubCommands returns a list of subcommands of the command
func getCobraSubCommands(cmd *cobra.Command) []subCommand {
	commands := []subCommand{}
	for _, command := range cmd.Commands() {
		// mark command for O(1)
		annotateCommandAdUIRelated(command)
		val, ok := command.Annotations[uiAnnotationKey]
		if ok && (val == "true") {
			commands = append(commands, subCommand{
				name:    command.Name(),
				command: command,
			})
		}
	}
	return commands
}

// annotateCommandAdUIRelated adds an annotation that indicates whether the command belongs to the UI
func annotateCommandAdUIRelated(cmd *cobra.Command) {
	// if command doesn't have annotation then mark it
	_, ok := cmd.Annotations[uiAnnotationKey]
	if !ok {
		if cmd.Annotations == nil {
			cmd.Annotations = map[string]string{}
		}
		if stringInSlice(cmd.Name(), ignoredCommands) {
			cmd.Annotations[uiAnnotationKey] = "false"
		} else {
			cmd.Annotations[uiAnnotationKey] = "true"
		}
	}
}
