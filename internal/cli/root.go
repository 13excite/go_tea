// Package cli implements the command line interface that is used from the entrypoint.
package cli

import (
	"os"

	"gotea/internal/cli/tui"
	"gotea/internal/cli/weather"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotea",
	Short: "A gotea CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(weather.NewWeatherCommand())
	rootCmd.AddCommand(tui.NewTUICommand())
}
