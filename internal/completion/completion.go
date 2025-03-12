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

func ProjectIDs() cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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
}

func TaskIDs(projectID string) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		client, err := loadClient()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		if projectID == "" {
			cfg, err := config.Load()
			if err != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			projectID = cfg.DefaultProjectID
		}

		tasks, err := client.ListTasks(projectID)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return TaskCompletions(tasks), cobra.ShellCompDirectiveNoFileComp
	}
}

func TaskCompletions(tasks []types.Task) []cobra.Completion {
	var completions []cobra.Completion
	for _, task := range tasks {
		completions = append(completions, cobra.CompletionWithDesc(task.ID, task.Title))
	}
	return completions
}

func ProjectCompletions(projects []types.Project, toComplete string) []cobra.Completion {
	var completions []cobra.Completion
	for _, project := range projects {
		completions = append(completions, cobra.CompletionWithDesc(project.ID, project.Name))
	}
	return completions
}
