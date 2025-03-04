package types

import (
	"fmt"
	"github.com/spf13/cobra"
)

type OutputFormat string

const (
	OutputSimple OutputFormat = "simple"
	OutputJSON   OutputFormat = "json"
)

func RegisterOutputFormatCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{string(OutputJSON), string(OutputSimple)}, cobra.ShellCompDirectiveDefault
}

func (o *OutputFormat) Set(value string) error {
	switch OutputFormat(value) {
	case OutputSimple, OutputJSON:
		*o = OutputFormat(value)
	default:
		return fmt.Errorf("invalid output format: %s", value)
	}
	return nil
}

func (o OutputFormat) String() string {
	return string(o)
}

func (o OutputFormat) Type() string {
	return "OutputFormat"
}
