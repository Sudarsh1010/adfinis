package workspace

import (
	"context"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
)

// Service defines the application layer contract for workspace operations.
// This service orchestrates business logic and coordinates with the repository layer.
// It enforces validation rules and transforms domain entities into DTOs.
type Service interface {
	// CreateWorkspace creates a new workspace with validation.
	// Returns ErrInvalidName if the name is empty.
	CreateWorkspace(
		ctx context.Context,
		req *CreateWorkspaceRequest,
	) (*Response, error)

	// GetWorkspaceByID retrieves a workspace by its unique identifier.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	GetWorkspaceByID(
		ctx context.Context,
		id workspace.ID,
	) (*Response, error)

	// ListWorkspaces retrieves all workspaces.
	// Returns an empty list if no workspaces exist.
	ListWorkspaces(ctx context.Context) ([]*Response, error)

	// UpdateWorkspaceName changes the name of an existing workspace.
	// Returns ErrInvalidName if the new name is empty.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	UpdateWorkspaceName(
		ctx context.Context,
		id workspace.ID,
		newName string,
	) (*Response, error)

	// DeleteWorkspace removes a workspace from the data store.
	// Returns ErrWorkspaceNotFound if the workspace does not exist.
	DeleteWorkspace(ctx context.Context, id workspace.ID) error
}

// workspaceService is the concrete implementation of WorkspaceService.
// It uses WorkspaceRepository for data persistence and transforms between
// domain entities and DTOs.
type workspaceService struct {
	repo Repository
}

// NewWorkspaceService creates a new WorkspaceService instance.
func NewWorkspaceService(repo Repository) Service {
	return &workspaceService{
		repo: repo,
	}
}

// CreateWorkspace creates a new workspace with validation.
func (s *workspaceService) CreateWorkspace(
	ctx context.Context,
	req *CreateWorkspaceRequest,
) (*Response, error) {
	if req.Name == "" {
		return nil, workspace.ErrInvalidName
	}

	entity := workspace.NewWorkspace(req.Name)
	entity.SetIsDefault(false)

	err := s.repo.Save(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:        string(entity.ID()),
		Name:      entity.Name(),
		IsDefault: entity.IsDefault(),
		CreatedAt: entity.CreatedAt(),
	}, nil
}

// GetWorkspaceByID retrieves a workspace by its unique identifier.
func (s *workspaceService) GetWorkspaceByID(
	ctx context.Context,
	id workspace.ID,
) (*Response, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, workspace.ErrWorkspaceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:        string(entity.ID()),
		Name:      entity.Name(),
		IsDefault: entity.IsDefault(),
		CreatedAt: entity.CreatedAt(),
	}, nil
}

// ListWorkspaces retrieves all workspaces.
func (s *workspaceService) ListWorkspaces(
	ctx context.Context,
) ([]*Response, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*Response, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, &Response{
			ID:        string(entity.ID()),
			Name:      entity.Name(),
			IsDefault: entity.IsDefault(),
			CreatedAt: entity.CreatedAt(),
		})
	}

	return responses, nil
}

// UpdateWorkspaceName changes the name of an existing workspace.
func (s *workspaceService) UpdateWorkspaceName(
	ctx context.Context,
	id workspace.ID,
	newName string,
) (*Response, error) {
	if newName == "" {
		return nil, workspace.ErrInvalidName
	}

	err := s.repo.UpdateName(ctx, id, newName)
	if err != nil {
		return nil, err
	}

	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, workspace.ErrWorkspaceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:        string(entity.ID()),
		Name:      entity.Name(),
		IsDefault: entity.IsDefault(),
		CreatedAt: entity.CreatedAt(),
	}, nil
}

// DeleteWorkspace removes a workspace from the data store.
func (s *workspaceService) DeleteWorkspace(ctx context.Context, id workspace.ID) error {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return workspace.ErrWorkspaceNotFound
	}
	if err != nil {
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
