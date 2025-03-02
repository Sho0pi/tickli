package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"strings"
)

type listOptions struct {
	filter string
}

func filterProjectByName(projects []*types.Project, name string) ([]*types.Project, error) {
	var matched []*types.Project
	nameLower := strings.ToLower(name)
	for i := range projects {
		if strings.Contains(strings.ToLower(projects[i].Name), nameLower) {
			matched = append(matched, projects[i])
		}
	}
	if len(matched) == 0 {
		return nil, fmt.Errorf("no project found with name '%s'", name)
	}
	return matched, nil
}

func newListCommand() *cobra.Command {
	opts := &listOptions{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Show all available projects without setting them as default",
		Long: `Show all available projects without setting any of them as the default. 
Lists each project with its details.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := TickliClient.ListProjects()
			if err != nil {
				return errors.Wrap(err, "failed to fetch projects")
			}

			projects, err = filterProjectByName(projects, opts.filter)
			if err != nil {
				return err
			}

			project, err := utils.FuzzySelectProject(projects, "")
			if err != nil {
				return errors.Wrap(err, "failed to select project")
			}

			fmt.Println(fmt.Sprintf("(%s) %s", project.ID, project.Name))
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.filter, "filter", "f", "", "Filter projects by name")

	return cmd
}
