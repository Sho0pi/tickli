package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

const (
	inboxID   = "inbox"
	inboxName = "inbox"
)

var noValidate bool

var useIDCmd = &cobra.Command{
	Use:   "use-id <project-id>",
	Short: "Switch to a project using its ID",
	Long: `Switch to a project using its exact ID.
This command requires an exact match and will fail if the ID doesn't exist.

Example:
  tickli use-id inbox   # Switch to inbox
  tickli use-id abc123   # Switch using exact ID match
  tickli use-id abc123 --no-noValidate   # Switch without validation`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

		// Validate if the flag is set
		if !noValidate && projectID != inboxID {
			projects, err := api.GetProjects()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch projects")
			}

			// Check if the project ID exists
			found := false
			for _, project := range projects {
				if project.ID == projectID {
					found = true
					break
				}
			}

			if !found {
				log.Fatal().Str("project_id", projectID).Msg("Project ID not found")
			}
		}
		// Save the project ID in config
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to load config")
		}

		cfg.DefaultProjectID = projectID
		if err := config.Save(cfg); err != nil {
			log.Fatal().Err(err).Msg("failed to save config")
		}

		log.Info().Str("project_id", cfg.DefaultProjectID).Msg("Switched to project")

	},
}

var useCmd = &cobra.Command{
	Use:   "use [project-name]",
	Short: "Switch to a project using name or interactive selection",
	Long: `Switch to a project using its name or interactive selection.
If no name is provided, opens an interactive fuzzy finder.
Special case: 'use inbox' switches to the inbox project.

Examples:
  tickli use              # Interactive fuzzy finder
  tickli use inbox       # Switch to inbox
  tickli use "Work"      # Search by name`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		if len(args) > 0 {
			projectName = args[0]
		}

		projects, err := api.GetProjects()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to fetch projects")
		}

		selectedProject, err := utils.FuzzySelectProject(projects, projectName)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to select project")
		}

		fmt.Println(selectedProject)
		fmt.Println(projectName)
		//matchingProjects := []api.Project{}
		//for _, project := range projects {
		//	if project.Name == projectName {
		//	}
		//}

	},
}

func init() {
	useIDCmd.Flags().BoolVar(&noValidate, "no-validate", false, "Skip project ID validation before switching")

	rootCmd.AddCommand(useIDCmd)
	rootCmd.AddCommand(useCmd)
}
