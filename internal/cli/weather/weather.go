// Package weather implements the cobra command for weather
package weather

import (
	"gotea/internal/cli/weather/city"

	"github.com/spf13/cobra"
)

// NewWeatherCommand creates weather cobra command
func NewWeatherCommand() *cobra.Command {
	weatherCmd := &cobra.Command{
		Use:   "weather",
		Short: "gets of weather",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
		DisableFlagParsing:    true,
		DisableFlagsInUseLine: true,
	}

	cityCmd := city.NewCityCommand()
	weatherCmd.AddCommand(cityCmd)
	return weatherCmd
}
