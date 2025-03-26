package api

import (
	"os"
	"testing"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/stretchr/testify/assert"
)

// TestAttachmentsService tests all methods of the AttachmentsService
// 测试 AttachmentsService 的所有方法
func TestAttachmentsService(t *testing.T) {
	// Initialize test client
	// 初始化测试客户端
	apiKey := os.Getenv("PLANE_API_KEY")
	if apiKey == "" {
		t.Skip("PLANE_API_KEY not set")
	}
	c := client.NewClient(apiKey)
	s := NewAttachmentsService(c)

	// Test data
	// 测试数据
	workspaceSlug := os.Getenv("PLANE_WORKSPACE_SLUG")
	projectID := os.Getenv("PLANE_PROJECT_ID")
	issueID := os.Getenv("PLANE_ISSUE_ID")

	if workspaceSlug == "" || projectID == "" || issueID == "" {
		t.Skip("Required environment variables not set")
	}

	var attachmentID string

	// Test List method
	// 测试 List 方法
	t.Run("List", func(t *testing.T) {
		attachments, err := s.List(workspaceSlug, projectID, issueID)
		assert.NoError(t, err)
		assert.NotNil(t, attachments)
	})

	// Test GetUploadCredentials method
	// 测试 GetUploadCredentials 方法
	t.Run("GetUploadCredentials", func(t *testing.T) {
		credentials, err := s.GetUploadCredentials(workspaceSlug, projectID, issueID, "test.txt", "text/plain", 100)
		assert.NoError(t, err)
		assert.NotNil(t, credentials)
		assert.NotEmpty(t, credentials.AssetID)
		assert.NotNil(t, credentials.UploadData)
		assert.NotEmpty(t, credentials.UploadData.URL)
		assert.NotNil(t, credentials.UploadData.Fields)
	})

	// Test CompleteUpload method
	// 测试 CompleteUpload 方法
	t.Run("CompleteUpload", func(t *testing.T) {
		// First get upload credentials
		// 首先获取上传凭证
		credentials, err := s.GetUploadCredentials(workspaceSlug, projectID, issueID, "test.txt", "text/plain", 100)
		assert.NoError(t, err)

		// Complete the upload
		// 完成上传
		attachment, err := s.CompleteUpload(workspaceSlug, projectID, issueID, credentials.AssetID)
		assert.NoError(t, err)
		assert.NotNil(t, attachment)
		assert.NotEmpty(t, attachment.ID)
		attachmentID = attachment.ID
	})

	// Test UploadFile method
	// 测试 UploadFile 方法
	t.Run("UploadFile", func(t *testing.T) {
		// Create a test file
		// 创建一个测试文件
		testFile := "test_upload.txt"
		err := os.WriteFile(testFile, []byte("Test file content"), 0644)
		assert.NoError(t, err)
		defer os.Remove(testFile)

		// Get upload credentials
		// 获取上传凭证
		credentials, err := s.GetUploadCredentials(workspaceSlug, projectID, issueID, testFile, "text/plain", 100)
		assert.NoError(t, err)

		// Upload the file
		// 上传文件
		err = s.UploadFile(credentials.UploadData.URL, credentials.UploadData.Fields, testFile)
		assert.NoError(t, err)

		// Complete the upload
		// 完成上传
		attachment, err := s.CompleteUpload(workspaceSlug, projectID, issueID, credentials.AssetID)
		assert.NoError(t, err)
		assert.NotNil(t, attachment)
		assert.NotEmpty(t, attachment.ID)
		attachmentID = attachment.ID
	})

	// Test UploadFileToIssue method
	// 测试 UploadFileToIssue 方法
	t.Run("UploadFileToIssue", func(t *testing.T) {
		// Create a test file
		// 创建一个测试文件
		testFile := "test_upload_to_issue.txt"
		err := os.WriteFile(testFile, []byte("Test file content for issue"), 0644)
		assert.NoError(t, err)
		defer os.Remove(testFile)

		// Upload the file to the issue
		// 将文件上传到问题中
		attachment, err := s.UploadFileToIssue(workspaceSlug, projectID, issueID, testFile)
		assert.NoError(t, err)
		assert.NotNil(t, attachment)
		assert.NotEmpty(t, attachment.ID)
		attachmentID = attachment.ID
	})

	// Clean up
	// 清理
	if attachmentID != "" {
		// Note: The API might not provide a direct method to delete attachments
		// 注意：API 可能不提供直接删除附件的方法
	}
}
