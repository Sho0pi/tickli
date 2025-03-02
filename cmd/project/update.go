package project

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type updateProjectOptions struct {
	projectID   string
	name        string
	color       types.ProjectColor
	viewMode    string
	kind        types.ProjectKind
	interactive bool
}

func newUpdateProjectCommand() *cobra.Command {
	opts := &updateProjectOptions{}
	cmd := &cobra.Command{
		Use:   "update [id/name]",
		Short: "Update an existing project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.projectID = args[0]
			}
			fmt.Println("Updating project:", opts.projectID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.name, "name", "n", "", "New name of the project")
	cmd.Flags().VarP(&opts.color, "color", "c", "New color (hex format, e.g., '#F18181')")
	cmd.Flags().StringVar(&opts.viewMode, "view-mode", "", "New view mode: list, kanban, timeline")
	cmd.Flags().Var(&opts.kind, "kind", "New project kind: TASK, NOTE")
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Update project interactively (prompt for fields)")

	return cmd
}
