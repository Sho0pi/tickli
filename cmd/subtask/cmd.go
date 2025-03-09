package subtask

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

func NewSubtaskCommand() *cobra.Command {
	var client api.Client
	cmd := &cobra.Command{
		Use:   "subtask",
		Short: "subtask commands",
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			token, err := config.LoadToken()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to load token")
			}
			client = *api.NewClient(token)
			// Put here to avoid runtime error
			log.Info().Interface("client", client).Msg("subtask commands")
			return nil
		},
	}

	return cmd
}
