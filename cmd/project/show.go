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
		Use:     "show [project-id]",
		Aliases: []string{"info"},
		Short:   "Show details of a project",
		Long: `Display detailed information about a specific project.
    
If no project ID is provided, shows the currently active project.
Can include associated tasks and switch between output formats.`,
		Example: `  # Show current project
  tickli project show
  
  # Show specific project
  tickli project show abc123def456
  
  # Show project with all its tasks
  tickli project show --with-tasks
  
  # Output in JSON format
  tickli project show -o json`,
		Args: cobra.MaximumNArgs(1),
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

	cmd.Flags().BoolVar(&opts.withTasks, "with-tasks", false, "Include all tasks belonging to this project")
	cmd.Flags().VarP(&opts.output, "output", "o", "Format for displaying results: simple (human-readable) or json (machine-readable)")
	_ = cmd.RegisterFlagCompletionFunc("output", types.RegisterOutputFormatCompletions)
	return cmd
}
