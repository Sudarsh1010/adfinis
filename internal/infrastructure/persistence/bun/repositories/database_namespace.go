package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/database_namespace"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"
	"github.com/uptrace/bun"
)

// DatabaseNamespaceRepository implements persistence for DatabaseNamespace domain entity.
type DatabaseNamespaceRepository struct {
	db *bun.DB
}

// NewDatabaseNamespaceRepository creates a new DatabaseNamespaceRepository with the given database connection.
func NewDatabaseNamespaceRepository(db *bun.DB) *DatabaseNamespaceRepository {
	return &DatabaseNamespaceRepository{db: db}
}

// toDomainEntity converts a models.DatabaseNamespace to a domain.DatabaseNamespace entity.
func toDatabaseNamespaceDomainEntity(m *models.DatabaseNamespace) *database_namespace.DatabaseNamespace {
	d := &database_namespace.DatabaseNamespace{}
	d.SetID(database_namespace.ID(m.ID))
	d.SetConnectionID(m.ConnectionID)
	d.SetName(m.Name)
	d.SetCreatedAt(m.CreatedAt)
	d.SetUpdatedAt(m.UpdatedAt)
	return d
}

// toDatabaseNamespaceModel converts a domain.DatabaseNamespace entity to a database model.
func toDatabaseNamespaceModel(d *database_namespace.DatabaseNamespace) *models.DatabaseNamespace {
	return &models.DatabaseNamespace{
		ID:           string(d.ID()),
		ConnectionID: d.ConnectionID(),
		Name:         d.Name(),
		CreatedAt:    d.CreatedAt(),
		UpdatedAt:    d.UpdatedAt(),
	}
}

// Save persists a database namespace entity to the database.
func (r *DatabaseNamespaceRepository) Save(ctx context.Context, d *database_namespace.DatabaseNamespace) error {
	model := toDatabaseNamespaceModel(d)
	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves a database namespace by its ID from the database.
func (r *DatabaseNamespaceRepository) FindByID(ctx context.Context, id database_namespace.ID) (*database_namespace.DatabaseNamespace, error) {
	var model models.DatabaseNamespace
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("database namespace not found: %w", err)
	}
	return toDatabaseNamespaceDomainEntity(&model), nil
}

// FindAll retrieves all database namespaces from the database.
func (r *DatabaseNamespaceRepository) FindAll(ctx context.Context) ([]*database_namespace.DatabaseNamespace, error) {
	var modelList []models.DatabaseNamespace
	err := r.db.NewSelect().Model(&modelList).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all database namespaces: %w", err)
	}

	namespaces := make([]*database_namespace.DatabaseNamespace, len(modelList))
	for i := range modelList {
		namespaces[i] = toDatabaseNamespaceDomainEntity(&modelList[i])
	}
	return namespaces, nil
}

// Update updates a database namespace in the database.
func (r *DatabaseNamespaceRepository) Update(ctx context.Context, d *database_namespace.DatabaseNamespace) error {
	model := toDatabaseNamespaceModel(d)
	result, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", string(d.ID())).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update database namespace: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return database_namespace.ErrDatabaseNamespaceNotFound
	}

	return nil
}

// Delete removes a database namespace from the database by its ID.
func (r *DatabaseNamespaceRepository) Delete(ctx context.Context, id database_namespace.ID) error {
	result, err := r.db.NewDelete().
		Model(&models.DatabaseNamespace{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete database namespace: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return database_namespace.ErrDatabaseNamespaceNotFound
	}

	return nil
}
