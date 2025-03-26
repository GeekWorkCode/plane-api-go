package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// WorklogsService handles communication with the worklog related endpoints
type WorklogsService struct {
	client *client.Client
}

// NewWorklogsService creates a new worklogs service
func NewWorklogsService(client *client.Client) *WorklogsService {
	return &WorklogsService{
		client: client,
	}
}

// WorklogCreateRequest represents the request body for creating a worklog
type WorklogCreateRequest struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"` // Duration in minutes
}

// WorklogUpdateRequest represents the request body for updating a worklog
type WorklogUpdateRequest struct {
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"` // Duration in minutes
}

// List returns all worklogs for an issue
func (s *WorklogsService) List(workspaceSlug string, projectID string, issueID string) ([]models.Worklog, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/worklogs/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var worklogs []models.Worklog
	_, err = s.client.Do(req, &worklogs)
	return worklogs, err
}

// Create creates a new worklog for an issue
func (s *WorklogsService) Create(workspaceSlug string, projectID string, issueID string, createRequest *WorklogCreateRequest) (*models.Worklog, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/worklogs/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	worklog := new(models.Worklog)
	_, err = s.client.Do(req, worklog)
	return worklog, err
}

// Update updates a worklog
func (s *WorklogsService) Update(workspaceSlug string, projectID string, issueID string, worklogID string, updateRequest *WorklogUpdateRequest) (*models.Worklog, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/worklogs/%s/", workspaceSlug, projectID, issueID, worklogID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	worklog := new(models.Worklog)
	_, err = s.client.Do(req, worklog)
	return worklog, err
}

// Delete deletes a worklog
func (s *WorklogsService) Delete(workspaceSlug string, projectID string, issueID string, worklogID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/worklogs/%s/", workspaceSlug, projectID, issueID, worklogID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// GetTotalTime returns the total time spent for all issues in a project
func (s *WorklogsService) GetTotalTime(workspaceSlug string, projectID string) ([]models.WorklogTotal, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/total-worklogs/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var totals []models.WorklogTotal
	_, err = s.client.Do(req, &totals)
	return totals, err
}

// Get returns a single worklog by ID
func (s *WorklogsService) Get(workspaceSlug string, projectID string, issueID string, worklogID string) (*models.Worklog, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/worklogs/%s/", workspaceSlug, projectID, issueID, worklogID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	worklog := new(models.Worklog)
	_, err = s.client.Do(req, worklog)
	return worklog, err
}
