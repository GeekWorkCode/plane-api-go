package plane

import (
	"github.com/GeekWorkCode/plane-api-go/api"
	"github.com/GeekWorkCode/plane-api-go/client"
	"github.com/GeekWorkCode/plane-api-go/models"
)

// Comments handles comment-related operations.
// It provides methods for listing, creating, updating, and deleting
// comments, with support for using either member IDs or display names.
type CommentsService interface {
	List(workspaceSlug string, projectID string, issueID string) ([]models.Comment, error)
	Get(workspaceSlug string, projectID string, issueID string, commentID string) (*models.Comment, error)
	Create(workspaceSlug string, projectID string, issueID string, request *api.CommentRequest) (*models.Comment, error)
	Update(workspaceSlug string, projectID string, issueID string, commentID string, request *api.CommentRequest) (*models.Comment, error)
	Delete(workspaceSlug string, projectID string, issueID string, commentID string) error
}

// Plane is the main API client
type Plane struct {
	// Client is the HTTP client used to communicate with the API
	client *client.Client

	// Services for the different parts of the Plane API
	Projects    *api.ProjectsService
	Issues      *api.IssuesService
	Cycles      *api.CyclesService
	Modules     *api.ModulesService
	Labels      *api.LabelsService
	States      *api.StatesService
	Comments    *api.CommentsService
	Links       *api.LinksService
	Attachments *api.AttachmentsService
	Worklogs    *api.WorklogsService
	Members     *api.MembersService
}

// NewClient returns a new Plane API client
func NewClient(apiKey string) *Plane {
	c := client.NewClient(apiKey)

	return &Plane{
		client:      c,
		Projects:    api.NewProjectsService(c),
		Issues:      api.NewIssuesService(c),
		Cycles:      api.NewCyclesService(c),
		Modules:     api.NewModulesService(c),
		Labels:      api.NewLabelsService(c),
		States:      api.NewStatesService(c),
		Comments:    api.NewCommentsService(c),
		Links:       api.NewLinksService(c),
		Attachments: api.NewAttachmentsService(c),
		Worklogs:    api.NewWorklogsService(c),
		Members:     api.NewMembersService(c),
	}
}

// SetDebug enables or disables debug logging
func (p *Plane) SetDebug(debug bool) {
	p.client.SetDebug(debug)
}

// SetBaseURL sets the base URL for API requests
func (p *Plane) SetBaseURL(baseURL string) {
	p.client.SetBaseURL(baseURL)
}
