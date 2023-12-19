// Package shell implements the cobra command for shell
package tui

import (
	"fmt"
	"gotea/internal/prompt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewShellCommand creates a cobra command called auth. It bundles authentication related subcommands
func NewTUICommand() *cobra.Command {
	shellCmd := &cobra.Command{
		Use:   "tui",
		Short: "run terminal UI for the cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdSelectedByUser := &prompt.UserChoice{}
			prompt.Interactive(cmd.Parent(), cmdSelectedByUser)

			var argList []string
			if cmdSelectedByUser.Command != nil {
				// generate a list of arguments for the Run method of commands
				// We cannot use the method Execute here. Therefore, we need to get
				// values which were set by the user and pass them to the Run method
				cmdSelectedByUser.Command.Flags().VisitAll(func(flag *pflag.Flag) {
					argList = append(argList, "--"+flag.Name, flag.Value.String())
				})
				// if command has a preRun method, then call it before Run method
				// this method can be used for validation of user input
				if cmdSelectedByUser.Command.PreRunE != nil {
					err := cmdSelectedByUser.Command.PreRunE(cmdSelectedByUser.Command, argList)
					if err != nil {
						return err
					}
				}
				err := cmdSelectedByUser.Command.RunE(cmdSelectedByUser.Command, argList)
				if err != nil {
					return fmt.Errorf("error apierred during the execution of the command: %v", err)
				}
			}

			return nil
		},
	}
	return shellCmd
}
