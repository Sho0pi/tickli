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
		Use:     "create",
		Aliases: []string{"add", "a"},
		Short:   "Create a new task",
		Long: `Create a new task in the current project or a specified project.
    
You can set various properties including title, content, priority, due date,
and tags. At minimum, a title is required unless using interactive mode.`,
		Example: `  # Create a basic task with just a title
  tickli task create -t "Buy groceries"
  
  # Create a task with priority and due date
  tickli task create -t "Submit report" -p high --due "tomorrow 5pm"
  
  # Create a task in a specific project
  tickli task create -t "Call client" --project-id abc123def456
  
  # Create a task with content and tags
  tickli task create -t "Team meeting" -c "Discuss Q3 roadmap" --tags meeting,work
  
  # Create a task interactively
  tickli task create -i`,
		Args: cobra.NoArgs,
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

	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Create in a specific project instead of the current one")
	cmd.Flags().StringVarP(&opts.title, "title", "t", "", "Title of the task (required)")
	cmd.MarkFlagRequired("title")
	cmd.Flags().StringVarP(&opts.content, "content", "c", "", "Additional details about the task")
	cmd.Flags().StringVarP(&opts.description, "desc", "d", "", "Description (for checklist)")
	cmd.Flags().MarkDeprecated("desc", "please use --content")
	cmd.Flags().BoolVarP(&opts.allDay, "all-day", "a", false, "Set as an all-day task without specific time")
	cmd.Flags().StringVar(&opts.startDate, "start", "", "When the task begins (ISO format: '2025-02-18T15:04:05Z')")
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "When the task is due (ISO format: '2025-02-18T18:00:00Z')")
	cmd.Flags().StringVar(&opts.date, "date", "", "Set date with natural language (e.g., 'today', 'next week')")
	cmd.Flags().StringVar(&opts.timeZone, "tz", "", "Timezone for date calculations (e.g., 'America/Los_Angeles')")
	cmd.Flags().StringSliceVar(&opts.reminders, "reminders", []string{}, "List of reminder triggers (e.g., '9h', '0s')")
	cmd.Flags().StringSliceVar(&opts.tags, "tags", []string{}, "Apply tags to categorize the task (comma-separated)")
	cmd.Flags().StringVar(&opts.repeat, "repeat", "", "Recurring rule (e.g., 'daily', 'weekly on monday')")
	cmd.Flags().VarP(&opts.priority, "priority", "p", "Task importance: none, low, medium, high (default: none)")
	_ = cmd.RegisterFlagCompletionFunc("priority", types.RegisterTaskPriorityCompletions)
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Create task by answering prompts")

	return cmd
}
