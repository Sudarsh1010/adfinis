package workspace

import "time"

// CreateWorkspaceRequest defines the input for creating a new workspace.
type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

// Response defines the output when returning workspace data.
type Response struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsDefault bool      `json:"isDefault"`
	CreatedAt time.Time `json:"createdAt"`
}
