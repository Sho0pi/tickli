package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var (
	projectID string
)

func NewTaskCommand() *cobra.Command {
	var client api.Client
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Work with TickTick tasks",
		Long: `Create, view, update, and manage tasks in your TickTick projects.
    
All task commands operate on the current active project by default.
You can change the current project with 'tickli project use' or
specify a different project with the --project-id flag.`,
		Example: `  # List all tasks in current project
  tickli task list
  
  # Create a new task
  tickli task create -t "Submit quarterly report"
  
  # Complete a task
  tickli task complete abc123def456`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			token, err := config.LoadToken()
			if err != nil {
				log.Fatal().Err(err).Msg("Please run 'tickli init' first")
			}
			client = *api.NewClient(token)
			if projectID == "" {
				cfg, err := config.Load()
				if err != nil {
					return errors.Wrap(err, "failed to load config")
				}
				projectID = cfg.DefaultProjectID
			}
			return nil
		},
	}

	cmd.AddCommand(
		newCompleteCmd(&client),
		newDeleteCommand(&client),
		newShowCommand(&client),
		newCreateCommand(&client),
		newListCommand(&client),
		newUncompleteCommand(&client),
		newUpdateCommand(&client),
	)

	RegisterProjectOverride(cmd)

	return cmd
}

func loadClient() *api.Client {
	token, err := config.LoadToken()
	if err != nil || token == "" {
		log.Fatal().Msg("Please run 'tickli init' first")
	}

	// Init the TickliClient
	client := api.NewClient(token)
	return client
}

func RegisterProjectOverride(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&projectID, "project-id", "P", "", "select another project")

	_ = cmd.RegisterFlagCompletionFunc("project-id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		client := loadClient()
		var cacheCompletions []string
		if client != nil {
			projects, err := client.ListProjects()
			if err == nil {
				for i := range projects {
					cacheCompletions = append(cacheCompletions, fmt.Sprintf("%s\t%s", projects[i].ID, projects[i].Name))
				}
				return cacheCompletions, cobra.ShellCompDirectiveNoFileComp
			}
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
}
