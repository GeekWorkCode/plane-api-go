package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
	"github.com/stretchr/testify/assert"
)

// TestWorklogsService tests all methods of the WorklogsService
// 测试 WorklogsService 的所有方法
func TestWorklogsService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewWorklogsService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	issueID := os.Getenv("PLANE_ISSUE_ID")

	if workspaceSlug == "" || projectID == "" || issueID == "" {
		t.Skip("Required environment variables not set")
	}

	var worklogID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		worklogs, err := s.List(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
		assert.NotNil(t, worklogs)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &WorklogCreateRequest{
			Description: "Test worklog description",
			Duration:    3600, // 1 hour in seconds
		}
		worklog, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, worklog)
		assert.NotEmpty(t, worklog.ID)
		worklogID = worklog.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		worklog, err := s.Get(workspaceSlug, projectID, issueID, worklogID)
		assert.NoError(t, err)
		assert.NotNil(t, worklog)
		assert.Equal(t, worklogID, worklog.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &WorklogUpdateRequest{
			Description: "Updated test worklog description",
			Duration:    7200, // 2 hours in seconds
		}
		worklog, err := s.Update(workspaceSlug, projectID, issueID, worklogID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, worklog)
		assert.Equal(t, worklogID, worklog.ID)
		assert.Equal(t, updateReq.Description, worklog.Description)
		assert.Equal(t, updateReq.Duration, worklog.Duration)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, issueID, worklogID)
		assert.NoError(t, err)
	})

	// Test GetTotalTime method
	// 测试 GetTotalTime 方法
	t.Run("GetTotalTime", func(t *testing.T) {
		// First create a test worklog
		// 首先创建一个测试工作日志
		createReq := &WorklogCreateRequest{
			Description: "Test worklog for total time",
			Duration:    1800, // 30 minutes in seconds
		}
		worklog, err := s.Create(workspaceSlug, projectID, issueID, createReq)
		assert.NoError(t, err)
		worklogID = worklog.ID

		// Get total time
		// 获取总时间
		totals, err := s.GetTotalTime(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, totals)
		assert.NotEmpty(t, totals)

		// Find the total for our issue
		// 找到我们问题的时间总计
		var issueTotal *models.WorklogTotal
		for _, total := range totals {
			if total.IssueID == issueID {
				issueTotal = &total
				break
			}
		}
		assert.NotNil(t, issueTotal)
		assert.Equal(t, createReq.Duration, issueTotal.Duration)

		// Clean up
		// 清理
		err = s.Delete(workspaceSlug, projectID, issueID, worklogID)
		assert.NoError(t, err)
	})
}
