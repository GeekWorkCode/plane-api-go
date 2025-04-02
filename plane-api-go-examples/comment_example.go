package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// CommentExample demonstrates the usage of comment-related API endpoints
func CommentExample(client *plane.Plane, workspaceSlug, projectID, issueID string) {
	fmt.Println("\n=== 测试评论相关接口 ===")

	// 从成员列表中选择第一个成员的显示名称用于测试
	displayName := "win8zhang"
	fmt.Printf("\n> 使用成员显示名称: %s 进行评论测试\n", displayName)

	var createdComment *models.Comment

	// 列出所有评论
	comments, err := client.Comments.List(workspaceSlug, projectID, issueID)
	if err != nil {
		log.Printf("获取评论列表失败: %v", err)
		return
	}

	// 打印评论的成员信息
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
	fmt.Println("\n=== 使用 DisplayName 更新评论 ===")
	updateReq := &api.CommentRequest{
		CommentHTML: fmt.Sprintf("<p>这是一条通过 DisplayName (%s) 更新的测试评论 %s</p>",
			displayName, time.Now().Format("20060102150405")),
		DisplayName: displayName,
	}

	updatedComment, err := client.Comments.Update(workspaceSlug, projectID, issueID, createdComment.ID, updateReq)
	if err != nil {
		log.Printf("使用 DisplayName 更新评论失败: %v", err)
	} else {
		fmt.Println("评论更新成功！")

		// 显示评论的成员信息
		if updatedComment.Member != nil {
			fmt.Printf("评论由 %s (%s) 更新\n", updatedComment.Member.DisplayName, updatedComment.Member.ID)
		} else {
			fmt.Printf("评论更新者ID: %s (无成员信息)\n", updatedComment.UpdatedBy)
		}
	}

	// 删除评论
	err = client.Comments.Delete(workspaceSlug, projectID, issueID, createdComment.ID)
	if err != nil {
		log.Printf("删除评论失败: %v", err)
		return
	}
	fmt.Println("评论删除成功")
}
