package project

import (
	"github.com/sho0pi/tickli/internal/api"
	"github.com/spf13/cobra"
)

// NewProjectCommand returns a cobra command for `project` subcommands
func NewProjectCommand(client *api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Work with TickTick projects.",
		Aliases: []string{"list"},
	}
	cmd.AddCommand(
		newListCommand(client),
		newCreateProjectCommand(client),
		newUpdateProjectCommand(client),
		newUseProjectCmd(client),
		newShowCommand(client),
		newDeleteCommand(client),
	)

	return cmd
}
