package completion

import (
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type ProjectsProvider interface {
	ListProjects() ([]types.Project, error)
}

func loadClient() (*api.Client, error) {
	token, err := config.LoadToken()
	if err != nil || token == "" {
		return nil, err
	}

	client := api.NewClient(token)
	return client, nil
}

func ProjectNames(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	client, err := loadClient()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	projects, err := client.ListProjects()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return ProjectCompletions(projects, toComplete), cobra.ShellCompDirectiveNoFileComp

}

func ProjectCompletions(projects []types.Project, toComplete string) []cobra.Completion {
	var completions []cobra.Completion
	for _, project := range projects {
		completions = append(completions, cobra.CompletionWithDesc(project.ID, project.Name))
	}
	return completions
}
