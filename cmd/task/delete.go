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
		Use:     "delete <task-id>",
		Aliases: []string{"rm", "remove"},
		Short:   "Remove a task permanently",
		Long: `Delete a task completely from your TickTick account.
    
This operation cannot be undone. By default, you will be asked to confirm
the deletion unless the --force flag is used.`,
		Example: `  # Delete with confirmation prompt
  tickli task delete abc123def456
  
  # Force delete without confirmation
  tickli task delete abc123def456 --force
  
  # Delete from specific project
  tickli task delete abc123def456 --project-id xyz789`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
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

	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation prompt and delete immediately")

	return cmd
}
