package api

import (
	"fmt"
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
	var sequenceID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		issues, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, issues)

		// 如果有问题存在，获取第一个问题的序列ID用于后续测试
		if len(issues) > 0 {
			// 在实际情况中，sequenceID通常以"PROJECT_IDENTIFIER-NUMBER"格式存在
			// 这里假设我们已知道项目标识符，或者可以从项目ID获取
			// 以下只是示例，实际情况需要根据API响应结构调整
			projectsService := NewProjectsService(client)
			project, err := projectsService.Get(workspaceSlug, projectID)
			if err == nil && project != nil {
				// 假设问题序列号为1，实际中应该从问题数据中获取
				sequenceID = fmt.Sprintf("%s-%d", project.Identifier, 1)
			}
		}
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

	// Test GetBySequenceID method (if sequenceID is available)
	// 测试 GetBySequenceID 方法（如果序列ID可用）
	t.Run("GetBySequenceID", func(t *testing.T) {
		if sequenceID == "" {
			t.Skip("No sequence ID available for testing")
		}

		issue, err := s.GetBySequenceID(workspaceSlug, sequenceID)
		assert.NoError(t, err)
		assert.NotNil(t, issue)
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

	// Test UpdateBySequenceID method (if sequenceID is available)
	// 测试 UpdateBySequenceID 方法（如果序列ID可用）
	t.Run("UpdateBySequenceID", func(t *testing.T) {
		if sequenceID == "" {
			t.Skip("No sequence ID available for testing")
		}

		updateReq := &IssueUpdateRequest{
			Name:        "Updated By Sequence ID",
			Description: "Issue updated by sequence ID",
		}
		issue, err := s.UpdateBySequenceID(workspaceSlug, sequenceID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, issue)
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
