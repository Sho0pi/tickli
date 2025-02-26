package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var TickliClient *api.Client

var rootCmd = &cobra.Command{
	Use:   "tickli",
	Short: "TickTick CLI - A modern command line interface for TickTick",
	Long: `tickli is a CLI tool that helps you manage your TickTick tasks from the command line.
Complete documentation is available at https://github.com/sho0pi/tickli`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Skip initialization check for init command
		if cmd.Name() != "init" && cmd.Name() != "reset" && cmd.Name() != "help" {
			token, err := config.LoadToken()
			if err != nil || token == "" {
				log.Fatal().Msg("Please run 'tickli init' first")
			}

			// Init the TickliClient
			TickliClient = api.NewClient(token)
		}
	},
}

func Execute() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05",
		FormatFieldName: func(i interface{}) string {
			return i.(string) + ":"
		},
		FormatFieldValue: func(i interface{}) string {
			return "'" + i.(string) + "'"
		},
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
