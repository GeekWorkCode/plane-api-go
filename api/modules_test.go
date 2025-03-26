package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestModulesService tests all methods of the ModulesService
// 测试 ModulesService 的所有方法
func TestModulesService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewModulesService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	var moduleID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		modules, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, modules)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &ModuleCreateRequest{
			Name:        "Test Module",
			Description: "Test module description",
		}
		module, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, module)
		assert.NotEmpty(t, module.ID)
		moduleID = module.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		module, err := s.Get(workspaceSlug, projectID, moduleID)
		assert.NoError(t, err)
		assert.NotNil(t, module)
		assert.Equal(t, moduleID, module.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &ModuleUpdateRequest{
			Name:        "Updated Test Module",
			Description: "Updated test module description",
		}
		module, err := s.Update(workspaceSlug, projectID, moduleID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, module)
		assert.Equal(t, moduleID, module.ID)
		assert.Equal(t, updateReq.Name, module.Name)
		assert.Equal(t, updateReq.Description, module.Description)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, moduleID)
		assert.NoError(t, err)
	})

	// Test AddIssue method
	// 测试 AddIssue 方法
	t.Run("AddIssue", func(t *testing.T) {
		// First create a test module
		// 首先创建一个测试模块
		createReq := &ModuleCreateRequest{
			Name:        "Test Module for Issue",
			Description: "Test module for issue",
		}
		module, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		moduleID = module.ID

		// Create a test issue
		// 创建一个测试问题
		issuesService := NewIssuesService(c)
		issueReq := &IssueCreateRequest{
			Name: "Test Issue for Module",
		}
		issue, err := issuesService.Create(workspaceSlug, projectID, issueReq)
		assert.NoError(t, err)

		// Add issue to module
		// 将问题添加到模块中
		err = s.AddIssues(workspaceSlug, projectID, moduleID, []string{issue.ID})
		assert.NoError(t, err)

		// Clean up
		// 清理
		err = s.Delete(workspaceSlug, projectID, moduleID)
		assert.NoError(t, err)
		err = issuesService.Delete(workspaceSlug, projectID, issue.ID)
		assert.NoError(t, err)
	})
}
