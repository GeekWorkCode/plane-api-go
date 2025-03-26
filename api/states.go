package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// StatesService handles communication with the state related endpoints
type StatesService struct {
	client *client.Client
}

// NewStatesService creates a new states service
func NewStatesService(client *client.Client) *StatesService {
	return &StatesService{
		client: client,
	}
}

// StateCreateRequest represents the request body for creating a state
type StateCreateRequest struct {
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description,omitempty"`
}

// StateUpdateRequest represents the request body for updating a state
type StateUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

// List returns all states in a project
func (s *StatesService) List(workspaceSlug string, projectID string) ([]models.State, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/states/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var states []models.State
	_, err = s.client.Do(req, &states)
	return states, err
}

// Get returns a state by its ID
func (s *StatesService) Get(workspaceSlug string, projectID string, stateID string) (*models.State, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/states/%s/", workspaceSlug, projectID, stateID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	state := new(models.State)
	_, err = s.client.Do(req, state)
	return state, err
}

// Create creates a new state
func (s *StatesService) Create(workspaceSlug string, projectID string, createRequest *StateCreateRequest) (*models.State, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/states/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	state := new(models.State)
	_, err = s.client.Do(req, state)
	return state, err
}

// Update updates a state
func (s *StatesService) Update(workspaceSlug string, projectID string, stateID string, updateRequest *StateUpdateRequest) (*models.State, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/states/%s/", workspaceSlug, projectID, stateID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	state := new(models.State)
	_, err = s.client.Do(req, state)
	return state, err
}

// Delete deletes a state
func (s *StatesService) Delete(workspaceSlug string, projectID string, stateID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/states/%s/", workspaceSlug, projectID, stateID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
