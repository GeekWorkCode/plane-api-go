package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// LinksService handles communication with the issue links related endpoints
type LinksService struct {
	client *client.Client
}

// NewLinksService creates a new links service
func NewLinksService(client *client.Client) *LinksService {
	return &LinksService{
		client: client,
	}
}

// LinkCreateRequest represents the request body for creating a link
type LinkCreateRequest struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url"`
}

// LinkUpdateRequest represents the request body for updating a link
type LinkUpdateRequest struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// List returns all links for an issue
func (s *LinksService) List(workspaceSlug string, projectID string, issueID string) ([]models.Link, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/links/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var links []models.Link
	_, err = s.client.Do(req, &links)
	return links, err
}

// Get returns a link by its ID
func (s *LinksService) Get(workspaceSlug string, projectID string, issueID string, linkID string) (*models.Link, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/links/%s", workspaceSlug, projectID, issueID, linkID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	link := new(models.Link)
	_, err = s.client.Do(req, link)
	return link, err
}

// Create creates a new link
func (s *LinksService) Create(workspaceSlug string, projectID string, issueID string, createRequest *LinkCreateRequest) (*models.Link, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/links/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	link := new(models.Link)
	_, err = s.client.Do(req, link)
	return link, err
}

// Update updates a link
func (s *LinksService) Update(workspaceSlug string, projectID string, issueID string, linkID string, updateRequest *LinkUpdateRequest) (*models.Link, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/links/%s", workspaceSlug, projectID, issueID, linkID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	link := new(models.Link)
	_, err = s.client.Do(req, link)
	return link, err
}

// Delete deletes a link
func (s *LinksService) Delete(workspaceSlug string, projectID string, issueID string, linkID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/links/%s", workspaceSlug, projectID, issueID, linkID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
