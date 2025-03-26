package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// ProjectsService handles communication with the project related endpoints
type ProjectsService struct {
	client *client.Client
}

// NewProjectsService creates a new projects service
func NewProjectsService(client *client.Client) *ProjectsService {
	return &ProjectsService{
		client: client,
	}
}

// ProjectCreateRequest represents the request body for creating a project
type ProjectCreateRequest struct {
	Name        string `json:"name"`
	Identifier  string `json:"identifier"`
	Description string `json:"description,omitempty"`
}

// ProjectUpdateRequest represents the request body for updating a project
type ProjectUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all projects in a workspace
func (s *ProjectsService) List(workspaceSlug string) ([]models.Project, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/", workspaceSlug)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.ProjectsResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// Get returns a project by its ID
func (s *ProjectsService) Get(workspaceSlug string, projectID string) (*models.Project, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	project := new(models.Project)
	_, err = s.client.Do(req, project)
	return project, err
}

// Create creates a new project
func (s *ProjectsService) Create(workspaceSlug string, createRequest *ProjectCreateRequest) (*models.Project, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/", workspaceSlug)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	project := new(models.Project)
	_, err = s.client.Do(req, project)
	return project, err
}

// Update updates a project
func (s *ProjectsService) Update(workspaceSlug string, projectID string, updateRequest *ProjectUpdateRequest) (*models.Project, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	project := new(models.Project)
	_, err = s.client.Do(req, project)
	return project, err
}

// Delete deletes a project
func (s *ProjectsService) Delete(workspaceSlug string, projectID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
