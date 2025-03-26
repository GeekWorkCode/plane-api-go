package models

import (
	"time"
)

// Common response structures
// PagedResponse 分页响应的通用结构
type PagedResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         interface{} `json:"results"`
}

// ProjectsResponse 项目列表的分页响应
type ProjectsResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         []Project   `json:"results"`
}

// IssuesResponse 问题列表的分页响应
type IssuesResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         []Issue     `json:"results"`
}

// CyclesResponse 周期列表的分页响应
type CyclesResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         []Cycle     `json:"results"`
}

// ModulesResponse 模块列表的分页响应
type ModulesResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         []Module    `json:"results"`
}

// AttachmentsResponse 附件列表的分页响应
type AttachmentsResponse struct {
	GroupedBy       interface{}  `json:"grouped_by"`
	SubGroupedBy    interface{}  `json:"sub_grouped_by"`
	TotalCount      int          `json:"total_count"`
	NextCursor      string       `json:"next_cursor"`
	PrevCursor      string       `json:"prev_cursor"`
	NextPageResults bool         `json:"next_page_results"`
	PrevPageResults bool         `json:"prev_page_results"`
	Count           int          `json:"count"`
	TotalPages      int          `json:"total_pages"`
	TotalResults    int          `json:"total_results"`
	ExtraStats      interface{}  `json:"extra_stats"`
	Results         []Attachment `json:"results"`
}

// CommentsResponse 评论列表的分页响应
type CommentsResponse struct {
	GroupedBy       interface{} `json:"grouped_by"`
	SubGroupedBy    interface{} `json:"sub_grouped_by"`
	TotalCount      int         `json:"total_count"`
	NextCursor      string      `json:"next_cursor"`
	PrevCursor      string      `json:"prev_cursor"`
	NextPageResults bool        `json:"next_page_results"`
	PrevPageResults bool        `json:"prev_page_results"`
	Count           int         `json:"count"`
	TotalPages      int         `json:"total_pages"`
	TotalResults    int         `json:"total_results"`
	ExtraStats      interface{} `json:"extra_stats"`
	Results         []Comment   `json:"results"`
}

type Pagination struct {
	Count        int  `json:"count"`
	TotalPages   int  `json:"total_pages"`
	NextPage     *int `json:"next_page"`
	PreviousPage *int `json:"previous_page"`
	Page         int  `json:"page"`
}

// Error response from the API
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Workspace represents a Plane workspace
type Workspace struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Project represents a Plane project
type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Identifier  string    `json:"identifier"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Workspace   string    `json:"workspace"`
}

// Issue represents a Plane issue
type Issue struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	State       string    `json:"state,omitempty"`
	Priority    string    `json:"priority,omitempty"`
	AssigneeID  string    `json:"assignee_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
}

// Cycle represents a Plane cycle
type Cycle struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	StartDate   string    `json:"start_date,omitempty"`
	EndDate     string    `json:"end_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
}

// Module represents a Plane module
type Module struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
}

// Label represents a Plane label
type Label struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Color       string    `json:"color,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
	Parent      *string   `json:"parent"`
}

// Comment represents a comment on an issue
type Comment struct {
	ID          string    `json:"id"`
	CommentHTML string    `json:"comment_html"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
	Issue       string    `json:"issue"`
}

// State represents a state in the project (e.g., Todo, In Progress, Done)
type State struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	Project     string    `json:"project"`
	Workspace   string    `json:"workspace"`
}

// Link represents a link attached to an issue
type Link struct {
	ID        string                 `json:"id"`
	Title     string                 `json:"title,omitempty"`
	URL       string                 `json:"url"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	CreatedBy string                 `json:"created_by"`
	UpdatedBy string                 `json:"updated_by"`
	Project   string                 `json:"project"`
	Workspace string                 `json:"workspace"`
	Issue     string                 `json:"issue"`
}

// Worklog represents time spent on a specific issue
type Worklog struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"` // Duration in minutes
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	ProjectID   string    `json:"project_id"`
	WorkspaceID string    `json:"workspace_id"`
	LoggedBy    string    `json:"logged_by"`
}

// WorklogTotal represents the aggregated time for an issue
type WorklogTotal struct {
	IssueID  string  `json:"issue_id"`
	Duration float64 `json:"duration"` // Duration in minutes
}

// Attachment represents a file attached to an issue
type Attachment struct {
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	DeletedAt       *time.Time             `json:"deleted_at"`
	Attributes      map[string]interface{} `json:"attributes"`
	Asset           string                 `json:"asset"`
	EntityType      string                 `json:"entity_type"`
	IsDeleted       bool                   `json:"is_deleted"`
	IsArchived      bool                   `json:"is_archived"`
	ExternalID      *string                `json:"external_id"`
	ExternalSource  *string                `json:"external_source"`
	Size            float64                `json:"size"`
	IsUploaded      bool                   `json:"is_uploaded"`
	StorageMetadata map[string]interface{} `json:"storage_metadata"`
	CreatedBy       string                 `json:"created_by"`
	UpdatedBy       *string                `json:"updated_by"`
	Workspace       string                 `json:"workspace"`
	Project         string                 `json:"project"`
	Issue           string                 `json:"issue"`
}

// S3UploadData contains the pre-signed URL and fields for direct S3 upload
type S3UploadData struct {
	URL    string            `json:"url"`
	Fields map[string]string `json:"fields"`
}

// UploadCredentials represents the response from the get upload credentials endpoint
type UploadCredentials struct {
	UploadData S3UploadData `json:"upload_data"`
	AssetID    string       `json:"asset_id"`
	Attachment Attachment   `json:"attachment"`
	AssetURL   string       `json:"asset_url"`
}
