package project

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
	opts := deleteOptions{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existing project",
		Args:  cobra.NoArgs,
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

			err := TickliClient.DeleteProject(opts.projectID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to delete project %s", opts.projectID))
			}

			fmt.Println("Deleting project:", args[0])

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.projectID, "project-id", "i", "", "Project ID to delete")
	cmd.MarkFlagRequired("project-id")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Force deletion without confirmation")

	return cmd
}
