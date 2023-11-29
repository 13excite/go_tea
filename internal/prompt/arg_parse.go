package prompt

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// argument represents a single argument for the command
type argument struct {
	name        string
	value       pflag.Value
	description string
}

// uiAnnotationKey is a key for the annotation that indicates whether the flag belongs to the UI
const uiAnnotationKey = "interactive-ui"

// ignoredFlags contains a list of flags that should be ignored
var ignoredFlags = []string{"version", "help", "yes"}

// annotateArgsAsUIRelated adds an annotation that indicates whether the flag belongs to the UI
func annotateArgsAsUIRelated(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden || stringInSlice(flag.Name, ignoredFlags) {
			return
		}
		if flag.Annotations == nil {
			flag.Annotations = map[string][]string{}
		}
		flag.Annotations[uiAnnotationKey] = []string{"true"}
	})
}

// getArgs returns a list of arguments of the command
func getArgs(cmd *cobra.Command) []argument {
	args := []argument{}
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		val, ok := flag.Annotations[uiAnnotationKey]
		if ok && (val[0] == "true") {
			args = append(args, argument{
				name:        flag.Name,
				value:       flag.Value,
				description: flag.Usage,
			})
		}
	})
	return args
}

// IsUserInputFlagsValid checks that each required flag has a value.
// For avoiding cases when user set empty string for required flag.
func IsUserInputFlagsValid(cmd *cobra.Command) bool {
	validInput := true
	cmd.Flags().VisitAll(func(pflag *pflag.Flag) {
		requiredAnnotation, ok := pflag.Annotations[cobra.BashCompOneRequiredFlag]
		if !ok {
			return
		}
		if (requiredAnnotation[0] == "true") && (pflag.Value.String() == "") {
			validInput = false
		}
	})
	return validInput
}

// stringInSlice checks that given string is in given slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
