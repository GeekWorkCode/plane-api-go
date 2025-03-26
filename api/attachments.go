package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// AttachmentsService handles communication with the issue attachment related endpoints
type AttachmentsService struct {
	client *client.Client
}

// NewAttachmentsService creates a new attachments service
func NewAttachmentsService(client *client.Client) *AttachmentsService {
	return &AttachmentsService{
		client: client,
	}
}

// UploadCredentialsRequest represents the request body for getting upload credentials
type UploadCredentialsRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

// List returns all attachments for an issue
func (s *AttachmentsService) List(workspaceSlug string, projectID string, issueID string) ([]models.Attachment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/issue-attachments/", workspaceSlug, projectID, issueID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var attachments []models.Attachment
	_, err = s.client.Do(req, &attachments)
	if err != nil {
		return nil, err
	}

	return attachments, nil
}

// GetUploadCredentials gets credentials to upload a file directly to cloud storage
func (s *AttachmentsService) GetUploadCredentials(workspaceSlug string, projectID string, issueID string, filename string, fileType string, fileSize int64) (*models.UploadCredentials, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/issue-attachments/get-upload-url/", workspaceSlug, projectID, issueID)

	requestBody := &UploadCredentialsRequest{
		Name: filename,
		Type: fileType,
		Size: fileSize,
	}

	req, err := s.client.NewRequest(http.MethodPost, path, requestBody)
	if err != nil {
		return nil, err
	}

	uploadCredentials := new(models.UploadCredentials)
	_, err = s.client.Do(req, uploadCredentials)
	return uploadCredentials, err
}

// CompleteUpload completes the upload process by notifying the API that the file has been uploaded
func (s *AttachmentsService) CompleteUpload(workspaceSlug string, projectID string, issueID string, assetID string) (*models.Attachment, error) {
	path := fmt.Sprintf("/workspaces/%s/projects/%s/issues/%s/issue-attachments/%s/", workspaceSlug, projectID, issueID, assetID)
	req, err := s.client.NewRequest(http.MethodPatch, path, nil)
	if err != nil {
		return nil, err
	}

	attachment := new(models.Attachment)
	_, err = s.client.Do(req, attachment)
	return attachment, err
}

// UploadFile uploads a file to S3 using the provided credentials
func (s *AttachmentsService) UploadFile(uploadURL string, fields map[string]string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields from S3 policy
	for key, val := range fields {
		if err := writer.WriteField(key, val); err != nil {
			return err
		}
	}

	// Add file
	fileWriter, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return err
	}

	if _, err = io.Copy(fileWriter, file); err != nil {
		return err
	}

	if err = writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, uploadURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("S3 upload failed: %s (Status: %d)\nBody: %s",
			uploadURL, resp.StatusCode, string(bodyBytes))
	}

	return nil
}

// UploadFileToIssue uploads a file to an issue in a single step
// 一步完成文件上传到问题，包括获取上传凭证、上传文件和完成上传过程
func (s *AttachmentsService) UploadFileToIssue(workspaceSlug string, projectID string, issueID string, filePath string) (*models.Attachment, error) {
	// 获取文件信息
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("无法获取文件信息: %w", err)
	}

	filename := filepath.Base(filePath)
	fileSize := fileInfo.Size()

	// 检测内容类型
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("无法读取文件内容: %w", err)
	}

	// 重置文件指针到开始位置
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("无法重置文件指针: %w", err)
	}

	contentType := http.DetectContentType(buffer)

	// 获取上传凭证
	credentials, err := s.GetUploadCredentials(workspaceSlug, projectID, issueID, filename, contentType, fileSize)
	if err != nil {
		return nil, fmt.Errorf("获取上传凭证失败: %w", err)
	}

	// 上传文件
	err = s.UploadFile(credentials.UploadData.URL, credentials.UploadData.Fields, filePath)
	if err != nil {
		return nil, fmt.Errorf("上传文件失败: %w", err)
	}

	// 完成上传过程
	attachment, err := s.CompleteUpload(workspaceSlug, projectID, issueID, credentials.AssetID)
	if err != nil {
		return nil, fmt.Errorf("完成上传过程失败: %w", err)
	}

	return attachment, nil
}
