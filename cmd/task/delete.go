package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	projectID string
	force     bool
}

func newDeleteCommand() *cobra.Command {
	opts := &deleteOptions{}
	cmd := &cobra.Command{
		Use:     "delete [task-id]",
		Aliases: []string{"rm", "remove"},
		Short:   "Delete a task",
		Args:    cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID == "" {
				opts.projectID = projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID := args[0]
			if !opts.force {
				var confirm string
				fmt.Printf("Are you sure you want to delete the task %s? (y/N): ", taskID)
				fmt.Scanln(&confirm)
				if confirm != "y" && confirm != "Y" {
					fmt.Println("Deletion aborted")
					return nil
				}
			}

			err := TickliClient.DeleteTask(opts.projectID, taskID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to delete task %s", taskID))
			}

			fmt.Printf("Task %s deleted\n", taskID)

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.projectID, "project-id", "i", "", "Project ID containing the task (default is current project)")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Force deletion without confirmation")

	return cmd
}
