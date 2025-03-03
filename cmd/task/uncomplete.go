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
		Use: "uncomplete [task-id]",
		Run: func(cmd *cobra.Command, args []string) {
			log.Warn().Msg("uncomplete command not implemented yet")
		},
	}

	cmd.Flags().StringVarP(&opts.projectID, "project-id", "i", "", "Project ID containing the task (default is current project)")
	return cmd
}
