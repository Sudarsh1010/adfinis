package databasenamespace

import (
	"context"

	domaindatabasenamespace "github.com/sudarsh1010/adfinis/internal/domain/database_namespace"
)

// Service defines the application layer contract for database namespace operations.
// This service orchestrates business logic and coordinates with the repository layer.
// It enforces validation rules and transforms domain entities into DTOs.
type Service interface {
	// CreateDatabaseNamespace creates a new database namespace with validation.
	// Returns ErrDatabaseNamespaceNotFound if the connection ID is empty.
	CreateDatabaseNamespace(
		ctx context.Context,
		req *CreateDatabaseNamespaceRequest,
	) (*Response, error)

	// GetDatabaseNamespace retrieves a database namespace by its unique identifier.
	// Returns ErrDatabaseNamespaceNotFound if the namespace does not exist.
	GetDatabaseNamespace(
		ctx context.Context,
		id domaindatabasenamespace.ID,
	) (*Response, error)

	// ListDatabaseNamespaces retrieves all database namespaces.
	// Returns an empty list if no namespaces exist.
	ListDatabaseNamespaces(ctx context.Context) ([]*Response, error)

	// DeleteDatabaseNamespace removes a database namespace from the data store.
	// Returns ErrDatabaseNamespaceNotFound if the namespace does not exist.
	DeleteDatabaseNamespace(ctx context.Context, id domaindatabasenamespace.ID) error
}

// databasenamespaceService is the concrete implementation of DatabaseNamespaceService.
// It uses DatabaseNamespaceRepository for data persistence and transforms between
// domain entities and DTOs.
type databasenamespaceService struct {
	repo Repository
}

// NewDatabaseNamespaceService creates a new DatabaseNamespaceService instance.
func NewDatabaseNamespaceService(repo Repository) Service {
	return &databasenamespaceService{
		repo: repo,
	}
}

// CreateDatabaseNamespace creates a new database namespace with validation.
func (s *databasenamespaceService) CreateDatabaseNamespace(
	ctx context.Context,
	req *CreateDatabaseNamespaceRequest,
) (*Response, error) {
	if req.ConnectionID == "" {
		return nil, domaindatabasenamespace.ErrDatabaseNamespaceNotFound
	}

	entity := domaindatabasenamespace.NewDatabaseNamespace(req.ConnectionID, req.Name)

	err := s.repo.Save(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:           string(entity.ID()),
		ConnectionID: entity.ConnectionID(),
		Name:         entity.Name(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
	}, nil
}

// GetDatabaseNamespace retrieves a database namespace by its unique identifier.
func (s *databasenamespaceService) GetDatabaseNamespace(
	ctx context.Context,
	id domaindatabasenamespace.ID,
) (*Response, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, domaindatabasenamespace.ErrDatabaseNamespaceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &Response{
		ID:           string(entity.ID()),
		ConnectionID: entity.ConnectionID(),
		Name:         entity.Name(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
	}, nil
}

// ListDatabaseNamespaces retrieves all database namespaces.
func (s *databasenamespaceService) ListDatabaseNamespaces(
	ctx context.Context,
) ([]*Response, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*Response, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, &Response{
			ID:           string(entity.ID()),
			ConnectionID: entity.ConnectionID(),
			Name:         entity.Name(),
			CreatedAt:    entity.CreatedAt(),
			UpdatedAt:    entity.UpdatedAt(),
		})
	}

	return responses, nil
}

// DeleteDatabaseNamespace removes a database namespace from the data store.
func (s *databasenamespaceService) DeleteDatabaseNamespace(ctx context.Context, id domaindatabasenamespace.ID) error {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return domaindatabasenamespace.ErrDatabaseNamespaceNotFound
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
