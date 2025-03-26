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

	// 测试项目相关接口
	testProjects(client, workspaceSlug)

	// 测试问题相关接口
	testIssues(client, workspaceSlug, projectID)

	// 测试周期相关接口
	testCycles(client, workspaceSlug, projectID)

	// 测试评论相关接口
	testComments(client, workspaceSlug, projectID, issueID)

	// 测试工作日志功能 - API可能尚未实现
	// fmt.Println("\n=== 测试工作日志 ===")
	// testWorklogs(client, workspaceSlug, projectID, issueID)

	// 附件相关接口需要特殊配置，暂不测试
	fmt.Println("\n=== 测试附件相关接口 ===")
	fmt.Println("附件上传功能需要更多配置和测试，略过演示。")
}

// 测试项目相关接口
func testProjects(client *plane.Plane, workspaceSlug string) {
	fmt.Println("\n=== 测试项目相关接口 ===")

	// 列出所有项目
	projects, err := client.Projects.List(workspaceSlug)
	if err != nil {
		log.Printf("获取项目列表失败: %v", err)
		return
	}
	fmt.Printf("项目列表: %+v\n", projects)

	// 创建新项目
	timestamp := time.Now().Format("20060102150405")
	newProject, err := client.Projects.Create(workspaceSlug, &api.ProjectCreateRequest{
		Name:        fmt.Sprintf("测试项目 %s", timestamp),
		Identifier:  fmt.Sprintf("T%s", timestamp[8:]), // 使用时分秒部分作为标识符的一部分
		Description: "通过 API 创建的测试项目",
	})
	if err != nil {
		log.Printf("创建项目失败: %v", err)
		return
	}
	fmt.Printf("创建的项目: %+v\n", newProject)

	// 获取项目详情
	project, err := client.Projects.Get(workspaceSlug, newProject.ID)
	if err != nil {
		log.Printf("获取项目详情失败: %v", err)
		return
	}
	fmt.Printf("项目详情: %+v\n", project)

	// 更新项目
	updatedProject, err := client.Projects.Update(workspaceSlug, newProject.ID, &api.ProjectUpdateRequest{
		Name:        fmt.Sprintf("更新后的项目 %s", time.Now().Format("20060102150405")),
		Description: "更新后的项目描述",
	})
	if err != nil {
		log.Printf("更新项目失败: %v", err)
		return
	}
	fmt.Printf("更新后的项目: %+v\n", updatedProject)

	// 删除项目
	// err = client.Projects.Delete(workspaceSlug, newProject.ID)
	// if err != nil {
	// 	log.Printf("删除项目失败: %v", err)
	// 	return
	// }
	// fmt.Println("项目删除成功")
}

// 测试问题相关接口
func testIssues(client *plane.Plane, workspaceSlug, projectID string) {
	fmt.Println("\n=== 测试问题相关接口 ===")

	// 列出所有问题
	issues, err := client.Issues.List(workspaceSlug, projectID)
	if err != nil {
		log.Printf("获取问题列表失败: %v", err)
		return
	}
	fmt.Printf("问题列表: %+v\n", issues)

	// 获取状态ID
	stateID := ""
	if len(issues) > 0 {
		stateID = issues[0].State
	} else {
		log.Printf("无法找到状态ID，跳过创建问题")
		return
	}

	// 创建新问题
	newIssue, err := client.Issues.Create(workspaceSlug, projectID, &api.IssueCreateRequest{
		Name:        fmt.Sprintf("测试问题 %s", time.Now().Format("20060102150405")),
		Description: "通过 API 创建的测试问题",
		State:       stateID,
		Priority:    "medium",
	})
	if err != nil {
		log.Printf("创建问题失败: %v", err)
		return
	}
	fmt.Printf("创建的问题: %+v\n", newIssue)

	// 获取问题详情
	issue, err := client.Issues.Get(workspaceSlug, projectID, newIssue.ID)
	if err != nil {
		log.Printf("获取问题详情失败: %v", err)
		return
	}
	fmt.Printf("问题详情: %+v\n", issue)

	// 更新问题
	updatedIssue, err := client.Issues.Update(workspaceSlug, projectID, newIssue.ID, &api.IssueUpdateRequest{
		Name:        fmt.Sprintf("更新后的问题 %s", time.Now().Format("20060102150405")),
		Description: "更新后的问题描述",
		State:       stateID,
		Priority:    "high",
	})
	if err != nil {
		log.Printf("更新问题失败: %v", err)
		return
	}
	fmt.Printf("更新后的问题: %+v\n", updatedIssue)

	// 删除问题
	// err = client.Issues.Delete(workspaceSlug, projectID, newIssue.ID)
	// if err != nil {
	// 	log.Printf("删除问题失败: %v", err)
	// 	return
	// }
	// fmt.Println("问题删除成功")
}

// 测试周期相关接口
func testCycles(client *plane.Plane, workspaceSlug, projectID string) {
	fmt.Println("\n=== 测试周期相关接口 ===")

	// 列出所有周期
	cycles, err := client.Cycles.List(workspaceSlug, projectID)
	if err != nil {
		log.Printf("获取周期列表失败: %v", err)
		return
	}
	fmt.Printf("周期列表: %+v\n", cycles)

	// 创建新周期
	fmt.Println("开始创建周期...")
	cycleRequest := &api.CycleCreateRequest{
		Name: fmt.Sprintf("测试周期 %s", time.Now().Format("20060102150405")),
	}
	fmt.Printf("周期创建请求: %+v\n", cycleRequest)

	newCycle, err := client.Cycles.Create(workspaceSlug, projectID, cycleRequest)
	if err != nil {
		log.Printf("创建周期失败: %v", err)
		return
	}
	fmt.Printf("创建的周期: %+v\n", newCycle)

	// 获取周期详情
	cycle, err := client.Cycles.Get(workspaceSlug, projectID, newCycle.ID)
	if err != nil {
		log.Printf("获取周期详情失败: %v", err)
		return
	}
	fmt.Printf("周期详情: %+v\n", cycle)

	// 更新周期
	updatedCycle, err := client.Cycles.Update(workspaceSlug, projectID, newCycle.ID, &api.CycleUpdateRequest{
		Name: fmt.Sprintf("更新后的周期 %s", time.Now().Format("20060102150405")),
	})
	if err != nil {
		log.Printf("更新周期失败: %v", err)
		return
	}
	fmt.Printf("更新后的周期: %+v\n", updatedCycle)

	// 删除周期
	// err = client.Cycles.Delete(workspaceSlug, projectID, newCycle.ID)
	// if err != nil {
	// 	log.Printf("删除周期失败: %v", err)
	// 	return
	// }
	// fmt.Println("周期删除成功")
}

// 测试评论相关接口
func testComments(client *plane.Plane, workspaceSlug, projectID, issueID string) {
	fmt.Println("\n=== 测试评论相关接口 ===")

	// 列出所有评论
	comments, err := client.Comments.List(workspaceSlug, projectID, issueID)
	if err != nil {
		log.Printf("获取评论列表失败: %v", err)
		return
	}
	fmt.Printf("评论列表: %+v\n", comments)

	// 创建新评论
	newComment, err := client.Comments.Create(workspaceSlug, projectID, issueID, &api.CommentCreateRequest{
		CommentHTML: fmt.Sprintf("<p>这是一条通过 API 创建的测试评论 %s</p>", time.Now().Format("20060102150405")),
	})
	if err != nil {
		log.Printf("创建评论失败: %v", err)
		return
	}
	fmt.Printf("创建的评论: %+v\n", newComment)

	// 获取评论详情
	comment, err := client.Comments.Get(workspaceSlug, projectID, issueID, newComment.ID)
	if err != nil {
		log.Printf("获取评论详情失败: %v", err)
		return
	}
	fmt.Printf("评论详情: %+v\n", comment)

	// 更新评论
	updatedComment, err := client.Comments.Update(workspaceSlug, projectID, issueID, newComment.ID, &api.CommentUpdateRequest{
		CommentHTML: fmt.Sprintf("<p>这是一条更新后的评论 %s</p>", time.Now().Format("20060102150405")),
	})
	if err != nil {
		log.Printf("更新评论失败: %v", err)
		return
	}
	fmt.Printf("更新后的评论: %+v\n", updatedComment)

	// 删除评论
	err = client.Comments.Delete(workspaceSlug, projectID, issueID, newComment.ID)
	if err != nil {
		log.Printf("删除评论失败: %v", err)
		return
	}
	fmt.Println("评论删除成功")
}

// 测试工作日志相关接口
func testWorklogs(client *plane.Plane, workspaceSlug, projectID, issueID string) {
	fmt.Println("\n=== 测试工作日志相关接口 ===")

	// 列出所有工作日志
	worklogs, err := client.Worklogs.List(workspaceSlug, projectID, issueID)
	if err != nil {
		log.Printf("获取工作日志列表失败: %v", err)
		return
	}
	fmt.Printf("工作日志列表: %+v\n", worklogs)

	// 创建新工作日志
	newWorklog, err := client.Worklogs.Create(workspaceSlug, projectID, issueID, &api.WorklogCreateRequest{
		Description: fmt.Sprintf("测试工作日志 %s", time.Now().Format("20060102150405")),
		Duration:    60, // 60分钟
	})
	if err != nil {
		log.Printf("创建工作日志失败: %v", err)
		return
	}
	fmt.Printf("创建的工作日志: %+v\n", newWorklog)

	// 获取工作日志详情
	worklog, err := client.Worklogs.Get(workspaceSlug, projectID, issueID, newWorklog.ID)
	if err != nil {
		log.Printf("获取工作日志详情失败: %v", err)
		return
	}
	fmt.Printf("工作日志详情: %+v\n", worklog)

	// 更新工作日志
	updatedWorklog, err := client.Worklogs.Update(workspaceSlug, projectID, issueID, newWorklog.ID, &api.WorklogUpdateRequest{
		Description: fmt.Sprintf("更新后的工作日志 %s", time.Now().Format("20060102150405")),
		Duration:    90, // 90分钟
	})
	if err != nil {
		log.Printf("更新工作日志失败: %v", err)
		return
	}
	fmt.Printf("更新后的工作日志: %+v\n", updatedWorklog)

	// 获取项目总工时
	totalTime, err := client.Worklogs.GetTotalTime(workspaceSlug, projectID)
	if err != nil {
		log.Printf("获取项目总工时失败: %v", err)
	} else {
		fmt.Printf("项目总工时: %+v\n", totalTime)
	}

	// 删除工作日志
	// err = client.Worklogs.Delete(workspaceSlug, projectID, issueID, newWorklog.ID)
	// if err != nil {
	// 	log.Printf("删除工作日志失败: %v", err)
	// 	return
	// }
	// fmt.Println("工作日志删除成功")
}
