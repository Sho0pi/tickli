package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/completion"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	projectID string
	force     bool
}

func newDeleteCommand(client *api.Client) *cobra.Command {
	opts := deleteOptions{}
	cmd := &cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete an existing project",
		Long: `Permanently delete a project by its ID.
    
This operation cannot be undone. By default, you will be asked to confirm
the deletion unless the --force flag is used.`,
		Example: `  # Delete with confirmation prompt
  tickli project delete abc123def456
  
  # Force delete without confirmation
  tickli project delete abc123def456 --force`,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completion.ProjectNames,
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.force {
				var confirm string
				fmt.Printf("Are you sure you want to delete the project %s? (y/N): ", opts.projectID)
				fmt.Scanln(&confirm)
				if confirm != "y" && confirm != "Y" {
					fmt.Println("Deletion aborted")
					return nil
				}
			}

			err := client.DeleteProject(opts.projectID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to delete project %s", opts.projectID))
			}

			fmt.Println("Deleting project:", args[0])

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation prompt and delete immediately")

	return cmd
}
