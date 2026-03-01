package collection

import (
	"context"

	"github.com/sudarsh1010/adfinis/internal/domain/collection"
)

// Repository defines the contract for collection data access.
// This interface isolates the application layer from infrastructure concerns.
// The implementation (in infrastructure/persistence/bun/repositories) provides
// actual database operations using Bun ORM.
type Repository interface {
	// FindByID retrieves a collection by its unique identifier.
	// Returns ErrCollectionNotFound if the collection does not exist.
	FindByID(ctx context.Context, id collection.ID) (*collection.Collection, error)

	// FindByDatabaseNamespaceID retrieves all collections for a specific database namespace.
	FindByDatabaseNamespaceID(ctx context.Context, databaseNamespaceID string) ([]*collection.Collection, error)

	// Delete removes a collection from the data store by its identifier.
	// Returns ErrCollectionNotFound if the collection does not exist.
	Delete(ctx context.Context, id collection.ID) error
}
