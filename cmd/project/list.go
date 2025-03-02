package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "show"},
		Short:   "Show all available projects without setting them as default",
		Long: `Show all available projects without setting any of them as the default. 
Lists each project with its details.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := TickliClient.ListProjects()
			if err != nil {
				return errors.Wrap(err, "failed to fetch projects")
			}

			project, err := utils.FuzzySelectProject(projects, "")
			if err != nil {
				return errors.Wrap(err, "failed to select project")
			}

			fmt.Println(fmt.Sprintf("(%s) %s", project.ID, project.Name))
			return nil
		},
	}

	return cmd
}
