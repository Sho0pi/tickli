package task

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/types/project"
	"github.com/sho0pi/tickli/internal/types/task"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"slices"
)

type listOptions struct {
	all       bool
	verbose   bool
	priority  task.Priority
	dueDate   string
	tag       string
	projectID string
}

func fetchProjectColor(projectID string) project.Color {
	p, err := TickliClient.GetProject(projectID)
	if err != nil {
		log.Warn().Err(err).Msg("failed to get p color, using default color")
		return project.DefaultColor
	}
	return p.Color
}

func fetchProjectColorAsync(ctx context.Context, projectID string) <-chan project.Color {
	colorChan := make(chan project.Color, 1)

	go func() {
		defer close(colorChan)

		select {
		case <-ctx.Done():
			return
		case colorChan <- fetchProjectColor(projectID):
		}
	}()

	return colorChan
}

type taskFilterResult struct {
	tasks []types.Task
	err   error
}

func fetchAndFilterTasksAsync(ctx context.Context, projectID string, opts *listOptions) <-chan taskFilterResult {
	resultChan := make(chan taskFilterResult, 1)

	go func() {
		defer close(resultChan)

		// Fetch tasks
		tasks, err := TickliClient.ListTasks(projectID)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case resultChan <- taskFilterResult{err: err}:
				return
			}
		}

		// Apply filters
		filteredTasks := filterTasks(tasks, opts)

		select {
		case <-ctx.Done():
			return
		case resultChan <- taskFilterResult{filteredTasks, nil}:
		}
	}()

	return resultChan
}

func filterTasks(tasks []types.Task, opts *listOptions) []types.Task {
	// Filter by priority
	tasks = Filter(tasks, func(t types.Task) bool {
		return t.Priority >= opts.priority
	})

	// Filter by tags
	tasks = Filter(tasks, func(t types.Task) bool {
		if opts.tag != "" {
			return slices.Contains(t.Tags, opts.tag)
		}
		return true
	})

	// Filter by completion status
	if !opts.all {
		//	tasks = Filter(tasks, func(t types.Task) bool {
		//		return !t.
		//	})
	}

	// TODO: implement due date filtering
	if opts.dueDate != "" {
		// Future implementation
	}

	return tasks
}

func newListCommand() *cobra.Command {
	opts := &listOptions{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Browse and select from available tasks",
		Long: `Display tasks in the current project or a specified project.
    
By default, only shows incomplete tasks. You can filter tasks by priority,
tags, and due date. Results are displayed in an interactive selector.`,
		Example: `  # List all incomplete tasks in current project
  tickli task list
  
  # List all tasks including completed ones
  tickli task list --all
  
  # List tasks with specific tag
  tickli task list -t important
  
  # List high priority tasks
  tickli task list -p high
  
  # List tasks in specific project
  tickli task list --project-id abc123def456`,
		Args: cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.projectID = projectID
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			colorChan := fetchProjectColorAsync(ctx, projectID)
			taskChan := fetchAndFilterTasksAsync(ctx, projectID, opts)

			// Wait for both operations to complete
			var projectColor project.Color
			var filteredTasks []types.Task

			// Get the task results
			taskResult := <-taskChan
			if taskResult.err != nil {
				cancel() // Cancel the color fetching if task fetching failed
				return taskResult.err
			}
			filteredTasks = taskResult.tasks

			// Get the project color
			select {
			case <-ctx.Done():
				projectColor = project.DefaultColor
			case color, ok := <-colorChan:
				if !ok {
					projectColor = project.DefaultColor
				} else {
					projectColor = color
				}
			}

			t, err := utils.FuzzySelectTask(filteredTasks, projectColor, "")
			if err != nil {
				log.Fatal().Err(err).Msg("failed to select task")
			}

			fmt.Println(utils.GetTaskDescription(t, projectColor))
			return nil
		},
	}
	cmd.Flags().BoolVarP(&opts.all, "all", "a", false, "Include completed tasks in the results")
	cmd.Flags().StringVarP(&opts.tag, "tag", "t", "", "Only show tasks with this specific tag")
	cmd.Flags().VarP(&opts.priority, "priority", "p", "Only show tasks with this priority level or higher")
	_ = cmd.RegisterFlagCompletionFunc("priority", task.PriorityCompletionFunc)
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "Filter by due date (today, tomorrow, this-week, overdue)")
	cmd.Flags().BoolVarP(&opts.verbose, "verbose", "v", false, "Show more details for each task in the list")

	return cmd
}
func Filter(tasks []types.Task, predicate func(task types.Task) bool) []types.Task {
	var result []types.Task
	for _, t := range tasks {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}
