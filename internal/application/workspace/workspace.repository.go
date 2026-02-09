package workspace

import (
	"context"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
)

// WorkspaceRepository defines the contract for workspace data access.
// This interface isolates the application layer from infrastructure concerns.
// The implementation (in infrastructure/persistence/bun/repositories) provides
// actual database operations using Bun ORM.
type WorkspaceRepository interface {
	// Save persists a new workspace to the data store.
	Save(ctx context.Context, ws *workspace.Workspace) error

	// FindByID retrieves a workspace by its unique identifier.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	FindByID(ctx context.Context, id workspace.WorkspaceID) (*workspace.Workspace, error)

	// FindAll retrieves all workspaces from the data store.
	FindAll(ctx context.Context) ([]*workspace.Workspace, error)

	// UpdateName changes the name of an existing workspace.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	UpdateName(ctx context.Context, id workspace.WorkspaceID, name string) error

	// Delete removes a workspace from the data store by its identifier.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	Delete(ctx context.Context, id workspace.WorkspaceID) error
}
