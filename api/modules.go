package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// ModulesService handles communication with the module related endpoints
type ModulesService struct {
	client *client.Client
}

// NewModulesService creates a new modules service
func NewModulesService(client *client.Client) *ModulesService {
	return &ModulesService{
		client: client,
	}
}

// ModuleCreateRequest represents the request body for creating a module
type ModuleCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// ModuleUpdateRequest represents the request body for updating a module
type ModuleUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all modules in a project
func (s *ModulesService) List(workspaceSlug string, projectID string) ([]models.Module, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var modules []models.Module
	_, err = s.client.Do(req, &modules)
	return modules, err
}

// Get returns a module by its ID
func (s *ModulesService) Get(workspaceSlug string, projectID string, moduleID string) (*models.Module, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s", workspaceSlug, projectID, moduleID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	module := new(models.Module)
	_, err = s.client.Do(req, module)
	return module, err
}

// Create creates a new module
func (s *ModulesService) Create(workspaceSlug string, projectID string, createRequest *ModuleCreateRequest) (*models.Module, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	module := new(models.Module)
	_, err = s.client.Do(req, module)
	return module, err
}

// Update updates a module
func (s *ModulesService) Update(workspaceSlug string, projectID string, moduleID string, updateRequest *ModuleUpdateRequest) (*models.Module, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s", workspaceSlug, projectID, moduleID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	module := new(models.Module)
	_, err = s.client.Do(req, module)
	return module, err
}

// Delete deletes a module
func (s *ModulesService) Delete(workspaceSlug string, projectID string, moduleID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s", workspaceSlug, projectID, moduleID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// ModuleIssueAddRequest represents the request body for adding issues to a module
type ModuleIssueAddRequest struct {
	Issues []string `json:"issues"`
}

// ListIssues returns all issues in a module
func (s *ModulesService) ListIssues(workspaceSlug string, projectID string, moduleID string) ([]models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s/module-issues/", workspaceSlug, projectID, moduleID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var issues []models.Issue
	_, err = s.client.Do(req, &issues)
	return issues, err
}

// AddIssues adds issues to a module
func (s *ModulesService) AddIssues(workspaceSlug string, projectID string, moduleID string, issueIDs []string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s/module-issues/", workspaceSlug, projectID, moduleID)

	req, err := s.client.NewRequest(http.MethodPost, path, &ModuleIssueAddRequest{
		Issues: issueIDs,
	})
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// RemoveIssue removes an issue from a module
func (s *ModulesService) RemoveIssue(workspaceSlug string, projectID string, moduleID string, issueID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/modules/%s/module-issues/%s", workspaceSlug, projectID, moduleID, issueID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
