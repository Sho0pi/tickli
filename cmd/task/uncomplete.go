package task

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type uncompleteOptions struct {
	projectID string
}

func newUncompleteCommand() *cobra.Command {
	opts := &uncompleteOptions{}
	cmd := &cobra.Command{
		Use:   "uncomplete <task-id>",
		Short: "Mark a completed task as active again",
		Long: `Change a task's status from completed back to active.
    
This command can be used to reactivate tasks that were previously completed
but need to be worked on again.`,
		Example: `  # Reactivate a completed task
  tickli task uncomplete abc123def456
  
  # Reactivate a task in a specific project
  tickli task uncomplete abc123def456 --project-id xyz789`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Warn().Msg("uncomplete command not implemented yet")
		},
	}

	return cmd
}
