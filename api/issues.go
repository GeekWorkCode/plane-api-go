package api

import (
	"fmt"
	"net/http"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// IssuesService handles communication with the issue related endpoints
type IssuesService struct {
	client *client.Client
}

// NewIssuesService creates a new issues service
func NewIssuesService(client *client.Client) *IssuesService {
	return &IssuesService{
		client: client,
	}
}

// IssueCreateRequest represents the request body for creating an issue
type IssueCreateRequest struct {
	Name          string   `json:"name"`
	Description   string   `json:"description,omitempty"`
	State         string   `json:"state,omitempty"` // 状态ID
	StateName     string   `json:"-"`               // 状态名称 (不发送到API)
	Priority      string   `json:"priority,omitempty"`
	AssigneeID    string   `json:"-"`                   // 分配人ID (不直接发送到API)
	Assignees     []string `json:"assignees,omitempty"` // 多个分配人ID
	AssigneeNames []string `json:"-"`                   // 多个分配人名称 (不发送到API)
}

// IssueUpdateRequest represents the request body for updating an issue
type IssueUpdateRequest struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	State         string   `json:"state,omitempty"` // 状态ID
	StateName     string   `json:"-"`               // 状态名称 (不发送到API)
	Priority      string   `json:"priority,omitempty"`
	AssigneeID    string   `json:"-"`                   // 分配人ID (不直接发送到API)
	Assignees     []string `json:"assignees,omitempty"` // 多个分配人ID
	AssigneeNames []string `json:"-"`                   // 多个分配人名称 (不发送到API)
}

// List returns all issues in a project
func (s *IssuesService) List(workspaceSlug string, projectID string) ([]models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new(models.IssuesResponse)
	_, err = s.client.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// Get returns an issue by its ID
func (s *IssuesService) Get(workspaceSlug string, projectID string, issueID string) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// GetBySequenceID returns an issue by its sequence ID
func (s *IssuesService) GetBySequenceID(workspaceSlug string, sequenceID string) (*models.Issue, error) {
	path := fmt.Sprintf("/workspaces/%s/issues/%s/", workspaceSlug, sequenceID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// findStateIDByName 通过状态名称查找状态ID
func (s *IssuesService) findStateIDByName(workspaceSlug string, projectID string, stateName string) (string, error) {
	// 获取项目所有状态
	statesService := NewStatesService(s.client)
	states, err := statesService.List(workspaceSlug, projectID)
	if err != nil {
		return "", fmt.Errorf("获取状态列表失败: %w", err)
	}

	for _, state := range states {
		if state.Name == stateName {
			return state.ID, nil
		}
	}
	return "", fmt.Errorf("未找到名称为 '%s' 的状态", stateName)
}

// findMemberIDByName 通过成员名称查找成员ID
func (s *IssuesService) findMemberIDByName(workspaceSlug string, projectID string, memberName string) (string, error) {
	// 获取项目所有成员
	membersService := NewMembersService(s.client)
	members, err := membersService.List(workspaceSlug, projectID)
	if err != nil {
		return "", fmt.Errorf("获取成员列表失败: %w", err)
	}

	for _, member := range members {
		if member.Member.DisplayName == memberName {
			return member.Member.ID, nil
		}
	}
	return "", fmt.Errorf("未找到名称为 '%s' 的成员", memberName)
}

// Create creates a new issue
func (s *IssuesService) Create(workspaceSlug string, projectID string, createRequest *IssueCreateRequest) (*models.Issue, error) {
	// 如果提供了状态名称，查找对应的状态ID
	if createRequest.StateName != "" {
		stateID, err := s.findStateIDByName(workspaceSlug, projectID, createRequest.StateName)
		if err != nil {
			return nil, fmt.Errorf("查找状态失败: %w", err)
		}
		createRequest.State = stateID
	}

	// 如果提供了分配人名称，查找对应的成员ID
	if len(createRequest.AssigneeNames) > 0 {
		assigneeIDs := make([]string, 0, len(createRequest.AssigneeNames))
		for _, name := range createRequest.AssigneeNames {
			memberID, err := s.findMemberIDByName(workspaceSlug, projectID, name)
			if err != nil {
				return nil, fmt.Errorf("查找成员失败: %w", err)
			}
			assigneeIDs = append(assigneeIDs, memberID)
		}
		createRequest.Assignees = assigneeIDs
	}

	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/", workspaceSlug, projectID)
	req, err := s.client.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// Update updates an issue
func (s *IssuesService) Update(workspaceSlug string, projectID string, issueID string, updateRequest *IssueUpdateRequest) (*models.Issue, error) {
	// 如果提供了状态名称,查找对应的状态ID
	if updateRequest.StateName != "" && updateRequest.State == "" {
		stateID, err := s.findStateIDByName(workspaceSlug, projectID, updateRequest.StateName)
		if err != nil {
			return nil, err
		}
		updateRequest.State = stateID
	}

	// 如果提供了成员名称,查找对应的成员ID
	if len(updateRequest.AssigneeNames) > 0 && len(updateRequest.Assignees) == 0 {
		memberIDs := make([]string, 0, len(updateRequest.AssigneeNames))
		for _, memberName := range updateRequest.AssigneeNames {
			memberID, err := s.findMemberIDByName(workspaceSlug, projectID, memberName)
			if err != nil {
				return nil, err
			}
			memberIDs = append(memberIDs, memberID)
		}
		updateRequest.Assignees = memberIDs
	}

	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	issue := new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}

// Delete deletes an issue
func (s *IssuesService) Delete(workspaceSlug string, projectID string, issueID string) error {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// UpdateBySequenceID updates an issue by its sequence ID
func (s *IssuesService) UpdateBySequenceID(workspaceSlug string, sequenceID string, updateRequest *IssueUpdateRequest) (*models.Issue, error) {
	// 获取项目ID，通过序列ID获取项目ID
	issue, err := s.GetBySequenceID(workspaceSlug, sequenceID)
	if err != nil {
		return nil, fmt.Errorf("通过序列ID获取问题失败: %w", err)
	}
	projectID := issue.Project

	// 如果提供了状态名称,查找对应的状态ID
	if updateRequest.StateName != "" && updateRequest.State == "" {
		stateID, err := s.findStateIDByName(workspaceSlug, projectID, updateRequest.StateName)
		if err != nil {
			return nil, err
		}
		updateRequest.State = stateID
	}

	// 如果提供了成员名称,查找对应的成员ID
	if len(updateRequest.AssigneeNames) > 0 && len(updateRequest.Assignees) == 0 {
		memberIDs := make([]string, 0, len(updateRequest.AssigneeNames))
		for _, memberName := range updateRequest.AssigneeNames {
			memberID, err := s.findMemberIDByName(workspaceSlug, projectID, memberName)
			if err != nil {
				return nil, err
			}
			memberIDs = append(memberIDs, memberID)
		}
		updateRequest.Assignees = memberIDs
	}

	path := fmt.Sprintf("/workspaces/%s/issues/%s/", workspaceSlug, sequenceID)
	req, err := s.client.NewRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	issue = new(models.Issue)
	_, err = s.client.Do(req, issue)
	return issue, err
}
