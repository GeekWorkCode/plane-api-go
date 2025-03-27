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

	// Test Create method with Actor
	// 测试使用MemberID创建评论
	t.Run("CreateWithActor", func(t *testing.T) {
		// Get a member ID to use for the comment
		var memberID string

		// Try to get a member ID from environment variable first
		memberID = os.Getenv("PLANE_MEMBER_ID")

		// If no member ID provided, try to fetch one from the project
		if memberID == "" {
			membersService := NewMembersService(client)
			members, err := membersService.List(workspaceSlug, projectID)
			if err == nil && len(members) > 0 {
				memberID = members[0].Member.ID
			}
		}

		if memberID == "" {
			t.Skip("No member ID available for testing")
		}

		createReq := &CommentRequest{
			CommentHTML: "Test comment with actor",
			Actor:       memberID,
		}

		t.Logf("Using actor ID: %s for comment creation", memberID)

		comment, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.NotEmpty(t, comment.ID)
		commentID = comment.ID

		// Verify member information
		if comment.Member != nil {
			assert.Equal(t, memberID, comment.Member.ID)
		} else {
			t.Log("Comment was created with actor but API didn't return member information")
		}
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		if commentID == "" {
			t.Skip("No comment ID available for testing")
		}

		comment, err := s.Get(workspaceSlug, projectID, issueID, commentID)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.Equal(t, commentID, comment.ID)
	})

	// Test Update method with DisplayName
	// 测试使用DisplayName更新评论
	t.Run("UpdateWithDisplayName", func(t *testing.T) {
		if commentID == "" {
			t.Skip("No comment ID available for testing")
		}

		// Get a member's display name to use for the comment update
		var displayName string

		// Fetch a member to get their display name
		membersService := NewMembersService(client)
		members, err := membersService.List(workspaceSlug, projectID)
		if err == nil && len(members) > 0 {
			displayName = members[0].Member.DisplayName
			t.Logf("Using display name: %s for comment update", displayName)
		}

		if displayName == "" {
			t.Skip("No member display name available for testing")
		}

		updateReq := &CommentRequest{
			CommentHTML: "Updated comment with display name",
			DisplayName: displayName,
		}

		comment, err := s.Update(workspaceSlug, projectID, issueID, commentID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.Equal(t, commentID, comment.ID)
		assert.Equal(t, updateReq.CommentHTML, comment.CommentHTML)

		// Verify member information if we provided a display name
		if displayName != "" && comment.Member != nil {
			assert.Equal(t, displayName, comment.Member.DisplayName)
		} else {
			t.Log("Comment was updated with display name but API didn't return member information")
		}
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		if commentID == "" {
			t.Skip("No comment ID available for testing")
		}

		err := s.Delete(workspaceSlug, projectID, issueID, commentID)
		assert.NoError(t, err)
	})

	// Test creating a comment with DisplayName
	// 测试使用DisplayName创建评论
	t.Run("CreateWithDisplayName", func(t *testing.T) {
		// Get a member's display name
		membersService := NewMembersService(client)
		members, err := membersService.List(workspaceSlug, projectID)
		assert.NoError(t, err)

		if len(members) == 0 {
			t.Skip("No members available for testing")
		}

		displayName := members[0].Member.DisplayName
		t.Logf("Using display name: %s for comment creation", displayName)

		createReq := &CommentRequest{
			CommentHTML: "Test comment with display name",
			DisplayName: displayName,
		}

		comment, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.NotEmpty(t, comment.ID)

		// Verify member information
		if comment.Member != nil {
			assert.Equal(t, displayName, comment.Member.DisplayName)
		} else {
			t.Log("Comment was created with display name but API didn't return member information")
		}

		// Clean up
		err = s.Delete(workspaceSlug, projectID, issueID, comment.ID)
		assert.NoError(t, err)
	})
}
