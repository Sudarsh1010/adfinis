package workspace

import (
	"context"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
)

// WorkspaceService defines the application layer contract for workspace operations.
// This service orchestrates business logic and coordinates with the repository layer.
// It enforces validation rules and transforms domain entities into DTOs.
type WorkspaceService interface {
	// CreateWorkspace creates a new workspace with validation.
	// Returns ErrInvalidName if the name is empty.
	CreateWorkspace(ctx context.Context, req *CreateWorkspaceRequest) (*WorkspaceResponse, error)

	// GetWorkspaceByID retrieves a workspace by its unique identifier.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	GetWorkspaceByID(ctx context.Context, id workspace.WorkspaceID) (*WorkspaceResponse, error)

	// ListWorkspaces retrieves all workspaces.
	// Returns an empty list if no workspaces exist.
	ListWorkspaces(ctx context.Context) ([]*WorkspaceResponse, error)

	// UpdateWorkspaceName changes the name of an existing workspace.
	// Returns ErrInvalidName if the new name is empty.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	UpdateWorkspaceName(ctx context.Context, id workspace.WorkspaceID, newName string) (*WorkspaceResponse, error)

	// DeleteWorkspace removes a workspace from the data store.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	DeleteWorkspace(ctx context.Context, id workspace.WorkspaceID) error
}

// workspaceService is the concrete implementation of WorkspaceService.
// It uses WorkspaceRepository for data persistence and transforms between
// domain entities and DTOs.
type workspaceService struct {
	repo WorkspaceRepository
}

// NewWorkspaceService creates a new WorkspaceService instance.
func NewWorkspaceService(repo WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		repo: repo,
	}
}
