// Package create implements the cobra command for creating microservices
package city

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCityCommand creates cobra command called city
func NewCityCommand() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "city",
		Short: "Get an info about weather in a specific city",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("WEATHER CITY STARTED")
			return nil
		},
	}
	return createCmd
}
