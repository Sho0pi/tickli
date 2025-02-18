package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [project-name]",
	Aliases: []string{"ls"},
	Short:   "List all tasks in a project",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectID string
		if len(args) != 0 {
			projectID = args[0]
		} else {
			cfg, err := config.Load()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to load config")
			}

			projectID = cfg.DefaultProjectID
		}

		tasks, err := TickliClient.ListTasks(projectID)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get tasks")
		}

		task, err := utils.FuzzySelectTask(tasks, "", "")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to select task")
		}

		fmt.Println(task)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
