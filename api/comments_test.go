package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestCommentsService tests all methods of the CommentsService
// 测试 CommentsService 的所有方法
func TestCommentsService(t *testing.T) {
	// Get API key from environment
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}

	// Get required IDs from environment
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	issueID := os.Getenv("PLANE_ISSUE_ID")

	if workspaceSlug == "" || projectID == "" || issueID == "" {
		t.Skip("Required environment variables not set")
	}

	// Create a new client
	client := client.NewClient(apiKey)
	s := NewCommentsService(client)

	var commentID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		comments, err := s.List(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
		assert.NotNil(t, comments)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &CommentCreateRequest{
			CommentHTML: "Test comment content",
		}
		comment, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.NotEmpty(t, comment.ID)
		commentID = comment.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		comment, err := s.Get(workspaceSlug, projectID, issueID, commentID)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.Equal(t, commentID, comment.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &CommentUpdateRequest{
			CommentHTML: "Updated test comment content",
		}
		comment, err := s.Update(workspaceSlug, projectID, issueID, commentID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.Equal(t, commentID, comment.ID)
		assert.Equal(t, updateReq.CommentHTML, comment.CommentHTML)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, issueID, commentID)
		assert.NoError(t, err)
	})
}
