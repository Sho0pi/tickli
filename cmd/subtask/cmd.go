package subtask

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

func NewSubtaskCommand() *cobra.Command {
	var client api.Client
	cmd := &cobra.Command{
		Use:   "subtask",
		Short: "subtask commands",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			client = utils.LoadClient()
			// Put here to avoid runtime error
			log.Info().Interface("client", client).Msg("subtask commands")
			return nil
		},
	}

	return cmd
}
