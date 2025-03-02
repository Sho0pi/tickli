package task

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/types"
	"github.com/sho0pi/tickli/internal/utils"
	"github.com/spf13/cobra"
)

type OutputFormat string

const (
	OutputSimple OutputFormat = "simple"
	OutputJSON   OutputFormat = "json"
)

func (o *OutputFormat) Set(value string) error {
	switch OutputFormat(value) {
	case OutputSimple, OutputJSON:
		*o = OutputFormat(value)
	default:
		return fmt.Errorf("invalid output format: %s", value)
	}
	return nil
}

func (o OutputFormat) String() string {
	return string(o)
}

func (o OutputFormat) Type() string {
	return "OutputFormat"
}

type showOptions struct {
	projectID string
	output    OutputFormat
}

func newShowCommand() *cobra.Command {
	opts := &showOptions{
		output: OutputSimple,
	}
	cmd := &cobra.Command{
		Use:     "show [task-id]",
		Aliases: []string{"info"},
		Short:   "Show details of a task",
		Args:    cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.projectID != "" {
				projectID = opts.projectID
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			taskId := args[0]

			task, err := TickliClient.GetTask(projectID, taskId)
			if err != nil {
				return err
			}
			if task.ID != taskId {
				log.Warn().Str("task-id", taskId).Str("project-id", projectID).Msg("task not found")
				return fmt.Errorf("task %s not found for porject %s", taskId, projectID)
			}
			switch opts.output {
			case OutputSimple:
				fmt.Println(utils.GetTaskDescription(task, types.DefaultColor))
			case OutputJSON:
				jsonData, err := json.MarshalIndent(task, "", "  ")
				if err != nil {
					return errors.Wrap(err, "failed to marshal output")
				}
				fmt.Println(string(jsonData))
			}
			fmt.Println(task.ID)
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.projectID, "project-id", "", "Project ID containing the task (default is current project)")
	cmd.Flags().VarP(&opts.output, "output", "o", "Output format. One of: simple, json")
	return cmd
}
