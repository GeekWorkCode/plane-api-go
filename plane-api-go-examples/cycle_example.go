package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GeekWorkCode/plane-api-go"
	"github.com/GeekWorkCode/plane-api-go/api"
)

// CycleExample demonstrates the usage of cycle-related API endpoints
func CycleExample(client *plane.Plane, workspaceSlug, projectID string) {
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
	err = client.Cycles.Delete(workspaceSlug, projectID, newCycle.ID)
	if err != nil {
		log.Printf("删除周期失败: %v", err)
		return
	}
	fmt.Println("周期删除成功")
}
