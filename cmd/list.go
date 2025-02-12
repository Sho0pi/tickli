package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/sho0pi/tickli/internal/api"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects or tasks",
	Long:  `List projects or tasks from TickTick.`,
}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"lsp"},
	Short:   "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load config")
		}

		client := api.NewClient(cfg.AccessToken)
		projects, err := client.ListProjects()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to list projects")
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tCOLOR\tSTATUS")
		for _, p := range projects {
			status := "Active"
			if p.Closed {
				status = "Closed"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", p.ID, p.Name, p.Color, status)
		}
		w.Flush()
	},
}

var listTasksCmd = &cobra.Command{
	Use:     "tasks [PROJECT_ID]",
	Aliases: []string{"ls"},
	Short:   "List all tasks in a project",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load config")
		}

		projectID := cfg.DefaultProjectID
		if len(args) > 0 {
			projectID = args[0]
		}

		if projectID == "" {
			log.Fatal().Msg("No project ID specified. Run 'tickli set-project' or provide PROJECT_ID as argument")
		}

		client := api.NewClient(cfg.AccessToken)
		tasks, err := client.ListTasks(projectID)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to list tasks")
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTITLE\tDUE DATE\tPRIORITY\tSTATUS")
		for _, t := range tasks {
			status := "Open"
			if t.Status == 2 {
				status = "Completed"
			}
			priority := getPriorityString(t.Priority)
			dueDate := "-"
			//TODO fix time problem
			//if t.DueDate != nil {
			//	dueDate = t.DueDate.Format("2006-01-02")
			//}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", t.ID, t.Title, dueDate, priority, status)
		}
		w.Flush()
	},
}

func getPriorityString(priority int) string {
	switch priority {
	case 0:
		return "None"
	case 1:
		return "Low"
	case 3:
		return "Medium"
	case 5:
		return "High"
	default:
		return "Unknown"
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listTasksCmd)
}
