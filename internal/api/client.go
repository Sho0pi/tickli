package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	baseURL     = "https://api.ticktick.com/open/v1"
	authURL     = "https://ticktick.com/oauth/authorize"
	tokenURL    = "https://ticktick.com/oauth/token"
	scope       = "tasks:write tasks:read"
	redirectURL = "http://localhost:8080"
)

type Client struct {
	http *resty.Client
}

func NewClient(token string) *Client {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Authorization", "Bearer "+token)

	return &Client{http: client}
}

func GetAuthURL(clientID string) string {
	return fmt.Sprintf("%s?scope=%s&client_id=%s&state=state&redirect_uri=%s&response_type=code",
		authURL, scope, clientID, redirectURL)
}

func GetAccessToken(clientID, clientSecret, code string) (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetBasicAuth(clientID, clientSecret).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":   "authorization_code",
			"code":         code,
			"redirect_uri": redirectURL,
		}).
		Post(tokenURL)

	if err != nil {
		return "", errors.Wrap(err, "requesting access token")
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", errors.Wrap(err, "parsing response")
	}

	return result.AccessToken, nil
}

func (c *Client) ListProjects() ([]Project, error) {
	var projects []Project
	resp, err := c.http.R().
		SetResult(&projects).
		Get("/project")

	if err != nil {
		return nil, errors.Wrap(err, "listing projects")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to list projects: %s", resp.String())
	}

	return projects, nil
}

func (c *Client) ListTasks(projectID string) ([]Task, error) {
	var projectData struct {
		Tasks []Task `json:"tasks"`
	}
	resp, err := c.http.R().
		SetResult(&projectData).
		Get(fmt.Sprintf("/project/%s/data", projectID))

	if err != nil {
		return nil, errors.Wrap(err, "listing tasks")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to list tasks: %s", resp.String())
	}

	return projectData.Tasks, nil
}

func (c *Client) CreateTask(projectID, title string) (*Task, error) {
	task := &Task{
		ProjectID: projectID,
		Title:     title,
	}

	resp, err := c.http.R().
		SetBody(task).
		SetResult(task).
		Post("/task")

	if err != nil {
		return nil, errors.Wrap(err, "creating task")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to create task: %s", resp.String())
	}

	return task, nil
}
