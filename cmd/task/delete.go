package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/completion"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	projectID string
	taskID    string
	force     bool
}

func newDeleteCommand(client *api.Client) *cobra.Command {
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
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completion.TaskIDs(projectID),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
			opts.taskID = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.force {
				var confirm string
				fmt.Printf("Are you sure you want to delete the task %s? (y/N): ", opts.taskID)
				fmt.Scanln(&confirm)
				if confirm != "y" && confirm != "Y" {
					fmt.Println("Deletion aborted")
					return nil
				}
			}

			err := client.DeleteTask(opts.projectID, opts.taskID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to delete task %s", opts.taskID))
			}

			fmt.Printf("Task %s deleted\n", opts.taskID)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation prompt and delete immediately")

	return cmd
}
