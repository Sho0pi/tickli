package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var setProjectCmd = &cobra.Command{
	Use:   "set-project PROJECT_ID",
	Short: "Set default project ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load config")
		}

		cfg.DefaultProjectID = args[0]
		if err := config.Save(cfg); err != nil {
			log.Fatal().Err(err).Msg("Failed to save config")
		}

		log.Info().Str("project_id", cfg.DefaultProjectID).Msg("Default project ID updated")
	},
}

func init() {
	rootCmd.AddCommand(setProjectCmd)
}
