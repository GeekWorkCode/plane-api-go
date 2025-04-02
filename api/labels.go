package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// LabelsService handles communication with the label related endpoints
type LabelsService struct {
	client *client.Client
}

// NewLabelsService creates a new labels service
func NewLabelsService(client *client.Client) *LabelsService {
	return &LabelsService{
		client: client,
	}
}

// LabelCreateRequest represents the request body for creating a label
type LabelCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Color       string  `json:"color,omitempty"`
	Parent      *string `json:"parent,omitempty"`
}

// LabelUpdateRequest represents the request body for updating a label
type LabelUpdateRequest struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       string  `json:"color,omitempty"`
	Parent      *string `json:"parent,omitempty"`
}

// List returns all labels in a project
func (s *LabelsService) List(workspaceSlug string, projectID string) ([]models.Label, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/labels/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.LabelsResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// Get returns a label by its ID
func (s *LabelsService) Get(workspaceSlug string, projectID string, labelID string) (*models.Label, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/labels/%s", workspaceSlug, projectID, labelID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	label := new(models.Label)
	_, err = s.client.Do(req, label)
	return label, err
}

// Create creates a new label
func (s *LabelsService) Create(workspaceSlug string, projectID string, createRequest *LabelCreateRequest) (*models.Label, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/labels/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	label := new(models.Label)
	_, err = s.client.Do(req, label)
	return label, err
}

// Update updates a label
func (s *LabelsService) Update(workspaceSlug string, projectID string, labelID string, updateRequest *LabelUpdateRequest) (*models.Label, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/labels/%s", workspaceSlug, projectID, labelID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	label := new(models.Label)
	_, err = s.client.Do(req, label)
	return label, err
}

// Delete deletes a label
func (s *LabelsService) Delete(workspaceSlug string, projectID string, labelID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/labels/%s", workspaceSlug, projectID, labelID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
