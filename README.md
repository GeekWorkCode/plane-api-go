# Plane API Go Client

A Go client library for the [Plane](https://plane.so) API.

## Installation

```bash
go get github.com/GeekWorkCode/plane-api-go
```

## Authentication

To use the Plane API, you'll need an API key. You can generate one from your workspace settings:

1. Log into your Plane account and go to `Workspace Settings`.
2. Go to `API tokens` in the list of tabs available.
3. Click `Add API token`.
4. Choose a title and description so you know why you are creating this token and where you will use is.
5. Choose an expiry if you want this to stop working after a point.

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/GeekWorkCode/plane-api-go"
)

func main() {
    // Create a new client
    client := plane.NewClient("your-api-key")

    // Set debug mode to see API requests and responses (optional)
    client.SetDebug(true)

    // Get all projects in a workspace
    projects, err := client.Projects.List("your-workspace-slug")
    if err != nil {
        log.Fatalf("Error getting projects: %v", err)
    }

    for _, project := range projects {
        fmt.Printf("Project: %s (ID: %s)\n", project.Name, project.ID)
    }
}
```

### Projects

```go
// List all projects
projects, err := client.Projects.List("your-workspace-slug")

// Get a project by ID
project, err := client.Projects.Get("your-workspace-slug", "project-id")

// Create a new project
newProject, err := client.Projects.Create("your-workspace-slug", &api.ProjectCreateRequest{
    Name:        "New Project",
    Identifier:  "NEW",
    Description: "A new project created via the API",
})

// Update a project
updatedProject, err := client.Projects.Update("your-workspace-slug", "project-id", &api.ProjectUpdateRequest{
    Name:        "Updated Project Name",
    Description: "Updated project description",
})

// Delete a project
err := client.Projects.Delete("your-workspace-slug", "project-id")
```

### Issues

```go
// List all issues in a project
issues, err := client.Issues.List("your-workspace-slug", "project-id")

// Get an issue by ID
issue, err := client.Issues.Get("your-workspace-slug", "project-id", "issue-id")

// Get an issue by sequence ID
issue, err := client.Issues.GetBySequenceID("your-workspace-slug", "sequence-id")

// Create a new issue
newIssue, err := client.Issues.Create("your-workspace-slug", "project-id", &api.IssueCreateRequest{
    Name:        "New Issue",
    Description: "A new issue created via the API",
})

// Update an issue
updatedIssue, err := client.Issues.Update("your-workspace-slug", "project-id", "issue-id", &api.IssueUpdateRequest{
    Name:        "Updated Issue Name",
    Description: "Updated issue description",
})

// Delete an issue
err := client.Issues.Delete("your-workspace-slug", "project-id", "issue-id")
```

### Cycles

```go
// List all cycles in a project
cycles, err := client.Cycles.List("your-workspace-slug", "project-id")

// Get a cycle by ID
cycle, err := client.Cycles.Get("your-workspace-slug", "project-id", "cycle-id")

// Create a new cycle
newCycle, err := client.Cycles.Create("your-workspace-slug", "project-id", &api.CycleCreateRequest{
    Name:        "New Cycle",
    Description: "A new cycle created via the API",
    StartDate:   "2023-01-01",
    EndDate:     "2023-01-15",
})

// Update a cycle
updatedCycle, err := client.Cycles.Update("your-workspace-slug", "project-id", "cycle-id", &api.CycleUpdateRequest{
    Name:        "Updated Cycle Name",
    Description: "Updated cycle description",
})

// Delete a cycle
err := client.Cycles.Delete("your-workspace-slug", "project-id", "cycle-id")

// List all issues in a cycle
issues, err := client.Cycles.ListIssues("your-workspace-slug", "project-id", "cycle-id")

// Add issues to a cycle
err := client.Cycles.AddIssues("your-workspace-slug", "project-id", "cycle-id", []string{"issue-id-1", "issue-id-2"})

// Remove an issue from a cycle
err := client.Cycles.RemoveIssue("your-workspace-slug", "project-id", "cycle-id", "issue-id")
```

### Attachments

```go
// List all attachments for an issue
attachments, err := client.Attachments.List("your-workspace-slug", "project-id", "issue-id")

// Get upload credentials for an attachment
credentials, err := client.Attachments.GetUploadCredentials(
    "your-workspace-slug", 
    "project-id", 
    "issue-id", 
    "filename.pdf", 
    "application/pdf", 
    12345 // filesize in bytes
)

// Mark an attachment as uploaded (after uploading to S3)
attachment, err := client.Attachments.CompleteUpload(
    "your-workspace-slug", 
    "project-id", 
    "issue-id", 
    credentials.AssetID
)

// Upload a file to S3 using the provided credentials
err := client.Attachments.UploadFile(
    credentials.UploadData.URL,
    credentials.UploadData.Fields,
    "path/to/local/file.pdf"
)

// Convenience method that handles the entire upload process
attachment, err := client.Attachments.UploadFileToIssue(
    "your-workspace-slug", 
    "project-id", 
    "issue-id", 
    "path/to/local/file.pdf"
)
```

## Running the Examples

1. Navigate to the examples directory
2. Copy the `.env.example` file to `.env` and fill in your API key and workspace slug:
   ```bash
   cp .env.example .env
   ```
3. Edit the `.env` file to add your API key and workspace slug
4. Run the example:
   ```bash
   go run main.go
   ```

## Debug Mode

You can enable debug mode to see the API requests and responses:

```go
client.SetDebug(true)
```

This is useful for debugging and understanding the API behavior.

## Custom Base URL

If you're using a self-hosted Plane instance, you can set a custom base URL:

```go
client.SetBaseURL("https://your-plane-instance.com/api/v1")
```

## License

MIT 