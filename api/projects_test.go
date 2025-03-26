package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestProjectsService tests all methods of the ProjectsService
// 测试 ProjectsService 的所有方法
func TestProjectsService(t *testing.T) {
	// Get API key from environment
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}

	// Get required IDs from environment
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")

	if workspaceSlug == "" {
		t.Skip("Required environment variables not set")
	}

	// Create a new client
	client := client.NewClient(apiKey)
	s := NewProjectsService(client)

	var projectID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		projects, err := s.List(workspaceSlug)
		assert.NoError(t, err)
		assert.NotNil(t, projects)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &ProjectCreateRequest{
			Name:        "Test Project",
			Identifier:  "TEST",
			Description: "Test Project Description",
		}
		project, err := s.Create(workspaceSlug, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.NotEmpty(t, project.ID)
		projectID = project.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		project, err := s.Get(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, projectID, project.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &ProjectUpdateRequest{
			Name:        "Updated Test Project",
			Description: "Updated Test Project Description",
		}
		project, err := s.Update(workspaceSlug, projectID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, projectID, project.ID)
		assert.Equal(t, updateReq.Name, project.Name)
		assert.Equal(t, updateReq.Description, project.Description)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID)
		assert.NoError(t, err)
	})
}
