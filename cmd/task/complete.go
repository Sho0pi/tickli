package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type completeOptions struct {
	projectID string
}

func newCompleteCmd() *cobra.Command {
	opts := &completeOptions{}
	cmd := &cobra.Command{
		Use:   "complete [task-id]",
		Short: "Complete a task",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID == "" {
				opts.projectID = projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID := args[0]
			err := TickliClient.CompleteTask(opts.projectID, taskID)
			if err != nil {
				return errors.Wrap(err, "failed to complete task")
			}

			fmt.Printf("Task %s completed\n", taskID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.projectID, "project-id", "i", "", "Project ID containing the task (default is current project)")
	return cmd
}
