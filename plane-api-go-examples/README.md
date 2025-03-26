# Plane API Go 示例项目

这是一个使用 Go 语言调用 Plane API 的完整示例项目，展示了如何使用 Plane API 进行项目管理、问题跟踪、周期管理、评论管理、工作日志记录和附件上传等操作。

## 功能特点

- 完整的 API 调用示例
- 支持环境变量配置
- 包含错误处理
- 中文注释说明
- 支持调试模式
- 所有API都确保URL格式正确，避免PATCH请求错误

## 环境要求

- Go 1.24.0 或更高版本
- Plane API 密钥
- Plane 工作空间和项目信息

## 安装

1. 
```
github.com/GeekWorkCode/plane-api-go v0.2.0
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
   - 更新问题信息
   - 删除问题

3. 周期管理
   - 列出所有周期
   - 创建新周期
   - 获取周期详情
   - 更新周期信息
   - 删除周期

4. 评论管理
   - 列出问题评论
   - 创建新评论
   - 获取评论详情
   - 更新评论内容
   - 删除评论

5. 工作日志管理
   - 记录问题工时
   - 列出工作日志
   - 更新工作日志
   - 获取项目总工时
   - 删除工作日志

6. 附件管理
   - 列出所有附件
   - 上传新附件
   - 获取上传凭证
   - 完成附件上传

## 注意事项

1. 请确保在使用前正确配置 `.env` 文件中的所有必要参数
2. 建议先在测试环境中运行示例程序
3. 示例程序会创建和删除测试数据，请谨慎使用
4. 如果启用调试模式，将会输出详细的 API 请求和响应信息
5. 对于PATCH和DELETE等请求，确保URL末尾有斜杠("/")，否则Django服务器可能无法处理请求

## 许可证

MIT 
