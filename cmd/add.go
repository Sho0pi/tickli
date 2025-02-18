package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/spf13/cobra"
)

var (
	addProjectID string
	addAllDay    bool
	addContent   string
	addPriority  types.TaskPriority
	addDate      string
	addStartDate string
	addDueDate   string
	addTags      []string
)

var (
	title string
)

var addCmd = &cobra.Command{
	Use:     "add [task title]",
	Aliases: []string{"a"},
	Short:   "Add a new task to TickTick",
	Long: `Create a new task in TickTick. You can specify additional properties such as:
- Project ID
- Task content
- Priority level (none, low, medium, high)
- Start and due dates (parsed using natural language)
- All-day flag`,
	Example: `
# Add a task with only a title
tickli add "Buy groceries"

# Add a task with content
tickli add "Finish report" --content "Include Q4 financials"

# Add a task with a specific project ID
tickli add "Plan trip" --project-id "6226ff9877acee87727f6bca"

# Add a high-priority task
tickli add "Urgent Fix" --priority high

# Add a task with a natural date format
tickli add "Doctor appointment" --date "next Monday at 10am"
	`,
	Args: cobra.MaximumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			title = args[0]
		} else {
			//TODO: add support to read from stdin
			//stat, _ := os.Stdin.Stat()
			//if (stat.Mode() & os.ModeCharDevice) == 0 {
			//	scanner := bufio.NewScanner(os.Stdin)
			//	if scanner.Scan() {
			//		title = scanner.Text()
			//	}
			//	if err := scanner.Err(); err != nil {
			//		return errors.Wrap(err, "reading standard input")
			//	}
			//}
		}

		if title == "" {
			return errors.New("must specify a task title")
		}

		if addProjectID == "" {
			cfg, err := config.Load()
			if err != nil {
				return errors.Wrap(err, "failed to load config")
			}
			addProjectID = cfg.DefaultProjectID
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		task := &types.Task{
			ProjectID: addProjectID,
			Title:     title,
			IsAllDay:  addAllDay,
			Content:   addContent,
			Priority:  addPriority,
			Tags:      addTags,
		}

		fmt.Println(task)

		newTask, err := TickliClient.CreateTask(task)
		if err != nil {
			return errors.Wrap(err, "failed to create task")
		}

		fmt.Println()
		fmt.Println(newTask)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&addProjectID, "project-id", "", "Project ID for the task (default is from configuration)")
	addCmd.Flags().BoolVarP(&addAllDay, "all-day", "a", false, "Mark task as an all-day task")
	addCmd.Flags().StringVarP(&addContent, "content", "c", "", "Additional content or description for the task")

	addCmd.Flags().VarP(&addPriority, "priority", "p", "Task priority (none, low, medium, high)")
	_ = addCmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"high", "medium", "low", "none"}, cobra.ShellCompDirectiveNoFileComp
	})

	addCmd.Flags().StringVarP(&addDate, "date", "d", "", "Task date range in natural language (e.g., 'today to tomorrow', 'next month')")
	addCmd.Flags().StringVar(&addStartDate, "start-date", "", "Start date in natural language (e.g., 'tomorrow', '3 feb 2022')")
	addCmd.Flags().StringVar(&addDueDate, "due-date", "", "Due date in natural language (e.g., 'next friday', '6 oct 2022')")
	addCmd.Flags().StringSliceVarP(&addTags, "tags", "t", []string{}, "Tags to add to the task")

}
