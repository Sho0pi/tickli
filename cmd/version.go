package cmd

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"runtime/debug"
)

// Version is set at build time using -ldflags
var Version = "dev"

func NewVersionCommand() *cobra.Command {
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version of Tickli",
		Long: `Display the current version of the Tickli CLI tool.
This command shows the version number, which can be set at build time 
or automatically detected from build information.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if Version == "dev" {
				if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
					Version = info.Main.Version
				}
			}
			if jsonOut {
				return utils.PrintJSON(map[string]string{"version": Version})
			}
			fmt.Println("tickli version", Version)
			return nil
		},
	}
	utils.AddJSONFlag(cmd, &jsonOut)
	return cmd
}
