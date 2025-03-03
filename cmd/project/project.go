package project

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var TickliClient *api.Client

var Cmd = &cobra.Command{
	Use:     "project",
	Short:   "Work with TickTick projects.",
	Aliases: []string{"list"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		token, err := config.LoadToken()
		if err != nil || token == "" {
			log.Fatal().Msg("Please run 'tickli init' first")
		}

		TickliClient = api.NewClient(token)
		return nil
	},
}

func init() {
	Cmd.AddCommand(newListCommand())
	Cmd.AddCommand(newCreateProjectCommand())
	Cmd.AddCommand(newUpdateProjectCommand())
	Cmd.AddCommand(newUseProjectCmd())
	Cmd.AddCommand(newDeleteCommand())
	Cmd.AddCommand(newShowCommand())
}
