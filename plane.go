package plane

import (
	"github.com/GeekWorkCode/plane-api-go/api"
	"github.com/GeekWorkCode/plane-api-go/client"
)

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
