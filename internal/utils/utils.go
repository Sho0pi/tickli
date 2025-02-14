package utils

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/sho0pi/tickli/internal/api"
)

func GetProjectDescription(project api.Project) string {
	var kindEmoji string
	switch project.Kind {
	case "TASK":
		kindEmoji = "📝"
	case "NOTE":
		kindEmoji = "🗒️"
	default:
		kindEmoji = "🔧"
	}

	var projectStatus string
	if project.Closed {
		projectStatus = "Closed"
	} else {
		projectStatus = "Open"
	}

	projectColor := color.HEX(project.Color, true)
	if project.Color == "" {
		projectColor = color.HEX("#000000", true)
	}
	projectLine := projectColor.Sprint("----------------------")

	description := fmt.Sprintf(`
Project Details:

%s
Name: %s - %s
ID: %s
Status: %s
Group: %s

Tasks:`,
		projectLine,
		kindEmoji,
		project.Name,
		project.ID,
		projectStatus,
		project.GroupID,
	)

	return description
}

func GetTaskDescription(task api.Task, projectColorHEX string) string {
	var statusEmoji string
	switch task.Status {
	case 0:
		statusEmoji = color.BgHiWhite.Sprint("[ ]")
	case 2:
		statusEmoji = color.BgHiGreen.Sprint("[✔]")
	default:
		statusEmoji = "[>>]"
	}

	projectColor := color.HEX(projectColorHEX, true)
	if projectColorHEX == "" {
		projectColor = color.HEX("#000000", true)
	}
	projectLine := projectColor.Sprint("----------------------")

	description := fmt.Sprintf(`
Task Details:

%s %s
%s
Desc: %s 
Content: %s
Group: %s

Tasks:`,
		statusEmoji,
		projectLine,
		task.Title,
		task.Desc,
		task.Content,
		task.ProjectID,
	)

	return description
}

func FuzzySelectProject(projects []api.Project, query string) (api.Project, error) {
	idx, err := fuzzyfinder.Find(
		projects,
		func(i int) string {
			return fmt.Sprintf("%s (%s)",
				projects[i].Name,
				projects[i].ID,
			)
		},
		fuzzyfinder.WithQuery(query),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return GetProjectDescription(projects[i])
		}),
		fuzzyfinder.WithPromptString("Search Project: "),
	)
	if err != nil {
		return api.Project{}, err
	}

	return projects[idx], nil
}

func FuzzySelectTask(tasks []api.Task, projectColorHEX string, query string) (api.Task, error) {
	idx, err := fuzzyfinder.Find(
		tasks,
		func(i int) string {
			return fmt.Sprintf("%s (%s)",
				tasks[i].Title,
				tasks[i].ID,
			)
		},
		fuzzyfinder.WithQuery(query),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return GetTaskDescription(tasks[i], projectColorHEX)
		}),
		fuzzyfinder.WithPromptString("Search Project: "),
	)
	if err != nil {
		return api.Task{}, err
	}

	return tasks[idx], nil
}
