package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/completion"
	"github.com/sho0pi/tickli/internal/types/project"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type updateProjectOptions struct {
	projectID   string
	name        string
	color       project.Color
	viewMode    project.ViewMode
	kind        project.Kind
	interactive bool
}

func newUpdateProjectCommand(client *api.Client) *cobra.Command {
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
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completion.ProjectNames,
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := client.GetProject(opts.projectID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to fetch project %s", opts.projectID))
			}
			if cmd.Flags().Changed("name") {
				p.Name = opts.name
			}
			if cmd.Flags().Changed("color") {
				p.Color = opts.color
			}
			if cmd.Flags().Changed("view-mode") {
				p.ViewMode = opts.viewMode
			}
			if cmd.Flags().Changed("kind") {
				p.Kind = opts.kind
			}
			p, err = client.UpdateProject(p)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to update project %s", opts.projectID))
			}
			fmt.Printf("Project %s updated successfully\n", p.ID)
			fmt.Println(utils.GetProjectDescription(p))
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.name, "name", "n", "", "Change the project name")
	cmd.Flags().VarP(&opts.color, "color", "c", "Change the project color (hex format, e.g., '#F18181')")
	_ = cmd.RegisterFlagCompletionFunc("color", project.ColorCompletionFunc)
	cmd.Flags().Var(&opts.viewMode, "view-mode", "Change how tasks are displayed: list, kanban, or timeline")
	_ = cmd.RegisterFlagCompletionFunc("view-mode", project.ViewModeCompletionFunc)
	cmd.Flags().Var(&opts.kind, "kind", "Change project type: TASK or NOTE")
	_ = cmd.RegisterFlagCompletionFunc("kind", project.KindCompletionFunc)
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Update project by answering prompts instead of using flags")

	return cmd
}
