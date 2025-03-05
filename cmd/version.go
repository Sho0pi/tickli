package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

// Version is set at build time using -ldflags
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Tickli",
	Long: `Display the current version of the Tickli CLI tool.
This command shows the version number, which can be set at build time 
or automatically detected from build information.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version == "dev" {
			if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
				Version = info.Main.Version
			}
		}
		fmt.Println("tickli version", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
