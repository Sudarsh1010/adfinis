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

	// Map Bun model → domain entity
	return &workspace.Workspace{
		id:        workspace.WorkspaceID(model.ID),
		name:      model.Name,
		isDefault: model.IsDefault,
		createdAt: model.CreatedAt,
	}, nil
}
