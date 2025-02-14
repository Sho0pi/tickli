package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tickli",
	Short: "TickTick CLI - A modern command line interface for TickTick",
	Long: `tickli is a CLI tool that helps you manage your TickTick tasks from the command line.
Complete documentation is available at https://github.com/sho0pi/tickli`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Skip initialization check for init command
		if cmd.Name() != "init" {
			token, err := config.LoadToken()
			if err != nil || token == "" {
				log.Fatal().Msg("Please run 'tickli init' first")
			}
		}
	},
}

func Execute() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
