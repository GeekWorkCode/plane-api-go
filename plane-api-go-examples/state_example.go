package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
)

// StateExample demonstrates the usage of state-related API endpoints
func StateExample(client *plane.Plane, workspaceSlug, projectID string) {
	fmt.Println("\n=== 测试状态相关接口 ===")

	// 列出所有状态
	states, err := client.States.List(workspaceSlug, projectID)
	if err != nil {
		log.Printf("获取状态列表失败: %v", err)
		return
	}
	fmt.Printf("状态列表: %+v\n", states)

	// 创建新状态
	timestamp := time.Now().Format("20060102150405")
	newState, err := client.States.Create(workspaceSlug, projectID, &api.StateCreateRequest{
		Name:        fmt.Sprintf("测试状态 %s", timestamp),
		Color:       "#FF0000",
		Description: "通过 API 创建的测试状态",
	})
	if err != nil {
		log.Printf("创建状态失败: %v", err)
		return
	}
	fmt.Printf("创建的状态: %+v\n", newState)

	// 获取状态详情
	state, err := client.States.Get(workspaceSlug, projectID, newState.ID)
	if err != nil {
		log.Printf("获取状态详情失败: %v", err)
		return
	}
	fmt.Printf("状态详情: %+v\n", state)

	// 更新状态
	updatedState, err := client.States.Update(workspaceSlug, projectID, newState.ID, &api.StateUpdateRequest{
		Name:        fmt.Sprintf("更新后的状态 %s", time.Now().Format("20060102150405")),
		Color:       "#00FF00",
		Description: "通过 API 更新的状态描述",
	})
	if err != nil {
		log.Printf("更新状态失败: %v", err)
		return
	}
	fmt.Printf("更新后的状态: %+v\n", updatedState)

	// 删除状态
	err = client.States.Delete(workspaceSlug, projectID, newState.ID)
	if err != nil {
		log.Printf("删除状态失败: %v", err)
		return
	}
	fmt.Println("状态删除成功")
}
