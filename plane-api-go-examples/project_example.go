package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
)

// ProjectExample demonstrates the usage of project-related API endpoints
func ProjectExample(client *plane.Plane, workspaceSlug string) {
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
		Description: "通过 API 更新的项目描述",
	})
	if err != nil {
		log.Printf("更新项目失败: %v", err)
		return
	}
	fmt.Printf("更新后的项目: %+v\n", updatedProject)

	// 删除项目
	err = client.Projects.Delete(workspaceSlug, newProject.ID)
	if err != nil {
		log.Printf("删除项目失败: %v", err)
		return
	}
	fmt.Println("项目删除成功")
}
