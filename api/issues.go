package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// IssuesService handles communication with the issue related endpoints
type IssuesService struct {
	client *client.Client
}

// NewIssuesService creates a new issues service
func NewIssuesService(client *client.Client) *IssuesService {
	return &IssuesService{
		client: client,
	}
}

// IssueCreateRequest represents the request body for creating an issue
type IssueCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	Priority    string `json:"priority,omitempty"`
	AssigneeID  string `json:"assignee_id,omitempty"`
}

// IssueUpdateRequest represents the request body for updating an issue
type IssueUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	Priority    string `json:"priority,omitempty"`
	AssigneeID  string `json:"assignee_id,omitempty"`
}

// List returns all issues in a project
func (s *IssuesService) List(workspaceSlug string, projectID string) ([]models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.IssuesResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// Get returns an issue by its ID
func (s *IssuesService) Get(workspaceSlug string, projectID string, issueID string) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// GetBySequenceID returns an issue by its sequence ID
func (s *IssuesService) GetBySequenceID(workspaceSlug string, sequenceID string) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/issues/%s/", workspaceSlug, sequenceID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// Create creates a new issue
func (s *IssuesService) Create(workspaceSlug string, projectID string, createRequest *IssueCreateRequest) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// Update updates an issue
func (s *IssuesService) Update(workspaceSlug string, projectID string, issueID string, updateRequest *IssueUpdateRequest) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// Delete deletes an issue
func (s *IssuesService) Delete(workspaceSlug string, projectID string, issueID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
