package project

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type showOptions struct {
	projectID string
	withTasks bool
	output    types.OutputFormat
}

func newShowCommand() *cobra.Command {
	opts := &showOptions{
		output: types.OutputSimple,
	}

	cmd := &cobra.Command{
		Use:     "show [id]",
		Aliases: []string{"info"},
		Short:   "Show details of a project",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Determine which project to show
			if len(args) > 0 {
				opts.projectID = args[0]
			} else {
				cfg, err := config.Load()
				if err != nil {
					return errors.Wrap(err, "failed to load config")
				}
				opts.projectID = cfg.DefaultProjectID
			}

			if opts.withTasks {
				projectData, err := TickliClient.GetProjectWithTasks(opts.projectID)
				if err != nil {
					return errors.Wrap(err, "failed to get project data")
				}
				switch opts.output {
				case types.OutputSimple:
					fmt.Println(utils.GetProjectDescription(&projectData.Project))
					for _, task := range projectData.Tasks {
						fmt.Println(utils.GetTaskDescription(&task, projectData.Project.Color))
					}
				case types.OutputJSON:
					jsonData, err := json.MarshalIndent(projectData, "", "  ")
					if err != nil {
						return errors.Wrap(err, "failed to marshal output")
					}
					fmt.Println(string(jsonData))
				}
			} else {
				project, err := TickliClient.GetProject(opts.projectID)
				if err != nil {
					return fmt.Errorf("project %s not found", opts.projectID)
				}
				switch opts.output {
				case types.OutputJSON:
					jsonData, err := json.MarshalIndent(project, "", "  ")
					if err != nil {
						return errors.Wrap(err, "failed to marshal output")
					}
					fmt.Println(string(jsonData))
				case types.OutputSimple:
					fmt.Println(utils.GetProjectDescription(project))
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&opts.withTasks, "with-tasks", false, "Include tasks in the output")
	cmd.Flags().VarP(&opts.output, "output", "o", "Output format. One of: simple, json")
	return cmd
}
