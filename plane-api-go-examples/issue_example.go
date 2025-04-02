package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
	"github.com/joho/godotenv"
)

// 使用状态名称和分配人名称创建和更新问题的示例
func issueNameExample() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file, using environment variables")
	}

	// 获取环境变量
	apiKey := os.Getenv("PLANE_API_KEY")
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	baseURL := os.Getenv("PLANE_API_BASE_URL")

	// 创建客户端
	client := plane.NewClient(apiKey)
	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}
	client.SetDebug(true)

	fmt.Println("\n=== 获取所有状态和成员 ===")
	// 获取所有状态
	states, err := client.States.List(workspaceSlug, projectID)
	if err != nil {
		log.Fatalf("获取状态列表失败: %v", err)
	}
	fmt.Printf("当前项目有 %d 个状态:\n", len(states))
	for i, state := range states {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, state.Name, state.ID)
	}
	if len(states) == 0 {
		log.Fatalf("项目没有可用的状态，请先创建状态")
	}

	// 获取所有成员
	members, err := client.Members.List(workspaceSlug, projectID)
	if err != nil {
		log.Fatalf("获取成员列表失败: %v", err)
	}
	fmt.Printf("\n当前项目有 %d 个成员:\n", len(members))
	for i, member := range members {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, member.Member.DisplayName, member.Member.ID)
	}
	if len(members) == 0 {
		log.Fatalf("项目没有可用的成员")
	}

	// 使用第一个状态和第一个成员的名称
	stateName := states[0].Name //"Done"

	memberName := members[0].Member.DisplayName

	fmt.Printf("\n=== 使用状态名称「%s」和分配人名称「%s」创建问题 ===\n", stateName, memberName)

	// 创建问题 - 使用状态名称和分配人名称
	createReq := &api.IssueCreateRequest{
		Name:          fmt.Sprintf("使用状态名称和分配人名称创建的问题 %s", time.Now().Format("20060102150405")),
		Description:   "这是一个使用状态名称和分配人名称创建的测试问题",
		StateName:     stateName, // 使用状态名称而不是ID
		Priority:      "medium",
		AssigneeNames: []string{memberName}, // 使用分配人名称而不是ID
	}

	fmt.Printf("\n请求内容: %+v\n", createReq)

	newIssue, err := client.Issues.Create(workspaceSlug, projectID, createReq)
	if err != nil {
		log.Fatalf("创建问题失败: %v", err)
	}
	fmt.Printf("问题创建成功! ID: %s\n", newIssue.ID)
	fmt.Printf("状态ID: %s\n", newIssue.State)
	fmt.Printf("分配人列表: %v\n", newIssue.Assignees)

	// 打印调试信息
	fmt.Printf("\n=== 调试信息 ===\n")
	fmt.Printf("状态名称: %s, 状态ID: %s\n", stateName, newIssue.State)
	fmt.Printf("分配人名称: %s, 分配人列表: %v\n", memberName, newIssue.Assignees)

	// 查看原始API回应中是否包含AssigneeID
	fmt.Printf("\n尝试再次获取问题详情以检查分配人ID\n")
	issueDetails, err := client.Issues.Get(workspaceSlug, projectID, newIssue.ID)
	if err != nil {
		log.Printf("获取问题详情失败: %v", err)
	} else {
		fmt.Printf("从API获取的问题详情:\n")
		fmt.Printf("  ID: %s\n", issueDetails.ID)
		fmt.Printf("  名称: %s\n", issueDetails.Name)
		fmt.Printf("  状态ID: %s\n", issueDetails.State)
		fmt.Printf("  分配人列表: %v\n", issueDetails.Assignees)
	}

	// 使用另一个状态名称和分配人名称更新问题
	var updateStateName, updateMemberName string
	if len(states) > 1 {
		updateStateName = states[1].Name
	} else {
		updateStateName = states[0].Name
	}

	if len(members) > 1 {
		updateMemberName = members[1].Member.DisplayName
	} else {
		updateMemberName = members[0].Member.DisplayName
	}

	fmt.Printf("\n=== 使用状态名称「%s」和分配人名称「%s」更新问题 ===\n", updateStateName, updateMemberName)

	// 更新问题 - 使用状态名称和分配人名称
	updateReq := &api.IssueUpdateRequest{
		Name:          fmt.Sprintf("使用状态名称和分配人名称更新的问题 %s", time.Now().Format("20060102150405")),
		Description:   "这是一个使用状态名称和分配人名称更新的测试问题",
		StateName:     updateStateName, // 使用状态名称而不是ID
		Priority:      "high",
		AssigneeNames: []string{updateMemberName}, // 使用分配人名称而不是ID
	}

	updatedIssue, err := client.Issues.Update(workspaceSlug, projectID, newIssue.ID, updateReq)
	if err != nil {
		log.Fatalf("更新问题失败: %v", err)
	}
	fmt.Printf("问题更新成功! ID: %s\n", updatedIssue.ID)
	fmt.Printf("新状态ID: %s\n", updatedIssue.State)
	fmt.Printf("新分配人列表: %v\n", updatedIssue.Assignees)

	// 使用多个分配人名称创建问题
	fmt.Println("\n=== 使用多个分配人名称创建问题 ===")

	// 收集所有成员的名称
	var memberNames []string
	for i, member := range members {
		if i < 3 { // 最多使用3个成员
			memberNames = append(memberNames, member.Member.DisplayName)
		}
	}

	fmt.Printf("使用以下成员名称: %v\n", memberNames)

	// 创建问题 - 使用多个分配人名称
	multiAssigneeReq := &api.IssueCreateRequest{
		Name:          fmt.Sprintf("使用多个分配人创建的问题 %s", time.Now().Format("20060102150405")),
		Description:   "这是一个使用多个分配人的测试问题",
		StateName:     stateName, // 使用状态名称
		Priority:      "medium",
		AssigneeNames: memberNames, // 使用多个分配人名称
	}

	multiAssigneeIssue, err := client.Issues.Create(workspaceSlug, projectID, multiAssigneeReq)
	if err != nil {
		log.Fatalf("创建多分配人问题失败: %v", err)
	}
	fmt.Printf("多分配人问题创建成功! ID: %s\n", multiAssigneeIssue.ID)
	fmt.Printf("分配人列表: %v\n", multiAssigneeIssue.Assignees)

	fmt.Println("\n=== 示例完成 ===")
}

// IssueNameExample exports the functionality for use by main.go
func IssueNameExample() {
	issueNameExample()
}
