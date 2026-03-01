package collection

import (
	"context"
	"errors"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/collection"
)

// Service defines the application layer contract for collection operations.
// This service orchestrates business logic and coordinates with the repository layer.
// It enforces validation rules and transforms domain entities into DTOs.
type Service interface {
	// Get retrieves a collection by its unique identifier.
	// Returns ErrCollectionNotFound if the collection does not exist.
	Get(
		ctx context.Context,
		id collection.ID,
	) (*Response, error)

	// ListByNamespace retrieves all collections for a specific database namespace.
	// Returns ErrCollectionNotFound if no collections exist for the namespace.
	// Returns ErrEmptyNamespace if namespace is empty.
	ListByNamespace(
		ctx context.Context,
		databaseNamespaceID string,
	) (*ListByNamespaceResponse, error)

	// Delete removes a collection from the data store.
	// Returns ErrCollectionNotFound if the collection does not exist.
	Delete(ctx context.Context, id collection.ID) error
}

type collectionService struct {
	repo Repository
}

// NewService creates a new CollectionService instance.
func NewService(repo Repository) Service {
	return &collectionService{
		repo: repo,
	}
}

func (s *collectionService) Get(
	ctx context.Context,
	id collection.ID,
) (*Response, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, collection.ErrCollectionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get collection: %w", err)
	}

	return &Response{
		ID:             string(entity.ID()),
		Name:           entity.Name(),
		EmbeddingID:    entity.EmbeddingProfileID(),
		VectorCount:    entity.VectorCount(),
		MetadataSchema: entity.MetadataSchemaJSON(),
		CreatedAt:      entity.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:      entity.UpdatedAt().Format("2006-01-02 15:04:05"),
		LastSyncedAt:   entity.LastSyncedAt().Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *collectionService) ListByNamespace(
	ctx context.Context,
	databaseNamespaceID string,
) (*ListByNamespaceResponse, error) {
	// Validate databaseNamespaceID is not empty
	if databaseNamespaceID == "" {
		return nil, errors.New("database namespace ID cannot be empty")
	}

	entities, err := s.repo.FindByDatabaseNamespaceID(ctx, databaseNamespaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list collections: %w", err)
	}

	// Convert domain entities to DTOs
	responses := make([]*Response, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, &Response{
			ID:             string(entity.ID()),
			Name:           entity.Name(),
			EmbeddingID:    entity.EmbeddingProfileID(),
			VectorCount:    entity.VectorCount(),
			MetadataSchema: entity.MetadataSchemaJSON(),
			CreatedAt:      entity.CreatedAt().Format("2006-01-02 15:04:05"),
			UpdatedAt:      entity.UpdatedAt().Format("2006-01-02 15:04:05"),
			LastSyncedAt:   entity.LastSyncedAt().Format("2006-01-02 15:04:05"),
		})
	}

	return &ListByNamespaceResponse{
		Collections: responses,
		Total:       int64(len(responses)),
	}, nil
}

func (s *collectionService) Delete(ctx context.Context, id collection.ID) error {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return collection.ErrCollectionNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to get collection: %w", err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %w", err)
	}

	return nil
}
