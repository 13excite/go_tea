// Package shell implements the cobra command for shell
package tui

import (
	"fmt"

	"gotea/internal/prompt"

	"github.com/spf13/cobra"
)

// NewShellCommand creates a cobra command called auth. It bundles authentication related subcommands
func NewTUICommand() *cobra.Command {
	shellCmd := &cobra.Command{
		Use:   "tui",
		Short: "run terminal UI for the cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TUI STARTED!")
			prompt.Interactive()

			return nil
		},
	}
	return shellCmd
}
