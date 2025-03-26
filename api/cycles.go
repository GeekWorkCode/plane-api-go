package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// CyclesService handles communication with the cycle related endpoints
type CyclesService struct {
	client *client.Client
}

// NewCyclesService creates a new cycles service
func NewCyclesService(client *client.Client) *CyclesService {
	return &CyclesService{
		client: client,
	}
}

// CycleCreateRequest represents the request body for creating a cycle
type CycleCreateRequest struct {
	Name string `json:"name"`
}

// CycleUpdateRequest represents the request body for updating a cycle
type CycleUpdateRequest struct {
	Name string `json:"name,omitempty"`
}

// List returns all cycles in a project
func (s *CyclesService) List(workspaceSlug string, projectID string) ([]models.Cycle, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.CyclesResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// Get returns a cycle by its ID
func (s *CyclesService) Get(workspaceSlug string, projectID string, cycleID string) (*models.Cycle, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/", workspaceSlug, projectID, cycleID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	cycle := new(models.Cycle)
	_, err = s.client.Do(req, cycle)
	return cycle, err
}

// Create creates a new cycle
func (s *CyclesService) Create(workspaceSlug string, projectID string, createRequest *CycleCreateRequest) (*models.Cycle, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	cycle := new(models.Cycle)
	resp, err := s.client.Do(req, cycle)
	if err != nil {
		// 如果服务器返回错误消息，尝试解析并返回
		if resp != nil && resp.StatusCode >= 400 {
			return nil, fmt.Errorf("创建周期失败(HTTP %d): %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("创建周期失败: %w", err)
	}
	return cycle, nil
}

// Update updates a cycle
func (s *CyclesService) Update(workspaceSlug string, projectID string, cycleID string, updateRequest *CycleUpdateRequest) (*models.Cycle, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/", workspaceSlug, projectID, cycleID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	cycle := new(models.Cycle)
	_, err = s.client.Do(req, cycle)
	return cycle, err
}

// Delete deletes a cycle
func (s *CyclesService) Delete(workspaceSlug string, projectID string, cycleID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/", workspaceSlug, projectID, cycleID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// CycleIssueAddRequest represents the request body for adding issues to a cycle
type CycleIssueAddRequest struct {
	Issues []string `json:"issues"`
}

// ListIssues returns all issues in a cycle
func (s *CyclesService) ListIssues(workspaceSlug string, projectID string, cycleID string) ([]models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/cycle-issues/", workspaceSlug, projectID, cycleID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var issues []models.Issue
	_, err = s.client.Do(req, &issues)
	return issues, err
}

// AddIssues adds issues to a cycle
func (s *CyclesService) AddIssues(workspaceSlug string, projectID string, cycleID string, issueIDs []string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/cycle-issues/", workspaceSlug, projectID, cycleID)

	req, err := s.client.NewRequest(http.MethodPost, path, &CycleIssueAddRequest{
		Issues: issueIDs,
	})
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// RemoveIssue removes an issue from a cycle
func (s *CyclesService) RemoveIssue(workspaceSlug string, projectID string, cycleID string, issueID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/cycles/%s/cycle-issues/%s/", workspaceSlug, projectID, cycleID, issueID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
