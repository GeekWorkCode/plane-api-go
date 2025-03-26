package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestLinksService tests all methods of the LinksService
// 测试 LinksService 的所有方法
func TestLinksService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewLinksService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	issueID := os.Getenv("PLANE_ISSUE_ID")

	if workspaceSlug == "" || projectID == "" || issueID == "" {
		t.Skip("Required environment variables not set")
	}

	var linkID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		links, err := s.List(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
		assert.NotNil(t, links)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &LinkCreateRequest{
			Title: "Test Link",
			URL:   "https://example.com",
		}
		link, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, link)
		assert.NotEmpty(t, link.ID)
		linkID = link.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		link, err := s.Get(workspaceSlug, projectID, issueID, linkID)
		assert.NoError(t, err)
		assert.NotNil(t, link)
		assert.Equal(t, linkID, link.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &LinkUpdateRequest{
			Title: "Updated Test Link",
			URL:   "https://updated-example.com",
		}
		link, err := s.Update(workspaceSlug, projectID, issueID, linkID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, link)
		assert.Equal(t, linkID, link.ID)
		assert.Equal(t, updateReq.Title, link.Title)
		assert.Equal(t, updateReq.URL, link.URL)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, issueID, linkID)
		assert.NoError(t, err)
	})
}
