package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/sudarsh1010/adfinis/internal/domain/collection"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"
	"github.com/uptrace/bun"
)
// CollectionRepository implements persistence for Collection domain entity.
type CollectionRepository struct {
	db *bun.DB
}

// NewCollectionRepository creates a new CollectionRepository with the given database connection.
func NewCollectionRepository(db *bun.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

// toDomainEntity converts a models.Collection to a domain.Collection entity.
func toCollectionDomainEntity(m *models.Collection) *collection.Collection {
	c := &collection.Collection{}
	c.SetID(collection.ID(m.ID))
	c.SetDatabaseNamespaceID(m.DatabaseNamespaceID)
	c.SetName(m.Name)
	c.SetEmbeddingProfileID(m.EmbeddingProfileID)
	c.SetVectorCount(int(m.VectorCount))
	c.SetMetadataSchemaJSON(m.MetadataSchemaJSON)
	c.SetCreatedAt(m.CreatedAt)
	c.SetUpdatedAt(m.UpdatedAt)
	if m.LastSyncedAt != nil {
		c.SetLastSyncedAt(*m.LastSyncedAt)
	}
	return c
}

// toCollectionModel converts a domain.Collection entity to a database model.
func toCollectionModel(c *collection.Collection) *models.Collection {
	m := &models.Collection{
		ID:                  string(c.ID()),
		DatabaseNamespaceID: c.DatabaseNamespaceID(),
		Name:                c.Name(),
		EmbeddingProfileID:  c.EmbeddingProfileID(),
		VectorCount:         int64(c.VectorCount()),
		MetadataSchemaJSON:  c.MetadataSchemaJSON(),
		CreatedAt:           c.CreatedAt(),
		UpdatedAt:           c.UpdatedAt(),
	}
	if !c.LastSyncedAt().IsZero() {
		t := c.LastSyncedAt()
		m.LastSyncedAt = &t
	}
	return m
}

// Save persists a collection entity to the database.
func (r *CollectionRepository) Save(ctx context.Context, c *collection.Collection) error {
	model := toCollectionModel(c)
	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves a collection by its ID from the database.
func (r *CollectionRepository) FindByID(ctx context.Context, id collection.ID) (*collection.Collection, error) {
	var model models.Collection
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("collection not found: %w", err)
	}
	return toCollectionDomainEntity(&model), nil
}

// FindAll retrieves all collections from the database.
func (r *CollectionRepository) FindAll(ctx context.Context) ([]*collection.Collection, error) {
	var modelList []models.Collection
	err := r.db.NewSelect().Model(&modelList).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all collections: %w", err)
	}

	collections := make([]*collection.Collection, len(modelList))
	for i := range modelList {
		collections[i] = toCollectionDomainEntity(&modelList[i])
	}
	return collections, nil
}

// FindByDatabaseNamespaceID retrieves all collections for a specific database namespace from the database.
func (r *CollectionRepository) FindByDatabaseNamespaceID(ctx context.Context, databaseNamespaceID string) ([]*collection.Collection, error) {
	var modelList []models.Collection
	err := r.db.NewSelect().Model(&modelList).Where("database_namespace_id = ?", databaseNamespaceID).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find collections by namespace: %w", err)
	}

	collections := make([]*collection.Collection, len(modelList))
	for i := range modelList {
		collections[i] = toCollectionDomainEntity(&modelList[i])
	}
	return collections, nil
}

// UpdateName updates the name of a collection in the database.
// Returns ErrCollectionNotFound if the collection does not exist.
func (r *CollectionRepository) UpdateName(ctx context.Context, id collection.ID, name string) error {
	result, err := r.db.NewUpdate().
		Model(&models.Collection{}).
		Set("name = ?", name).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update collection name: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return collection.ErrCollectionNotFound
	}

	return nil
}

// UpdateVectorCount updates the vector count of a collection in the database.
// Returns ErrCollectionNotFound if the collection does not exist.
func (r *CollectionRepository) UpdateVectorCount(ctx context.Context, id collection.ID, vectorCount int) error {
	result, err := r.db.NewUpdate().
		Model(&models.Collection{}).
		Set("vector_count = ?", vectorCount).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update collection vector count: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return collection.ErrCollectionNotFound
	}

	return nil
}

// UpdateLastSyncedAt updates the last synced at timestamp of a collection in the database.
// Returns ErrCollectionNotFound if the collection does not exist.
func (r *CollectionRepository) UpdateLastSyncedAt(ctx context.Context, id collection.ID, lastSyncedAt time.Time) error {
	result, err := r.db.NewUpdate().
		Model(&models.Collection{}).
		Set("last_synced_at = ?", lastSyncedAt).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update collection last synced at: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return collection.ErrCollectionNotFound
	}

	return nil
}

// Update updates a collection in the database.
func (r *CollectionRepository) Update(ctx context.Context, c *collection.Collection) error {
	model := toCollectionModel(c)
	result, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", string(c.ID())).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update collection: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return collection.ErrCollectionNotFound
	}

	return nil
}

// Delete removes a collection from the database by its ID.
func (r *CollectionRepository) Delete(ctx context.Context, id collection.ID) error {
	result, err := r.db.NewDelete().
		Model(&models.Collection{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return collection.ErrCollectionNotFound
	}

	return nil
}
