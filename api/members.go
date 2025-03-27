package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// MembersService handles communication with the project members related endpoints
type MembersService struct {
	client *client.Client
}

// NewMembersService creates a new members service
func NewMembersService(client *client.Client) *MembersService {
	return &MembersService{
		client: client,
	}
}

// MemberUserResponse represents the simplified user response from the members endpoint
type MemberUserResponse struct {
	ID          string      `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	Avatar      string      `json:"avatar"`
	AvatarURL   interface{} `json:"avatar_url"`
	DisplayName string      `json:"display_name"`
}

// List returns all members for a project
func (s *MembersService) List(workspaceSlug string, projectID string) ([]models.Member, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/members/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// First, get the raw response as a slice of simplified users
	var memberUsers []MemberUserResponse
	_, err = s.client.Do(req, &memberUsers)
	if err != nil {
		return nil, fmt.Errorf("获取成员列表失败: %w", err)
	}

	// Convert the simplified users to full member models
	members := make([]models.Member, len(memberUsers))
	for i, user := range memberUsers {
		// Create a member model for each user
		member := models.Member{
			// For the first version, we'll just use the user ID as the member ID
			// In a real implementation, you'd want to fetch the full member details
			ID: user.ID,
			Member: models.MemberUser{
				ID:              user.ID,
				FirstName:       user.FirstName,
				LastName:        user.LastName,
				Avatar:          user.Avatar,
				AvatarURL:       user.AvatarURL,
				DisplayName:     user.DisplayName,
				Email:           user.Email,
				LastLoginMedium: "", // Not available in the simple response
			},
			// Other fields would be filled in from a detailed response
			Role:     20, // Default role as member
			IsActive: true,
		}
		members[i] = member
	}

	return members, nil
}

// Get returns a member by its ID
func (s *MembersService) Get(workspaceSlug string, projectID string, memberID string) (*models.Member, error) {
	// Since direct member lookup is returning 404, we'll use the List method to get all members
	// and then filter to find the one with the requested ID
	members, err := s.List(workspaceSlug, projectID)
	if err != nil {
		return nil, fmt.Errorf("获取成员列表失败: %w", err)
	}

	// Find the requested member
	for _, member := range members {
		if member.Member.ID == memberID {
			return &member, nil
		}
	}

	return nil, fmt.Errorf("成员未找到: %s", memberID)
}
