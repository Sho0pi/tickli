package task

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
	"slices"
)

type listOptions struct {
	all       bool
	verbose   bool
	priority  types.TaskPriority
	dueDate   string
	tag       string
	projectID string
}

func fetchProjectColor(projectID string) types.ProjectColor {
	project, err := TickliClient.GetProject(projectID)
	if err != nil {
		log.Warn().Err(err).Msg("failed to get project color, using default color")
		return types.DefaultColor
	}
	return project.Color
}

func fetchProjectColorAsync(ctx context.Context, projectID string) <-chan types.ProjectColor {
	colorChan := make(chan types.ProjectColor, 1)

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
	tasks []*types.Task
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

func filterTasks(tasks []*types.Task, opts *listOptions) []*types.Task {
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
		Short:   "List tasks in a project",
		Long:    "List tasks in the current project. By default, this only lists open issues.",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID != "" {
				projectID = opts.projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			colorChan := fetchProjectColorAsync(ctx, projectID)
			taskChan := fetchAndFilterTasksAsync(ctx, projectID, opts)

			// Wait for both operations to complete
			var projectColor types.ProjectColor
			var filteredTasks []*types.Task

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
				projectColor = types.DefaultColor
			case color, ok := <-colorChan:
				if !ok {
					projectColor = types.DefaultColor
				} else {
					projectColor = color
				}
			}

			task, err := utils.FuzzySelectTask(filteredTasks, projectColor, "")
			if err != nil {
				log.Fatal().Err(err).Msg("failed to select task")
			}

			fmt.Println(utils.GetTaskDescription(task, projectColor))
			return nil
		},
	}
	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Project ID for the action scope (default is current project)")
	cmd.Flags().BoolVarP(&opts.all, "all", "a", false, "Include completed tasks")
	cmd.Flags().StringVarP(&opts.tag, "tag", "t", "", "Tags to add to the task")
	cmd.Flags().VarP(&opts.priority, "priority", "p", "Filter by priority (none, low, medium, high)")
	_ = cmd.RegisterFlagCompletionFunc("priority", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"none", "low", "medium", "high"}, cobra.ShellCompDirectiveDefault
	})
	cmd.Flags().StringVar(&opts.dueDate, "due", "", "Filter by due date (today, tomorrow, this-week, overdue)")
	cmd.Flags().BoolVarP(&opts.verbose, "verbose", "v", false, "Show additional details for each task")

	return cmd
}
func Filter(tasks []*types.Task, predicate func(task types.Task) bool) []*types.Task {
	var result []*types.Task
	for _, task := range tasks {
		if predicate(*task) {
			result = append(result, task)
		}
	}
	return result
}
