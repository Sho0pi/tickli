package project

import (
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type updateProjectOptions struct {
	projectID   string
	name        string
	color       types.ProjectColor
	viewMode    types.ViewMode
	kind        types.ProjectKind
	interactive bool
}

func newUpdateProjectCommand() *cobra.Command {
	opts := &updateProjectOptions{}
	cmd := &cobra.Command{
		Use:   "update <project-id>",
		Short: "Update an existing project's properties",
		Long: `Modify the properties of an existing project.
    
You can update a project's name, color, view mode, or kind.
Changes only the properties you specify - others remain unchanged.`,
		Example: `  # Update project name
  tickli project update abc123def456 -n "New Project Name"
  
  # Update multiple properties
  tickli project update abc123def456 --name "Work Tasks" --color "#00FF00" --view-mode kanban
  
  # Update interactively
  tickli project update abc123def456 -i`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Warn().Msg("Update project is not implemented yet")
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.name, "name", "n", "", "Change the project name")
	cmd.Flags().VarP(&opts.color, "color", "c", "Change the project color (hex format, e.g., '#F18181')")
	cmd.Flags().Var(&opts.viewMode, "view-mode", "Change how tasks are displayed: list, kanban, or timeline")
	_ = cmd.RegisterFlagCompletionFunc("view-mode", types.RegisterViewModeCompletions)
	cmd.Flags().Var(&opts.kind, "kind", "Change project type: TASK or NOTE")
	_ = cmd.RegisterFlagCompletionFunc("kind", types.RegisterProjectKindCompletions)
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Update project by answering prompts instead of using flags")

	return cmd
}
