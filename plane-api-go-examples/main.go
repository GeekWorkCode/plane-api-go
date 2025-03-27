package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
	"github.com/GeekWorkCode/plane-api-go/models"
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
	// testProjects(client, workspaceSlug)

	// 测试问题相关接口
	// testIssues(client, workspaceSlug, projectID)

	// 测试周期相关接口
	// testCycles(client, workspaceSlug, projectID)

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

	// 通过序列ID获取问题
	// 需要先获取项目详情以获取项目标识符
	fmt.Println("\n=== 通过序列ID获取问题 ===")
	project, err := client.Projects.Get(workspaceSlug, projectID)
	if err != nil {
		log.Printf("获取项目详情失败: %v", err)
		return
	}

	// 构造序列ID (格式通常为: PROJECT_IDENTIFIER-SEQUENCE_NUMBER)
	// 注意: 我们需要从Issue响应中自行提取序列号
	// 这里假设问题的序列号可以从问题名称或其他信息中获取，或者使用现有问题的序列号
	// 在真实环境中，序列号通常是从问题ID或问题详情中获取
	sequenceID := ""
	// 尝试从已有问题中获取一个有效的序列ID
	if len(issues) > 0 {
		// 这里假设可以通过某种方式获取到序列ID
		// 在实际环境中可能需要通过API获取或从其他地方提取
		sequenceID = fmt.Sprintf("%s-%d", project.Identifier, 1) // 使用项目标识符和假设的序列号1
	} else {
		sequenceID = fmt.Sprintf("%s-%d", project.Identifier, 1) // 默认使用1作为序列号
	}

	issueBySequence, err := client.Issues.GetBySequenceID(workspaceSlug, sequenceID)
	if err != nil {
		log.Printf("通过序列ID获取问题失败: %v", err)
	} else {
		fmt.Printf("通过序列ID获取的问题: %+v\n", issueBySequence)

		// 使用此问题ID进行后续测试
		newIssue.ID = issueBySequence.ID
	}

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

	// 通过序列ID更新问题
	fmt.Println("\n=== 通过序列ID更新问题 ===")
	testUpdateIssueBySequenceID(client, workspaceSlug, sequenceID)

	// 删除问题
	// err = client.Issues.Delete(workspaceSlug, projectID, newIssue.ID)
	// if err != nil {
	// 	log.Printf("删除问题失败: %v", err)
	// 	return
	// }
	// fmt.Println("问题删除成功")
}

// 测试通过序列ID更新问题
func testUpdateIssueBySequenceID(client *plane.Plane, workspaceSlug, sequenceID string) {
	// 更新问题
	// TODO: 待实现 UpdateBySequenceID 方法的完整集成
	// 当前API库尚未完全导出此方法
	/*
		updatedIssue, err := client.Issues.UpdateBySequenceID(workspaceSlug, sequenceID, &api.IssueUpdateRequest{
			Name:        fmt.Sprintf("通过序列ID更新的问题 %s", time.Now().Format("20060102150405")),
			Description: "通过序列ID更新的问题描述",
			Priority:    "urgent",
		})
		if err != nil {
			log.Printf("通过序列ID更新问题失败: %v", err)
			return
		}
		fmt.Printf("通过序列ID更新后的问题: %+v\n", updatedIssue)
	*/

	// 临时替代方案：使用GetBySequenceID获取问题后，通过标准Update方法更新
	issueBySequence, err := client.Issues.GetBySequenceID(workspaceSlug, sequenceID)
	if err != nil {
		log.Printf("通过序列ID获取问题失败: %v", err)
		return
	}

	// 使用标准更新方法
	updatedIssue, err := client.Issues.Update(workspaceSlug, issueBySequence.Project, issueBySequence.ID, &api.IssueUpdateRequest{
		Name:        fmt.Sprintf("通过序列ID获取并更新的问题 %s", time.Now().Format("20060102150405")),
		Description: "通过序列ID获取并更新的问题描述",
		Priority:    "urgent",
	})
	if err != nil {
		log.Printf("更新问题失败: %v", err)
		return
	}
	fmt.Printf("通过序列ID获取并更新后的问题: %+v\n", updatedIssue)
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

	// 从成员列表中选择第一个成员的显示名称用于测试
	displayName := "zs"
	fmt.Printf("\n> 使用成员显示名称: %s 进行评论测试\n", displayName)

	var createdComment *models.Comment

	// 列出所有评论
	comments, err := client.Comments.List(workspaceSlug, projectID, issueID)
	if err != nil {
		log.Printf("获取评论列表失败: %v", err)
		return
	}

	// fmt.Printf("当前评论数量: %d\n", len(comments))

	// // 打印评论的成员信息
	if len(comments) > 0 {
		fmt.Println("\n现有评论的作者信息:")
		for i, comment := range comments {
			if comment.Member != nil {
				fmt.Printf("评论 %d: 由 %s (%s) 创建\n", i+1, comment.Member.DisplayName, comment.Member.ID)
			} else {
				fmt.Printf("评论 %d: 创建者ID %s (无成员信息)\n", i+1, comment.CreatedBy)
			}
		}
	}

	// 如果没有从评论中获取到成员ID，尝试从成员列表中获取
	// if memberID == "" {
	// 	// 获取成员列表
	// 	members, err := client.Members.List(workspaceSlug, projectID)
	// 	if err == nil && len(members) > 0 {
	// 		memberID = members[0].Member.ID
	// 		fmt.Printf("从成员列表获取到成员ID: %s\n", memberID)
	// 	}
	// }

	// 使用 DisplayName 创建评论
	fmt.Println("\n=== 使用 DisplayName 创建评论 ===")
	createReq := &api.CommentRequest{
		CommentHTML: fmt.Sprintf("<p>这是一条通过 DisplayName (%s) 创建的测试评论 %s</p>",
			displayName, time.Now().Format("20060102150405")),
		DisplayName: displayName,
	}

	createdComment, err = client.Comments.Create(workspaceSlug, projectID, issueID, createReq)
	if err != nil {
		log.Printf("使用 DisplayName 创建评论失败: %v", err)
		return
	}

	fmt.Println("评论创建成功！")

	// 显示评论的成员信息
	if createdComment.Member != nil {
		fmt.Printf("评论由 %s (%s) 创建\n", createdComment.Member.DisplayName, createdComment.Member.ID)
	} else {
		fmt.Printf("评论创建者ID: %s (无成员信息)\n", createdComment.CreatedBy)
	}

	// 获取评论详情
	comment, err := client.Comments.Get(workspaceSlug, projectID, issueID, createdComment.ID)
	if err != nil {
		log.Printf("获取评论详情失败: %v", err)
		return
	}
	fmt.Println("\n=== 获取评论详情 ===")
	fmt.Printf("评论ID: %s\n", comment.ID)
	fmt.Printf("评论内容: %s\n", comment.CommentHTML)
	fmt.Printf("创建时间: %s\n", comment.CreatedAt.Format("2006-01-02 15:04:05"))

	// 使用 DisplayName 更新评论
	// fmt.Println("\n=== 使用 DisplayName 更新评论 ===")
	// updateReq := &api.CommentRequest{
	// 	CommentHTML: fmt.Sprintf("<p>这是一条通过 DisplayName (%s) 更新的测试评论 %s</p>",
	// 		displayName, time.Now().Format("20060102150405")),
	// 	DisplayName: displayName,
	// }

	// updatedComment, err := client.Comments.Update(workspaceSlug, projectID, issueID, createdComment.ID, updateReq)
	// if err != nil {
	// 	log.Printf("使用 DisplayName 更新评论失败: %v", err)
	// } else {
	// 	fmt.Println("评论更新成功！")

	// 	// 显示评论的成员信息
	// 	if updatedComment.Member != nil {
	// 		fmt.Printf("评论由 %s (%s) 更新\n", updatedComment.Member.DisplayName, updatedComment.Member.ID)
	// 	} else {
	// 		fmt.Printf("评论更新者ID: %s (无成员信息)\n", updatedComment.UpdatedBy)
	// 	}
	// }

	// 使用 CreatedBy 直接创建评论
	// if memberID != "" {
	// 	fmt.Println("\n=== 使用 CreatedBy (MemberID) 创建评论 ===")
	// 	directCreateReq := &api.CommentRequest{
	// 		CommentHTML: fmt.Sprintf("<p>这是一条通过 CreatedBy (MemberID: %s) 创建的测试评论 %s</p>",
	// 			memberID, time.Now().Format("20060102150405")),
	// 		CreatedBy: memberID,
	// 	}

	// 	directComment, err := client.Comments.Create(workspaceSlug, projectID, issueID, directCreateReq)
	// 	if err != nil {
	// 		log.Printf("使用 CreatedBy 创建评论失败: %v", err)
	// 	} else {
	// 		fmt.Println("使用 CreatedBy 创建评论成功！")

	// 		// 显示评论的成员信息
	// 		if directComment.Member != nil {
	// 			fmt.Printf("评论由 %s (%s) 创建\n", directComment.Member.DisplayName, directComment.Member.ID)
	// 		} else {
	// 			fmt.Printf("评论创建者ID: %s (无成员信息)\n", directComment.CreatedBy)
	// 		}

	// 		// 删除通过 CreatedBy 创建的评论
	// 		fmt.Println("\n=== 删除通过 CreatedBy 创建的评论 ===")
	// 		err = client.Comments.Delete(workspaceSlug, projectID, issueID, directComment.ID)
	// 		if err != nil {
	// 			log.Printf("删除通过 CreatedBy 创建的评论失败: %v", err)
	// 		} else {
	// 			fmt.Println("通过 CreatedBy 创建的评论删除成功")
	// 		}
	// 	}
	// }

	// 删除创建的评论
	// fmt.Println("\n=== 删除通过 DisplayName 创建的评论 ===")
	// err = client.Comments.Delete(workspaceSlug, projectID, issueID, createdComment.ID)
	// if err != nil {
	// 	log.Printf("删除评论失败: %v", err)
	// } else {
	// 	fmt.Println("评论删除成功")
	// }

	// 再次列出所有评论，确认删除成功
	// comments, err = client.Comments.List(workspaceSlug, projectID, issueID)
	// if err != nil {
	// 	log.Printf("获取评论列表失败: %v", err)
	// 	return
	// }

	// fmt.Printf("\n删除后的评论数量: %d\n", len(comments))
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
