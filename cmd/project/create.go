package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/types/project"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type createProjectOptions struct {
	name        string
	color       project.Color
	viewMode    project.ViewMode
	kind        project.Kind
	interactive bool
}

func newCreateProjectCommand() *cobra.Command {
	opts := &createProjectOptions{
		kind:     project.KindTask,
		viewMode: project.ViewModeList,
		color:    project.DefaultColor,
	}
	cmd := &cobra.Command{
		Use:   "create [project-name]",
		Short: "Create a new project",
		Long: `Create a new TickTick project with customizable properties.
    
You can specify a name, color, view mode, and project type. The command
supports both direct parameter input and interactive mode.`,
		Example: `  # Create a basic task project
  tickli project create "My New Project"
  
  # Create a project with custom properties
  tickli project create -n "Work Tasks" -c "#FF5733" --view-mode kanban --kind TASK
  
  # Create a project in interactive mode
  tickli project create -i`,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := &types.Project{
				Name:     opts.name,
				Color:    opts.color,
				ViewMode: opts.viewMode,
				Kind:     opts.kind,
			}

			p, err := TickliClient.CreateProject(p)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to create project %s", p.Name))
			}

			fmt.Println(utils.GetProjectDescription(*p))
			fmt.Println(p.ID)
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.name, "name", "n", "", "Name of the new project")
	_ = cmd.MarkFlagRequired("name")
	cmd.Flags().VarP(&opts.color, "color", "c", "Color for the project (hex format, e.g., '#F18181')")
	_ = cmd.RegisterFlagCompletionFunc("color", project.ColorCompletionFunc)
	cmd.Flags().Var(&opts.viewMode, "view-mode", "How to display tasks: list, kanban, or timeline (default: list)")
	_ = cmd.RegisterFlagCompletionFunc("view-mode", project.ViewModeCompletionFunc)
	cmd.Flags().Var(&opts.kind, "kind", "Project type: TASK for action items or NOTE for information (default: TASK)")
	_ = cmd.RegisterFlagCompletionFunc("kind", project.KindCompletionFunc)
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Create project by answering prompts instead of using flags")

	return cmd
}
