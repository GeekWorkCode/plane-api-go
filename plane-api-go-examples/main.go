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

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 获取环境变量
	apiKey := os.Getenv("PLANE_API_KEY")
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	issueID := os.Getenv("PLANE_ISSUE_ID")
	baseURL := os.Getenv("PLANE_API_BASE_URL")
	debug := os.Getenv("PLANE_API_DEBUG") == "true"

	// 创建 Plane 客户端
	client := plane.NewClient(apiKey)
	client.SetDebug(debug)
	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	// 如果PLANE_ISSUE_ID为空或为"your_issue_id"，则创建一个新问题
	if issueID == "" || issueID == "your_issue_id" {
		fmt.Println("\n=== 创建测试问题 ===")
		issue, err := client.Issues.Create(workspaceSlug, projectID, &api.IssueCreateRequest{
			Name:        fmt.Sprintf("测试问题 %s", time.Now().Format("20060102150405")),
			Description: "这是一个通过API创建的测试问题",
		})
		if err != nil {
			log.Printf("创建问题失败: %v", err)
		} else {
			issueID = issue.ID
			fmt.Printf("已创建问题，ID: %s\n", issueID)
			// 建议在测试完成后更新.env文件
			fmt.Println("请更新.env文件中的PLANE_ISSUE_ID为:", issueID)
		}
	}

	// 运行各个示例
	ProjectExample(client, workspaceSlug)
	IssueNameExample()
	StateExample(client, workspaceSlug, projectID)
	CycleExample(client, workspaceSlug, projectID)
	CommentExample(client, workspaceSlug, projectID, issueID)

}
