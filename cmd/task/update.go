package task

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/types/project"
	"github.com/sho0pi/tickli/internal/types/task"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	projectID string
	taskID    string

	title       string
	content     string
	description string
	priority    task.Priority
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
		Use:   "update <task-id>",
		Short: "Modify an existing task's properties",
		Long: `Update any property of an existing task identified by its ID.
    
Changes only the properties you specify - others remain unchanged.
This command allows modifying title, content, priority, dates, and more.`,
		Example: `  # Update a task's title
  tickli task update abc123def456 -t "New title"
  
  # Update priority and add content
  tickli task update abc123def456 -p high -c "Additional details here"
  
  # Change due date
  tickli task update abc123def456 --due "next Friday 5pm"
  
  # Update interactively
  tickli task update abc123def456 -i`,
		Args: cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
			opts.taskID = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := TickliClient.GetTask(opts.projectID, opts.taskID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to get t with ID %s", opts.taskID))
			}
			if cmd.Flags().Changed("title") {
				t.Title = opts.title
			}
			if cmd.Flags().Changed("content") {
				t.Content = opts.content
			}
			if cmd.Flags().Changed("desc") {
				t.Desc = opts.description
			}
			if cmd.Flags().Changed("priority") {
				t.Priority = opts.priority
			}
			if cmd.Flags().Changed("tags") {
				t.Tags = opts.tags
			}
			if cmd.Flags().Changed("all-day") {
				t.IsAllDay = opts.allDay
			}
			t, err = TickliClient.UpdateTask(t)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to update t %s", opts.taskID))
			}
			fmt.Printf("Task %s updated successfully\n", t.ID)
			fmt.Println(utils.GetTaskDescription(t, project.DefaultColor))
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.title, "title", "t", "", "Change the task title")
	cmd.Flags().StringVarP(&opts.content, "content", "c", "", "Change or add content/description")
	cmd.Flags().StringVarP(&opts.description, "desc", "d", "", "New description (for checklist)")
	cmd.Flags().MarkDeprecated("desc", "please use --content")
	cmd.Flags().BoolVarP(&opts.allDay, "all-day", "a", false, "Toggle all-day status for the task")
	cmd.Flags().StringVar(&opts.startDate, "start", "", "Change when the task begins")
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "Change when the task is due")
	cmd.Flags().StringVar(&opts.timeZone, "timezone", "", "Change timezone for date calculations")
	cmd.Flags().StringSliceVar(&opts.reminders, "reminders", []string{}, "Set reminders (e.g., '10m', '1h before')")
	cmd.Flags().StringVar(&opts.repeat, "repeat", "", "New recurring rule (e.g., 'daily', 'weekly on monday')")
	cmd.Flags().Var(&opts.priority, "priority", "Change task importance: none, low, medium, high")
	_ = cmd.RegisterFlagCompletionFunc("priority", task.PriorityCompletionFunc)
	cmd.Flags().BoolVarP(&opts.interactive, "interactive", "i", false, "Update task by answering prompts")

	return cmd
}
