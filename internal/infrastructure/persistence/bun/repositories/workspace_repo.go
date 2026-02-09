package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"

	"github.com/uptrace/bun"
)

type WorkspaceRepository struct {
	db *bun.DB
}

func NewWorkspaceRepository(db *bun.DB) *WorkspaceRepository {
	return &WorkspaceRepository{db: db}
}

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

func (r *WorkspaceRepository) FindByID(ctx context.Context, id workspace.WorkspaceID) (*workspace.Workspace, error) {
	var model models.Workspace
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("workspace not found: %w", err)
	}

	ws := &workspace.Workspace{}
	ws.SetID(workspace.WorkspaceID(model.ID))
	ws.SetName(model.Name)
	ws.SetIsDefault(model.IsDefault)
	ws.SetCreatedAt(model.CreatedAt)
	return ws, nil
}

func (r *WorkspaceRepository) FindAll(ctx context.Context) ([]*workspace.Workspace, error) {
	var models []models.Workspace
	err := r.db.NewSelect().Model(&models).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all workspaces: %w", err)
	}

	workspaces := make([]*workspace.Workspace, len(models))
	for i, model := range models {
		ws := &workspace.Workspace{}
		ws.SetID(workspace.WorkspaceID(model.ID))
		ws.SetName(model.Name)
		ws.SetIsDefault(model.IsDefault)
		ws.SetCreatedAt(model.CreatedAt)
		workspaces[i] = ws
	}
	return workspaces, nil
}

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
