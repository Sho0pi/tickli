package utils

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// AddJSONFlag registers the shared --json flag on cmd.
// In JSON mode commands print machine-readable output to stdout and never prompt.
func AddJSONFlag(cmd *cobra.Command, target *bool) {
	cmd.Flags().BoolVar(target, "json", false, "Print result as JSON to stdout (non-interactive)")
}

// PrintJSON writes v to stdout as indented JSON.
func PrintJSON(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal output")
	}
	fmt.Println(string(data))
	return nil
}
