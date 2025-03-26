package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestStatesService tests all methods of the StatesService
// 测试 StatesService 的所有方法
func TestStatesService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewStatesService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	var stateID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		states, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, states)
	})

	// Test Create method
	// 测试 Create 方法
	t.Run("Create", func(t *testing.T) {
		createReq := &StateCreateRequest{
			Name:        "Test State",
			Description: "Test State Description",
			Color:       "#FF0000",
		}
		state, err := s.Create(workspaceSlug, projectID, createReq)
		assert.NoError(t, err)
		assert.NotNil(t, state)
		assert.NotEmpty(t, state.ID)
		stateID = state.ID
	})

	// Test Get method
	// 测试 Get 方法
	t.Run("Get", func(t *testing.T) {
		state, err := s.Get(workspaceSlug, projectID, stateID)
		assert.NoError(t, err)
		assert.NotNil(t, state)
		assert.Equal(t, stateID, state.ID)
	})

	// Test Update method
	// 测试 Update 方法
	t.Run("Update", func(t *testing.T) {
		updateReq := &StateUpdateRequest{
			Name:        "Updated Test State",
			Description: "Updated Test State Description",
			Color:       "#00FF00",
		}
		state, err := s.Update(workspaceSlug, projectID, stateID, updateReq)
		assert.NoError(t, err)
		assert.NotNil(t, state)
		assert.Equal(t, stateID, state.ID)
		assert.Equal(t, updateReq.Name, state.Name)
		assert.Equal(t, updateReq.Description, state.Description)
		assert.Equal(t, updateReq.Color, state.Color)
	})

	// Test Delete method
	// 测试 Delete 方法
	t.Run("Delete", func(t *testing.T) {
		err := s.Delete(workspaceSlug, projectID, stateID)
		assert.NoError(t, err)
	})
}
