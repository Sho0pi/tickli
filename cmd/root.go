package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// TODO: think about how to avoid reading the token in almost every api call
var Token string

var rootCmd = &cobra.Command{
	Use:   "tickli",
	Short: "TickTick CLI - A modern command line interface for TickTick",
	Long: `tickli is a CLI tool that helps you manage your TickTick tasks from the command line.
Complete documentation is available at https://github.com/sho0pi/tickli`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Skip initialization check for init command
		if cmd.Name() != "init" {
			Token, err := config.LoadToken()
			if err != nil || Token == "" {
				log.Fatal().Msg("Please run 'tickli init' first")
			}
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
