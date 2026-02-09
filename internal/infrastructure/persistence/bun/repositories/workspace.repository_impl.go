package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"

	"github.com/uptrace/bun"
)

// WorkspaceRepository implements persistence for Workspace domain entity.
// This repository maps between domain.Workspace entities and models.Workspace database models
// using the Bun ORM.
type WorkspaceRepository struct {
	db *bun.DB
}

// NewWorkspaceRepository creates a new WorkspaceRepository with the given database connection.
func NewWorkspaceRepository(db *bun.DB) *WorkspaceRepository {
	return &WorkspaceRepository{db: db}
}

// toDomainEntity converts a models.Workspace to a domain.Workspace entity.
func toDomainEntity(m *models.Workspace) *workspace.Workspace {
	ws := &workspace.Workspace{}
	ws.SetID(workspace.WorkspaceID(m.ID))
	ws.SetName(m.Name)
	ws.SetIsDefault(m.IsDefault)
	ws.SetCreatedAt(m.CreatedAt)
	return ws
}

// Save persists a workspace entity to the database.
// Converts the domain entity to a database model and inserts it.
func (r *WorkspaceRepository) Save(ctx context.Context, ws *workspace.Workspace) error {
	model := &models.Workspace{
		ID:        string(ws.ID()),
		Name:      ws.Name(),
		IsDefault: ws.IsDefault(),
		CreatedAt: ws.CreatedAt(),
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves a workspace by its ID from the database.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (r *WorkspaceRepository) FindByID(ctx context.Context, id workspace.WorkspaceID) (*workspace.Workspace, error) {
	var model models.Workspace
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("workspace not found: %w", err)
	}

	return toDomainEntity(&model), nil
}

// FindAll retrieves all workspaces from the database.
func (r *WorkspaceRepository) FindAll(ctx context.Context) ([]*workspace.Workspace, error) {
	var models []models.Workspace
	err := r.db.NewSelect().Model(&models).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all workspaces: %w", err)
	}

	workspaces := make([]*workspace.Workspace, len(models))
	for i := range models {
		workspaces[i] = toDomainEntity(&models[i])
	}
	return workspaces, nil
}

// UpdateName updates the name of a workspace in the database.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (r *WorkspaceRepository) UpdateName(ctx context.Context, id workspace.WorkspaceID, name string) error {
	result, err := r.db.NewUpdate().
		Model(&models.Workspace{}).
		Set("name = ?", name).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update workspace name: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return workspace.ErrWorkspaceNotFound
	}

	return nil
}

// Delete removes a workspace from the database by its ID.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (r *WorkspaceRepository) Delete(ctx context.Context, id workspace.WorkspaceID) error {
	result, err := r.db.NewDelete().
		Model(&models.Workspace{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete workspace: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return workspace.ErrWorkspaceNotFound
	}

	return nil
}
