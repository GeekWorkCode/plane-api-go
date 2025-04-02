# Plane API Go 示例项目

这是一个使用 Go 语言调用 Plane API 的完整示例项目，展示了如何使用 Plane API 进行项目管理、问题跟踪、周期管理、评论管理、工作日志记录和附件上传等操作。

## 功能特点

- 完整的 API 调用示例
- 支持环境变量配置
- 包含错误处理
- 中文注释说明
- 支持调试模式
- 所有API都确保URL格式正确，避免PATCH请求错误
- 支持通过序列ID（SequenceID）操作问题
- 支持通过分配人名称（AssigneeName）而非ID创建和更新问题
- 支持多分配人（Multiple Assignees）功能
- 支持灵活的评论创建方式，可通过成员显示名称（DisplayName）或成员ID（CreatedBy）创建评论
- 自动处理评论作者显示不一致的问题，确保评论显示正确的作者
- 支持状态（State）管理，可创建、更新和删除问题状态

## 环境要求

- Go 1.24.0 或更高版本
- Plane API 密钥
- Plane 工作空间和项目信息

## 安装

1. 确保 `plane-api-go` 库位于正确位置：
```bash
# 示例项目应该位于 plane-api-go 库的同级目录下
# 目录结构如下：
# - plane-api-go/         # API 库
# - plane-api-go-examples/ # 示例项目
```

2. 安装依赖：
```bash
go mod tidy
```

3. 配置环境变量：
编辑 `.env` 文件，填入你的配置信息：
```
PLANE_API_KEY=your_api_key_here
PLANE_WORKSPACE_SLUG=your_workspace_slug
PLANE_PROJECT_ID=your_project_id
PLANE_ISSUE_ID=your_issue_id
PLANE_API_BASE_URL=https://api.plane.so/api/v1
PLANE_API_DEBUG=true
```

## 使用方法

运行示例程序：
```bash
go run main.go
```

## 示例功能

1. 项目管理
   - 列出所有项目
   - 创建新项目
   - 获取项目详情
   - 更新项目信息
   - 删除项目

2. 问题管理
   - 列出所有问题
   - 创建新问题
   - 获取问题详情
   - 通过序列ID获取问题详情
   - 更新问题信息
   - 通过序列ID更新问题
   - 删除问题

3. 状态管理
   - 列出项目中的所有状态
   - 创建自定义状态
   - 获取状态详情
   - 更新状态（名称、描述、颜色）
   - 删除状态
   - 验证状态在问题中的使用

4. 周期管理
   - 列出所有周期
   - 创建新周期
   - 获取周期详情
   - 更新周期信息
   - 删除周期

5. 评论管理
   - 列出问题评论
   - 创建新评论（支持两种方式）：
     - 通过成员显示名称（DisplayName）
     - 通过成员ID（CreatedBy）
   - 获取评论详情
   - 更新评论内容（支持两种方式）：
     - 通过成员显示名称（DisplayName）
     - 通过成员ID（Actor）
   - 删除评论

6. 工作日志管理
   - 记录问题工时
   - 列出工作日志
   - 更新工作日志
   - 获取项目总工时
   - 删除工作日志

7. 附件管理
   - 列出所有附件
   - 上传新附件
   - 获取上传凭证
   - 完成附件上传

## 问题管理功能

### 通过状态名称和分配人名称创建问题

Plane API 支持通过状态的名称和分配人的名称来创建问题，无需预先知道它们的 ID：

```go
// 使用状态名称和分配人名称创建问题
createReq := &api.IssueCreateRequest{
    Name:          "通过状态名称和分配人名称创建的问题",
    Description:   "这是通过API使用状态名称和分配人名称创建的问题",
    StateName:     "Backlog",    // 使用状态名称，而不是ID
    Priority:      "medium",
    AssigneeNames: []string{"win8zhang"},  // 使用分配人名称，而不是ID
}

newIssue, err := client.Issues.Create(workspaceSlug, projectID, createReq)
```

在这个示例中，API 客户端将自动：
1. 查找名为 "Backlog" 的状态并获取其 ID
2. 查找显示名称为 "win8zhang" 的成员并获取其 ID
3. 使用这些 ID 创建问题

### 创建带有多个分配人的问题

Plane 支持为一个问题分配多个处理人。您可以通过提供一个名称数组来实现：

```go
// 使用多个分配人名称创建问题
memberNames := []string{"win8zhang", "zs"}
multiAssigneeReq := &api.IssueCreateRequest{
    Name:          "带有多个分配人的问题",
    Description:   "这是通过API使用多个分配人名称创建的问题",
    StateName:     "Backlog",
    Priority:      "medium",
    AssigneeNames: memberNames, // 使用多个分配人名称
}

multiAssigneeIssue, err := client.Issues.Create(workspaceSlug, projectID, multiAssigneeReq)
```

同样，API 客户端会自动查找这些名称对应的成员 ID，并将它们分配给问题。

## 评论功能说明

评论功能是与 Plane 系统交互的重要部分，我们对评论系统进行了增强，以提供更灵活的使用方式。

### 创建评论的两种方式

1. 通过成员显示名称（DisplayName）创建评论:
   ```go
   commentReq := &api.CommentRequest{
       CommentHTML: "<p>这是一条通过显示名称创建的评论</p>",
       DisplayName: "用户显示名称",
   }
   
   comment, err := client.Comments.Create(workspaceSlug, projectID, issueID, commentReq)
   ```

   使用这种方式时，库会自动查找匹配的成员ID，无需预先知道ID值。

2. 通过成员ID（CreatedBy）直接创建评论:
   ```go
   commentReq := &api.CommentRequest{
       CommentHTML: "<p>这是一条通过成员ID创建的评论</p>",
       CreatedBy: "成员ID",
   }
   
   comment, err := client.Comments.Create(workspaceSlug, projectID, issueID, commentReq)
   ```

### 评论作者显示处理

Plane API在处理评论创建时有一个特殊行为：有时在创建评论后，API返回的评论对象可能显示的作者不是预期的用户，这会导致UI中评论显示为系统管理员创建而非实际用户。

我们的库中实现了自动修正机制，在创建评论后检测并处理这种情况：

1. 当创建评论时，库会保存原始的创建者ID
2. 检查API返回的评论对象中显示的作者是否与提供的创建者匹配
3. 如果发现不匹配，库会自动执行一次更新操作，确保评论显示正确的作者
4. 这个过程对用户完全透明，无需额外代码

示例代码中的`testComments`函数展示了这两种评论创建方式，并会显示评论作者的详细信息进行验证。

## 状态管理功能

Plane 中的状态用于跟踪问题的进展情况（例如：待处理、进行中、已完成等）。示例代码展示了如何管理这些状态：

1. 创建自定义状态：
   ```go
   createReq := &api.StateCreateRequest{
       Name:        "测试状态",
       Description: "通过 API 创建的测试状态",
       Color:       "#4CAF50", // 绿色
   }
   
   newState, err := client.States.Create(workspaceSlug, projectID, createReq)
   ```

2. 更新状态属性：
   ```go
   updateReq := &api.StateUpdateRequest{
       Name:        "更新后的状态",
       Description: "通过 API 更新的测试状态",
       Color:       "#FFC107", // 黄色
   }
   
   updatedState, err := client.States.Update(workspaceSlug, projectID, stateID, updateReq)
   ```

3. 使用自定义状态创建问题：
   ```go
   issueReq := &api.IssueCreateRequest{
       Name:        "使用新状态的测试问题",
       Description: "这是一个使用新创建状态的测试问题",
       State:       newState.ID,
       Priority:    "medium",
   }
   
   newIssue, err := client.Issues.Create(workspaceSlug, projectID, issueReq)
   ```

示例代码中的`testStates`函数提供了完整的状态管理流程演示。

### 优势

- 更加灵活的接口，同时支持显示名称和成员ID
- 自动处理成员ID查找，简化开发
- 自动处理作者显示问题，确保UI中显示正确的评论作者
- 统一的API接口使代码更加清晰

## 注意事项

1. 请确保在使用前正确配置 `.env` 文件中的所有必要参数
2. 建议先在测试环境中运行示例程序
3. 示例程序会创建和删除测试数据，请谨慎使用
4. 如果启用调试模式，将会输出详细的 API 请求和响应信息
5. 对于PATCH和DELETE等请求，确保URL末尾有斜杠("/")，否则Django服务器可能无法处理请求
6. 序列ID（SequenceID）通常是"PROJECT_IDENTIFIER-NUMBER"格式，例如"PRJ-123"
7. 使用序列ID操作问题时，需要确保序列ID是正确的，否则API将返回404错误
8. 评论创建时如果遇到作者显示不正确的情况，库会自动尝试修正，无需手动处理
9. 如果某个状态正在被问题使用，可能无法删除，API会返回相应的错误信息
10. 使用分配人名称（AssigneeName）时，确保提供的名称在系统中存在并且拼写正确
11. 问题的分配人信息在API响应中以`assignees`数组形式返回，而不是单个`assignee_id`字段
12. 当需要获取分配人信息时，请使用`issue.Assignees`数组而不是`issue.AssigneeID`字段

## 许可证

MIT 
