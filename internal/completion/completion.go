package completion

import (
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type ProjectsProvider interface {
	ListProjects() ([]types.Project, error)
}

func ProjectNames(provider ProjectsProvider) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		projects, err := provider.ListProjects()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return ProjectCompletions(projects, toComplete), cobra.ShellCompDirectiveNoFileComp

	}
}

func ProjectCompletions(projects []types.Project, toComplete string) []cobra.Completion {
	var completions []cobra.Completion
	for _, project := range projects {
		completions = append(completions, cobra.CompletionWithDesc(project.ID, project.Name))
	}
	return completions
}
