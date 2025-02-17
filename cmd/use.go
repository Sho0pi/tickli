package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"strings"
)

var (
	projectName string
	projectID   string
)

type projectSelector struct {
	projects []types.Project
}

func newProjectSelector(projects []types.Project) *projectSelector {
	return &projectSelector{projects: projects}
}

func (ps *projectSelector) byID(id string) (*types.Project, error) {
	for _, p := range ps.projects {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("no project found with ID '%s'", id)
}

func (ps *projectSelector) byName(name string) ([]types.Project, error) {
	var matched []types.Project
	nameLower := strings.ToLower(name)
	for _, p := range ps.projects {
		if strings.Contains(strings.ToLower(p.Name), nameLower) {
			matched = append(matched, p)
		}
	}
	if len(matched) == 0 {
		return nil, fmt.Errorf("no project found with name '%s'", name)
	}
	return matched, nil
}

var useCmd = &cobra.Command{
	Use:   "use [-n name | -i id]",
	Short: "Switch active project context",
	Long: `Switch the active project context for subsequent commands.

This command allows you to change your active project in three ways:
1. Interactive selection (no arguments)
2. Direct selection by project name
3. Direct selection by project ID

The selected project becomes the default context for future commands.`,
	Example: `  # Interactive project selection
  tickli use

  # Switch by partial or full project name
  tickli use -n "My Project"

  # Switch by project ID
  tickli use -i abc123`,
	Args: cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// first get all the available projects.
		projects, err := api.GetProjects()
		if err != nil {
			return errors.Wrap(err, "failed to get projects")
		}

		selector := newProjectSelector(projects)
		var selectedProject types.Project

		switch {
		case projectID != "":
			project, err := selector.byID(projectID)
			if err != nil {
				return err
			}
			selectedProject = *project
		case projectName != "":
			matches, err := selector.byName(projectName)
			if err != nil {
				return err
			}
			if len(matches) == 1 {
				selectedProject = matches[0]
			} else {
				project, err := utils.FuzzySelectProject(matches, projectName)
				if err != nil {
					return errors.Wrap(err, "failed to select project")
				}
				selectedProject = *project
			}
		default:
			project, err := utils.FuzzySelectProject(projects, "")
			if err != nil {
				return errors.Wrap(err, "failed to select project")
			}
			selectedProject = *project
		}

		cfg, err := config.Load()
		if err != nil {
			return errors.Wrap(err, "failed to load config")
		}

		cfg.DefaultProjectID = selectedProject.ID
		if err := config.Save(cfg); err != nil {
			return errors.Wrap(err, "failed to save config")
		}

		log.Info().Str("project_id", cfg.DefaultProjectID).Msg("Default project ID updated")
		return nil
	},
}

func init() {
	useCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project to switch to")
	useCmd.Flags().StringVarP(&projectID, "id", "i", "", "ID of the project to switch to")
	useCmd.MarkFlagsMutuallyExclusive("name", "id")

	rootCmd.AddCommand()
	rootCmd.AddCommand(useCmd)
}
