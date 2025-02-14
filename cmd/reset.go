package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset tickli authentication",
	Long: `Reset tickli by removing the current access token and re-running the initialization process.
This is useful if you need to reauthenticate with TickTick.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.DeleteToken(); err != nil {
			log.Fatal().Err(err).Msg("Failed to remove access token")
		}

		log.Info().Msg("Successfully removed access token. Running initialization...")
		initCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
