package task

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	taskID string

	title       string
	content     string
	description string
	priority    types.TaskPriority
	tags        []string

	// time specific vars
	allDay    bool
	date      string
	startDate string
	dueDate   string
	timeZone  string

	// reminders and repeat are more advanced features not implemented yet
	reminders []string
	repeat    string

	// interactive indicates if you should prompt to get title and content
	interactive bool
}

func newUpdateCommand() *cobra.Command {
	opts := &updateOptions{}
	cmd := &cobra.Command{
		Use:   "update [task-id]",
		Short: "Update an existing task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.taskID = args[0]
			fmt.Println("Updating task:", opts.taskID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.title, "title", "t", "", "New task title")
	cmd.Flags().StringVarP(&opts.content, "content", "c", "", "New additional content or description")
	cmd.Flags().StringVarP(&opts.description, "desc", "d", "", "New description (for checklist)")
	cmd.Flags().BoolVarP(&opts.allDay, "all-day", "a", false, "Mark as an all-day task")
	cmd.Flags().StringVar(&opts.startDate, "start", "", "New start date/time (natural language: 'today 3pm', 'tomorrow', etc)")
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "New due date/time (natural language: 'today 6pm', 'next friday', etc)")
	cmd.Flags().StringVar(&opts.timeZone, "timezone", "", "New timezone (e.g., 'America/Los_Angeles')")
	cmd.Flags().StringSliceVar(&opts.reminders, "reminders", []string{}, "New list of reminder triggers (e.g., '9h', '0s')")
	cmd.Flags().StringVar(&opts.repeat, "repeat", "", "New recurring rule (e.g., 'daily', 'weekly on monday')")
	cmd.Flags().Var(&opts.priority, "priority", "New priority level: none, low, medium, high")
	_ = cmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"none", "low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Update task interactively (prompt for fields)")

	return cmd
}
