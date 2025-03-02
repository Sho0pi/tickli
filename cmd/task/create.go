package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

type createOptions struct {
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

	projectID string
}

func newCreateCommand() *cobra.Command {
	opts := &createOptions{}
	cmd := &cobra.Command{
		Short:   "Create a new task",
		Use:     "create",
		Aliases: []string{"add", "a"},
		Long:    "Create a new task with an optional title. If the title is not provided, you will be prompted to enter one.",
		Args:    cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID != "" {
				projectID = opts.projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			task := &types.Task{
				ProjectID: projectID,
				Title:     opts.title,
				Content:   opts.content,
				Desc:      opts.description,

				Priority: opts.priority,
				Tags:     opts.tags,

				IsAllDay: opts.allDay,
			}

			task, err := TickliClient.CreateTask(task)
			if err != nil {
				return errors.Wrap(err, "failed to create task")
			}

			fmt.Println(task.ID)
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Project ID for the action scope (default is current project)")
	cmd.Flags().StringVarP(&opts.title, "title", "t", "", "Task title")
	cmd.MarkFlagRequired("title")
	cmd.Flags().StringVarP(&opts.content, "content", "c", "", "Additional content or description")
	cmd.Flags().StringVarP(&opts.description, "desc", "d", "", "Description (for checklist)")
	cmd.Flags().MarkDeprecated("desc", "please use --content")
	cmd.Flags().BoolVarP(&opts.allDay, "all-day", "a", false, "Mark as an all-day task")
	cmd.Flags().StringVar(&opts.startDate, "start", "", "Start date/time in ISO format (e.g., '2025-02-18T15:04:05Z')")
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "Due date/time in ISO format (e.g., '2025-02-18T18:00:00Z')")
	cmd.Flags().StringVar(&opts.date, "date", "", "Date range with natural language processing (e.g., 'today', 'next week')")
	cmd.Flags().StringVar(&opts.timeZone, "tz", "", "Timezone (e.g., 'America/Los_Angeles')")
	cmd.Flags().StringSliceVar(&opts.reminders, "reminders", []string{}, "List of reminder triggers (e.g., '9h', '0s')")
	cmd.Flags().StringSliceVar(&opts.tags, "tags", []string{}, "Tags to add to the task")
	cmd.Flags().StringVar(&opts.repeat, "repeat", "", "Recurring rule (e.g., 'daily', 'weekly on monday')")
	cmd.Flags().VarP(&opts.priority, "priority", "p", "Priority level (none, low, medium, high) (default 'none')")
	_ = cmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"none", "low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Create task interactively (prompt for fields)")

	return cmd
}
