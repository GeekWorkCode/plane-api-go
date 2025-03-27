package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestMembersService tests all methods of the MembersService
// 测试 MembersService 的所有方法
func TestMembersService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewMembersService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")

	if workspaceSlug == "" || projectID == "" {
		t.Skip("Required environment variables not set")
	}

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		members, err := s.List(workspaceSlug, projectID)
		assert.NoError(t, err)
		assert.NotNil(t, members)

		// Verify at least one member exists
		if len(members) > 0 {
			// Test the Get method with the first member's ID
			memberID := members[0].ID

			// Test Get method
			// 测试 Get 方法
			t.Run("Get", func(t *testing.T) {
				member, err := s.Get(workspaceSlug, projectID, memberID)
				assert.NoError(t, err)
				assert.NotNil(t, member)
				assert.Equal(t, memberID, member.ID)

				// Verify the critical fields from the member model
				assert.NotEmpty(t, member.Member.ID)
				assert.NotEmpty(t, member.Member.DisplayName)
			})
		}
	})
}
