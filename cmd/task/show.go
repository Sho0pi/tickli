package task

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type showOptions struct {
	projectID string
	output    types.OutputFormat
}

func newShowCommand() *cobra.Command {
	opts := &showOptions{
		output: types.OutputSimple,
	}
	cmd := &cobra.Command{
		Use:     "show <task-id>",
		Aliases: []string{"info"},
		Short:   "Display detailed information about a task",
		Long: `Show complete information about a specific task identified by its ID.
    
Displays title, content, dates, priority, tags, and other properties.
You can choose between human-readable output or machine-readable JSON.`,
		Example: `  # Show task details in human-readable format
  tickli task show abc123def456
  
  # Show task from specific project
  tickli task show abc123def456 -i xyz789
  
  # Show task details in JSON format
  tickli task show abc123def456 -o json`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID != "" {
				projectID = opts.projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId := args[0]

			task, err := TickliClient.GetTask(projectID, taskId)
			if err != nil {
				return err
			}
			if task.ID != taskId {
				log.Warn().Str("task-id", taskId).Str("project-id", projectID).Msg("task not found")
				return fmt.Errorf("task %s not found for porject %s", taskId, projectID)
			}
			switch opts.output {
			case types.OutputSimple:
				fmt.Println(utils.GetTaskDescription(task, types.DefaultColor))
			case types.OutputJSON:
				jsonData, err := json.MarshalIndent(task, "", "  ")
				if err != nil {
					return errors.Wrap(err, "failed to marshal output")
				}
				fmt.Println(string(jsonData))
			}
			fmt.Println(task.ID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.projectID, "project-id", "i", "", "Project containing the task (if not in current project)")
	cmd.Flags().VarP(&opts.output, "output", "o", "Display format: simple (human-readable) or json (machine-readable)")
	return cmd
}
