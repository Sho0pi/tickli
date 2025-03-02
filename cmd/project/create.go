package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type createProjectOptions struct {
	name        string
	color       types.ProjectColor
	viewMode    types.ViewMode
	kind        types.ProjectKind
	interactive bool
}

func newCreateProjectCommand() *cobra.Command {
	opts := &createProjectOptions{
		kind:     types.KindTask,
		viewMode: types.ViewModeList,
		color:    types.DefaultColor,
	}
	cmd := &cobra.Command{
		Use:   "create [project-name]",
		Short: "Create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			project := &types.Project{
				Name:     opts.name,
				Color:    opts.color,
				ViewMode: opts.viewMode,
				Kind:     opts.kind,
			}

			project, err := TickliClient.CreateProject(project)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to create project %s", project.Name))
			}

			fmt.Println(utils.GetProjectDescription(project))
			fmt.Println(project.ID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.name, "name", "n", "", "Name of the project (required if not using interactive mode)")
	cmd.MarkFlagRequired("name")
	cmd.Flags().VarP(&opts.color, "color", "c", "Color of the project (hex format, e.g., '#F18181')")
	cmd.Flags().Var(&opts.viewMode, "view-mode", "View mode: list, kanban, timeline (default 'list')")
	cmd.Flags().Var(&opts.kind, "kind", "Project kind: TASK, NOTE (default 'TASK')")
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Create project interactively (prompt for fields)")

	return cmd
}
