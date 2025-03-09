package project

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

// NewProjectCommand returns a cobra command for `project` subcommands
func NewProjectCommand() *cobra.Command {
	var client api.Client
	cmd := &cobra.Command{
		Use:     "project",
		Short:   "Work with TickTick projects.",
		Aliases: []string{"list"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			token, err := config.LoadToken()
			if err != nil {
				log.Fatal().Err(err).Msg("Please run 'tickli init' first")
			}
			client = *api.NewClient(token)
			return nil
		},
	}

	cmd.AddCommand(
		newListCommand(&client),
		newCreateProjectCommand(&client),
		newUpdateProjectCommand(&client),
		newUseProjectCmd(&client),
		newShowCommand(&client),
		newDeleteCommand(&client),
	)

	return cmd
}
