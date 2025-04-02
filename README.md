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

// Create an issue using state name and assignee name
newIssue, err := client.Issues.Create("your-workspace-slug", "project-id", &api.IssueCreateRequest{
    Name:          "Issue with State Name and Assignee Name",
    Description:   "An issue created using state name and assignee name",
    StateName:     "Backlog",       // Will be converted to state ID automatically
    Priority:      "medium",
    AssigneeNames: []string{"John Doe"}, // Will be converted to member ID automatically
})

// Create an issue with multiple assignees
newIssue, err := client.Issues.Create("your-workspace-slug", "project-id", &api.IssueCreateRequest{
    Name:          "Issue with Multiple Assignees",
    Description:   "An issue created with multiple assignees",
    StateName:     "Backlog",
    Priority:      "high",
    AssigneeNames: []string{"John Doe", "Jane Smith"}, // Will be converted to member IDs automatically
})

// Update an issue
updatedIssue, err := client.Issues.Update("your-workspace-slug", "project-id", "issue-id", &api.IssueUpdateRequest{
    Name:        "Updated Issue Name",
    Description: "Updated issue description",
})

// Delete an issue
err := client.Issues.Delete("your-workspace-slug", "project-id", "issue-id")
```

### States

```go
// List all states in a project
states, err := client.States.List("your-workspace-slug", "project-id")

// Get a state by ID
state, err := client.States.Get("your-workspace-slug", "project-id", "state-id")

// Create a new state
newState, err := client.States.Create("your-workspace-slug", "project-id", &api.StateCreateRequest{
    Name:        "New State",
    Description: "A new state created via the API",
    Color:       "#4CAF50",  // Green color
})

// Update a state
updatedState, err := client.States.Update("your-workspace-slug", "project-id", "state-id", &api.StateUpdateRequest{
    Name:        "Updated State Name",
    Description: "Updated state description",
    Color:       "#FFC107",  // Yellow color
})

// Delete a state
err := client.States.Delete("your-workspace-slug", "project-id", "state-id")

// Create an issue with a specific state
newIssue, err := client.Issues.Create("your-workspace-slug", "project-id", &api.IssueCreateRequest{
    Name:        "New Issue with Custom State",
    Description: "An issue using a custom state",
    State:       newState.ID,
})
```

### Comments

```go
// List all comments for an issue
comments, err := client.Comments.List("your-workspace-slug", "project-id", "issue-id")

// Get a comment by ID
comment, err := client.Comments.Get("your-workspace-slug", "project-id", "issue-id", "comment-id")

// Create a comment using DisplayName (library will look up the member ID)
comment, err := client.Comments.Create("your-workspace-slug", "project-id", "issue-id", &api.CommentRequest{
    CommentHTML: "<p>This is a comment created via DisplayName</p>",
    DisplayName: "John Doe",
})

// Create a comment using MemberID directly
comment, err := client.Comments.Create("your-workspace-slug", "project-id", "issue-id", &api.CommentRequest{
    CommentHTML: "<p>This is a comment created via MemberID</p>",
    CreatedBy: "member-id",
})

// Update a comment using DisplayName
updatedComment, err := client.Comments.Update("your-workspace-slug", "project-id", "issue-id", "comment-id", &api.CommentRequest{
    CommentHTML: "<p>This is an updated comment via DisplayName</p>",
    DisplayName: "John Doe",
})

// Update a comment using MemberID (through Actor field)
updatedComment, err := client.Comments.Update("your-workspace-slug", "project-id", "issue-id", "comment-id", &api.CommentRequest{
    CommentHTML: "<p>This is an updated comment via Actor</p>",
    Actor: "member-id",
})

// Delete a comment
err := client.Comments.Delete("your-workspace-slug", "project-id", "issue-id", "comment-id")
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

### Members

```go
// List all members in a project
members, err := client.Members.List("your-workspace-slug", "project-id")

// Get a member by ID
member, err := client.Members.Get("your-workspace-slug", "project-id", "member-id")
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

## Comment Author Handling

The Plane API has a specific behavior regarding comment creation and author display. The library automatically handles this behavior to ensure the comment author is displayed correctly:

1. When creating a comment, you can use either `DisplayName` or `CreatedBy` (Member ID).
2. If you use `DisplayName`, the library will automatically look up the corresponding Member ID.
3. The library includes logic to handle potential discrepancies between the provided creator ID and the displayed creator in the API response.
4. If needed, the library will automatically perform a follow-up update operation to ensure the correct author is displayed.

This approach ensures a consistent experience when creating comments through the API, matching what users would see in the Plane UI.

## States and Issue Workflow

The States API allows you to manage the workflow states for issues in your projects:

1. States represent different stages in your issue workflow (e.g., "Todo", "In Progress", "Done").
2. Each state has a name, description, and color for visual identification in the UI.
3. States can be created, updated, and deleted through the API.
4. When creating or updating issues, you can specify which state to use by providing the state ID.

Note that states that are currently being used by issues may not be deletable until those issues are moved to different states.

## License

MIT 