package utils

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/sho0pi/tickli/internal/types"
)

func GetProjectDescription(project types.Project) string {
	var projectStatus string
	if project.Closed {
		projectStatus = "Closed"
	} else {
		projectStatus = "Open"
	}

	projectLine := project.Color.Sprint("■■■■■■■■■■■■■■■■■■■■■■■■", project.Color)

	description := fmt.Sprintf(`
Project Details:

%s
Name: %s
ID: %s
Type: %s 
Status: %s
Group: %s

Tasks:`,
		projectLine,
		project.Name,
		project.ID,
		project.Kind,
		projectStatus,
		project.GroupID,
	)

	return description
}

func GetTaskDescription(task types.Task, projectColorHEX string) string {
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
Priority: %s
Group: %s

Time: 
StartDate: %s
DueDate: %s
CompletedTime: %s

Tasks:`,
		task.Status,
		projectLine,
		task.Title,
		task.Desc,
		task.Content,
		task.Priority.String(),
		task.ProjectID,
		task.StartDate.Humanize(),
		task.DueDate,
		task.CompletedTime.String(),
	)

	return description
}

func FuzzySelectProject(projects []types.Project, query string) (*types.Project, error) {
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
		return &types.Project{}, err
	}

	return &projects[idx], nil
}

func FuzzySelectTask(tasks []types.Task, projectColorHEX string, query string) (*types.Task, error) {
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
		return &types.Task{}, err
	}

	return &tasks[idx], nil
}
