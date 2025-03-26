package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestLabelsService tests all methods of the LabelsService
// 测试 LabelsService 的所有方法
func TestLabelsService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewLabelsService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	var labelID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		labels, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, labels)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &LabelCreateRequest{
			Name:        "Test Label",
			Description: "Test label description",
			Color:       "#FF0000",
		}
		label, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, label)
		assert.NotEmpty(t, label.ID)
		labelID = label.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		label, err := s.Get(workspaceSlug, projectID, labelID)
		assert.NoError(t, err)
		assert.NotNil(t, label)
		assert.Equal(t, labelID, label.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &LabelUpdateRequest{
			Name:        "Updated Test Label",
			Description: "Updated test label description",
			Color:       "#00FF00",
		}
		label, err := s.Update(workspaceSlug, projectID, labelID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, label)
		assert.Equal(t, labelID, label.ID)
		assert.Equal(t, updateReq.Name, label.Name)
		assert.Equal(t, updateReq.Description, label.Description)
		assert.Equal(t, updateReq.Color, label.Color)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, labelID)
		assert.NoError(t, err)
	})
}
