package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/embedding_profile"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"
	"github.com/uptrace/bun"
)

// EmbeddingProfileRepository implements persistence for EmbeddingProfile domain entity.
type EmbeddingProfileRepository struct {
	db *bun.DB
}

// NewEmbeddingProfileRepository creates a new EmbeddingProfileRepository with the given database connection.
func NewEmbeddingProfileRepository(db *bun.DB) *EmbeddingProfileRepository {
	return &EmbeddingProfileRepository{db: db}
}

// toDomainEntity converts a models.EmbeddingProfile to a domain.EmbeddingProfile entity.
func toEmbeddingProfileDomainEntity(m *models.EmbeddingProfile) *embedding_profile.EmbeddingProfile {
	e := &embedding_profile.EmbeddingProfile{}
	e.SetID(embedding_profile.ID(m.ID))
	e.SetProviderAccountID(m.ProviderAccountID)
	e.SetModelName(m.ModelName)
	e.SetDimension(m.Dimension)
	e.SetMetric(m.Metric)
	e.SetNormalize(m.Normalize)
	e.SetCreatedAt(m.CreatedAt)
	e.SetUpdatedAt(m.UpdatedAt)
	return e
}

// toEmbeddingProfileModel converts a domain.EmbeddingProfile entity to a database model.
func toEmbeddingProfileModel(e *embedding_profile.EmbeddingProfile) *models.EmbeddingProfile {
	return &models.EmbeddingProfile{
		ID:                string(e.ID()),
		ProviderAccountID: e.ProviderAccountID(),
		ModelName:         e.ModelName(),
		Dimension:         e.Dimension(),
		Metric:            e.Metric(),
		Normalize:         e.Normalize(),
		CreatedAt:         e.CreatedAt(),
		UpdatedAt:         e.UpdatedAt(),
	}
}

// Save persists an embedding profile entity to the database.
func (r *EmbeddingProfileRepository) Save(ctx context.Context, e *embedding_profile.EmbeddingProfile) error {
	model := toEmbeddingProfileModel(e)
	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves an embedding profile by its ID from the database.
func (r *EmbeddingProfileRepository) FindByID(ctx context.Context, id embedding_profile.ID) (*embedding_profile.EmbeddingProfile, error) {
	var model models.EmbeddingProfile
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("embedding profile not found: %w", err)
	}
	return toEmbeddingProfileDomainEntity(&model), nil
}

// FindAll retrieves all embedding profiles from the database.
func (r *EmbeddingProfileRepository) FindAll(ctx context.Context) ([]*embedding_profile.EmbeddingProfile, error) {
	var modelList []models.EmbeddingProfile
	err := r.db.NewSelect().Model(&modelList).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all embedding profiles: %w", err)
	}

	profiles := make([]*embedding_profile.EmbeddingProfile, len(modelList))
	for i := range modelList {
		profiles[i] = toEmbeddingProfileDomainEntity(&modelList[i])
	}
	return profiles, nil
}

// Update updates an embedding profile in the database.
func (r *EmbeddingProfileRepository) Update(ctx context.Context, e *embedding_profile.EmbeddingProfile) error {
	model := toEmbeddingProfileModel(e)
	result, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", string(e.ID())).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update embedding profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return embedding_profile.ErrEmbeddingProfileNotFound
	}

	return nil
}

// Delete removes an embedding profile from the database by its ID.
func (r *EmbeddingProfileRepository) Delete(ctx context.Context, id embedding_profile.ID) error {
	result, err := r.db.NewDelete().
		Model(&models.EmbeddingProfile{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete embedding profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return embedding_profile.ErrEmbeddingProfileNotFound
	}

	return nil
}
