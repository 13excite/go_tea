// Package version implements the version command
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCommand function creates a new cobra command called version
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version of binary",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1.0")
		},
	}
}
