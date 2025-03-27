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

// CommentRequest represents the request body for creating or updating a comment
type CommentRequest struct {
	CommentHTML string `json:"comment_html"`
	CreatedBy   string `json:"created_by,omitempty"` // ID of the member who created the comment
	Actor       string `json:"actor,omitempty"`      // ID of the member who updated the comment (used for update only)
	DisplayName string `json:"-"`                    // 成员显示名称，不会直接发送到API
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

// 根据显示名称查找成员ID
func (s *CommentsService) findMemberIDByDisplayName(workspaceSlug string, projectID string, displayName string) (string, error) {
	// 使用Members服务获取所有成员
	membersService := NewMembersService(s.client)
	members, err := membersService.List(workspaceSlug, projectID)
	if err != nil {
		return "", fmt.Errorf("获取成员列表失败: %w", err)
	}

	// 查找匹配DisplayName的成员
	for _, member := range members {
		if member.Member.DisplayName == displayName {
			return member.Member.ID, nil
		}
	}

	return "", fmt.Errorf("找不到显示名称为 '%s' 的成员", displayName)
}

// prepareCommentRequest 处理评论请求，如果提供了DisplayName则转换为MemberID
func (s *CommentsService) prepareCommentRequest(workspaceSlug string, projectID string, request *CommentRequest) error {
	// 如果提供了DisplayName但没有CreatedBy/Actor，尝试查找对应的MemberID
	if request.DisplayName != "" {
		if request.CreatedBy == "" && request.Actor == "" {
			memberID, err := s.findMemberIDByDisplayName(workspaceSlug, projectID, request.DisplayName)
			if err != nil {
				return err
			}
			// 根据是创建还是更新操作，设置不同的字段
			request.CreatedBy = memberID
			request.Actor = memberID
		}
	}

	// 清除DisplayName字段，确保它不会被发送到API
	request.DisplayName = ""
	return nil
}

// Create creates a new comment
// 支持通过DisplayName或CreatedBy(MemberID)创建评论
func (s *CommentsService) Create(workspaceSlug string, projectID string, issueID string, request *CommentRequest) (*models.Comment, error) {
	// 处理DisplayName，如果有的话
	err := s.prepareCommentRequest(workspaceSlug, projectID, request)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	comment := new(models.Comment)
	_, err = s.client.Do(req, comment)
	if err != nil {
		return nil, fmt.Errorf("创建评论失败: %w", err)
	}

	// 保存创建者ID，用于后续可能的校正
	originalCreatorID := request.CreatedBy

	// 如果创建者ID被提供了，但响应中显示的创建者不匹配，尝试通过更新操作修正这个问题
	if originalCreatorID != "" && (comment.Member == nil || comment.CreatedBy != originalCreatorID) {
		// 使用刚创建的评论ID，同样的内容再执行一次更新操作，
		// 但这次使用Actor字段传递创建者ID
		updateReq := &CommentRequest{
			CommentHTML: request.CommentHTML,
			Actor:       originalCreatorID,
		}

		// 执行更新操作
		updatedComment, updateErr := s.Update(workspaceSlug, projectID, issueID, comment.ID, updateReq)
		if updateErr == nil {
			// 如果更新成功，返回更新后的评论
			return updatedComment, nil
		}
		// 如果更新失败，记录日志但继续返回原始创建的评论
		fmt.Printf("警告: 尝试修正评论作者显示失败: %v\n", updateErr)
	}

	// If a memberID was provided and the comment was created successfully,
	// but the API didn't populate the Member field, we'll set it manually
	if request.CreatedBy != "" && comment.Member == nil {
		// Fetch member details if needed
		membersService := NewMembersService(s.client)
		member, err := membersService.Get(workspaceSlug, projectID, request.CreatedBy)
		if err == nil && member != nil {
			comment.Member = &member.Member
		}
	}

	return comment, nil
}

// Update updates a comment
// 支持通过DisplayName或Actor(MemberID)更新评论
func (s *CommentsService) Update(workspaceSlug string, projectID string, issueID string, commentID string, request *CommentRequest) (*models.Comment, error) {
	// 处理DisplayName，如果有的话
	err := s.prepareCommentRequest(workspaceSlug, projectID, request)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/comments/%s/", workspaceSlug, projectID, issueID, commentID)
	req, err := s.client.NewRequest(http.MethodPatch, path, request)
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
