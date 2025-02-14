package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
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

		tasks, err := api.GetTasks(projectID)
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

func getPriorityString(priority int) string {
	switch priority {
	case 0:
		return "None"
	case 1:
		return "Low"
	case 3:
		return "Medium"
	case 5:
		return "High"
	default:
		return "Unknown"
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
