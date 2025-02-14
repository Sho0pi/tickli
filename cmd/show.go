package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all available projects without setting them as default",
	Long: `Show all available projects without setting any of them as the default. 
Lists each project with its details.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := api.GetProjects()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to fetch projects")
		}

		project, err := utils.FuzzySelectProject(projects, "")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to select project")
		}

		fmt.Println(fmt.Sprintf("(%s) %s", project.ID, project.Name))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
