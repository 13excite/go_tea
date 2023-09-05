// Package create implements the cobra command for creating microservices
package city

import (
	"fmt"
	"log"
	"strings"

	"gotea/pkg/api"

	"github.com/spf13/cobra"
)

var (
	cityName string
	// cityMap contains valid values of the 'cityName' argument
	cityMap = map[string]struct{}{
		"Berlin":    {},
		"Munchen":   {},
		"Frankfurt": {},
		"Leipzig":   {},
		"Longon":    {},
		"Paris":     {},
		"Liverpool": {},
		"Koln":      {},
		"Lion":      {},
		"Flensburg": {},
		"Bordo":     {},
		"Erfurt":    {},
		"Dresden":   {},
	}
)

type commandOpt func(command *cobra.Command)

// NewCityCommand creates cobra command called city
func NewCityCommand() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "city",
		Short: "Get an info about weather in a specific city",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("WEATHER CITY STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			api := api.New()
			fmt.Println(api.GetWetherByCity(cityName))
			return nil
		},
	}

	setupCityNameFlag()(createCmd)
	return createCmd
}

func setupCityNameFlag() commandOpt {
	const flagName = "name"
	return func(cmd *cobra.Command) {

		cmd.Flags().StringVarP(&cityName, flagName, "n", "", "city name")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
		// register the autocomplete func
		if err := cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			out := make([]string, 0)
			for city := range cityMap {
				if !strings.HasPrefix(city, toComplete) {
					continue
				}
				out = append(out, city)
			}
			return out, cobra.ShellCompDirectiveNoFileComp
		}); err != nil {
			log.Printf("failed to register autocomplete func: %s", err)
		}
	}
}
