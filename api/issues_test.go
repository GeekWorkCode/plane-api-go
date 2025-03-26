package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestIssuesService tests all methods of the IssuesService
// 测试 IssuesService 的所有方法
func TestIssuesService(t *testing.T) {
	// Get API key from environment
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}

	// Get required IDs from environment
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	// Create a new client
	client := client.NewClient(apiKey)
	s := NewIssuesService(client)

	var issueID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		issues, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, issues)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &IssueCreateRequest{
			Name:        "Test Issue",
			Description: "Test Issue Description",
		}
		issue, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, issue)
		assert.NotEmpty(t, issue.ID)
		issueID = issue.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		issue, err := s.Get(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
		assert.NotNil(t, issue)
		assert.Equal(t, issueID, issue.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &IssueUpdateRequest{
			Name:        "Updated Test Issue",
			Description: "Updated Test Issue Description",
		}
		issue, err := s.Update(workspaceSlug, projectID, issueID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, issue)
		assert.Equal(t, issueID, issue.ID)
		assert.Equal(t, updateReq.Name, issue.Name)
		assert.Equal(t, updateReq.Description, issue.Description)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
	})
}
