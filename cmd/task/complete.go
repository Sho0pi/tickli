package task

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/spf13/cobra"
)

type completeOptions struct {
	projectID string
	taskID    string
}

func newCompleteCmd(client *api.Client) *cobra.Command {
	opts := &completeOptions{}
	cmd := &cobra.Command{
		Use:   "complete <task-id>",
		Short: "Mark a task as completed",
		Long: `Change a task's status to completed.
    
Takes a task ID and marks it as done. The task remains in the system
but will no longer appear in default listings unless using the --all flag.`,
		Example: `  # Complete a task in current project
  tickli task complete abc123def456
  
  # Complete a task in a specific project
  tickli task complete abc123def456 --project-id xyz789`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
			opts.taskID = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := client.CompleteTask(opts.projectID, opts.taskID)
			if err != nil {
				return errors.Wrap(err, "failed to complete task")
			}

			fmt.Printf("%s Task %s completed\n", color.Green.Sprint("☑"), opts.taskID)
			return nil
		},
	}

	return cmd
}
