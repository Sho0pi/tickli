package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"strings"
)

func findProjectByID(projects []*types.Project, id string) (*types.Project, error) {
	for i := range projects {
		if projects[i].ID == id {
			return projects[i], nil
		}
	}
	return nil, fmt.Errorf("project not found with ID '%s'", id)
}

func findProjectsByName(projects []*types.Project, name string) ([]*types.Project, error) {
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

type useProjectOptions struct {
	projectName string
	projectID   string
}

func newUseProjectCmd() *cobra.Command {
	opts := &useProjectOptions{}
	cmd := &cobra.Command{
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
			projects, err := TickliClient.ListProjects()
			if err != nil {
				return errors.Wrap(err, "could not fetch projects")
			}

			var selectedProject *types.Project

			switch {
			case opts.projectID != "":
				project, err := findProjectByID(projects, opts.projectID)
				if err != nil {
					return err
				}
				selectedProject = project
			case opts.projectName != "":
				matches, err := findProjectsByName(projects, opts.projectName)
				if err != nil {
					return err
				}
				if len(matches) == 1 {
					selectedProject = matches[0]
				} else {
					project, err := utils.FuzzySelectProject(matches, "")
					if err != nil {
						return errors.Wrap(err, "could not select project")
					}
					selectedProject = project
				}
			default:
				project, err := utils.FuzzySelectProject(projects, "")
				if err != nil {
					return errors.Wrap(err, "could not select project")
				}
				selectedProject = project
			}

			cfg, err := config.Load()
			if err != nil {
				return errors.Wrap(err, "could not load config")
			}

			cfg.DefaultProjectID = selectedProject.ID
			if err := config.Save(cfg); err != nil {
				return errors.Wrap(err, "failed to save config")
			}
			log.Info().
				Str("project_id", cfg.DefaultProjectID).
				Str("project_name", selectedProject.Name).
				Msg("Switched to project")
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.projectName, "name", "n", "", "Name of the project to switch to")
	cmd.Flags().StringVarP(&opts.projectID, "id", "i", "", "ID of the project to switch to")
	cmd.MarkFlagsMutuallyExclusive("name", "id")
	return cmd
}
