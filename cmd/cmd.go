package cmd

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/cmd/project"
	"github.com/sho0pi/tickli/cmd/subtask"
	"github.com/sho0pi/tickli/cmd/task"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func NewTickliCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tickli",
		Short: "TickTick CLI - A modern command line interface for TickTick",
		Long: `tickli is a CLI tool that helps you manage your TickTick tasks from the command line.
Complete documentation is available at https://github.com/sho0pi/tickli`,
		SilenceErrors: true,
		SilenceUsage:  false,
	}
	cmd.AddCommand(
		NewInitCommand(),
		NewResetCommand(),
		NewVersionCommand(),
		task.NewTaskCommand(),
		project.NewProjectCommand(),
		subtask.NewSubtaskCommand(),
	)

	return cmd
}

func Execute() {
	cmd := NewTickliCommand()
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

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal().Err(err).Msg("Failed to execute command")
	}
}
