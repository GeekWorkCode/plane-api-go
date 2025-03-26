package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// CommentsService handles communication with the issue comments related endpoints
type CommentsService struct {
	client *client.Client
}

// NewCommentsService creates a new comments service
func NewCommentsService(client *client.Client) *CommentsService {
	return &CommentsService{
		client: client,
	}
}

// CommentCreateRequest represents the request body for creating a comment
type CommentCreateRequest struct {
	CommentHTML string `json:"comment_html"`
}

// CommentUpdateRequest represents the request body for updating a comment
type CommentUpdateRequest struct {
	CommentHTML string `json:"comment_html"`
}

// List returns all comments for an issue
func (s *CommentsService) List(workspaceSlug string, projectID string, issueID string) ([]models.Comment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	response := new(models.CommentsResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, fmt.Errorf("获取评论列表失败: %w", err)
	}
	return response.Results, nil
}

// Get returns a comment by its ID
func (s *CommentsService) Get(workspaceSlug string, projectID string, issueID string, commentID string) (*models.Comment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/%s/", workspaceSlug, projectID, issueID, commentID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	comment := new(models.Comment)
	_, err = s.client.Do(req, comment)
	if err != nil {
		return nil, fmt.Errorf("获取评论详情失败: %w", err)
	}
	return comment, nil
}

// Create creates a new comment
func (s *CommentsService) Create(workspaceSlug string, projectID string, issueID string, createRequest *CommentCreateRequest) (*models.Comment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	comment := new(models.Comment)
	_, err = s.client.Do(req, comment)
	if err != nil {
		return nil, fmt.Errorf("创建评论失败: %w", err)
	}
	return comment, nil
}

// Update updates a comment
func (s *CommentsService) Update(workspaceSlug string, projectID string, issueID string, commentID string, updateRequest *CommentUpdateRequest) (*models.Comment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/%s/", workspaceSlug, projectID, issueID, commentID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	comment := new(models.Comment)
	_, err = s.client.Do(req, comment)
	if err != nil {
		return nil, fmt.Errorf("更新评论失败: %w", err)
	}
	return comment, nil
}

// Delete deletes a comment
func (s *CommentsService) Delete(workspaceSlug string, projectID string, issueID string, commentID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/%s/", workspaceSlug, projectID, issueID, commentID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return fmt.Errorf("删除评论失败: %w", err)
	}
	return nil
}
