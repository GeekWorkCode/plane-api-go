package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestCyclesService tests all methods of the CyclesService
// 测试 CyclesService 的所有方法
func TestCyclesService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewCyclesService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	var cycleID string
	var issueID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		cycles, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, cycles)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &CycleCreateRequest{
			Name:        "Test Cycle",
			Description: "Test cycle description",
			StartDate:   "2024-01-01",
			EndDate:     "2024-01-31",
		}
		cycle, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, cycle)
		assert.NotEmpty(t, cycle.ID)
		cycleID = cycle.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		cycle, err := s.Get(workspaceSlug, projectID, cycleID)
		assert.NoError(t, err)
		assert.NotNil(t, cycle)
		assert.Equal(t, cycleID, cycle.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &CycleUpdateRequest{
			Name:        "Updated Test Cycle",
			Description: "Updated test cycle description",
		}
		cycle, err := s.Update(workspaceSlug, projectID, cycleID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, cycle)
		assert.Equal(t, cycleID, cycle.ID)
		assert.Equal(t, updateReq.Name, cycle.Name)
		assert.Equal(t, updateReq.Description, cycle.Description)
	})

	// Test ListIssues method
	// 测试 ListIssues 方法
	t.Run("ListIssues", func(t *testing.T) {
		issues, err := s.ListIssues(workspaceSlug, projectID, cycleID)
		assert.NoError(t, err)
		assert.NotNil(t, issues)
	})

	// Test AddIssues method
	// 测试 AddIssues 方法
	t.Run("AddIssues", func(t *testing.T) {
		// First create a test issue
		// 首先创建一个测试问题
		issuesService := NewIssuesService(c)
		issueReq := &IssueCreateRequest{
			Name: "Test Issue for Cycle",
		}
		issue, err := issuesService.Create(workspaceSlug, projectID, issueReq)
		assert.NoError(t, err)
		issueID = issue.ID

		// Add the issue to the cycle
		// 将问题添加到周期中
		err = s.AddIssues(workspaceSlug, projectID, cycleID, []string{issueID})
		assert.NoError(t, err)
	})

	// Test RemoveIssue method
	// 测试 RemoveIssue 方法
	t.Run("RemoveIssue", func(t *testing.T) {
		err := s.RemoveIssue(workspaceSlug, projectID, cycleID, issueID)
		assert.NoError(t, err)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, cycleID)
		assert.NoError(t, err)
	})
}
