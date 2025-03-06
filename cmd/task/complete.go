package task

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type completeOptions struct {
	projectID string
}

func newCompleteCmd() *cobra.Command {
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
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID := args[0]
			err := TickliClient.CompleteTask(opts.projectID, taskID)
			if err != nil {
				return errors.Wrap(err, "failed to complete task")
			}

			fmt.Printf("%s Task %s completed\n", color.Green.Sprint("☑"), taskID)
			return nil
		},
	}

	return cmd
}
