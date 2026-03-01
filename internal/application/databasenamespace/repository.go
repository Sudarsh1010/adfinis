package databasenamespace

import (
	"context"

	domaindatabasenamespace "github.com/sudarsh1010/adfinis/internal/domain/database_namespace"
)

// Repository defines the contract for database namespace data access.
// This interface isolates the application layer from infrastructure concerns.
// The implementation (in infrastructure/persistence/bun/repositories) provides
// actual database operations using Bun ORM.
type Repository interface {
	// Save persists a new database namespace to the data store.
	Save(ctx context.Context, ns *domaindatabasenamespace.DatabaseNamespace) error

	// FindByID retrieves a database namespace by its unique identifier.
	// Returns ErrDatabaseNamespaceNotFound if the namespace does not exist.
	FindByID(ctx context.Context, id domaindatabasenamespace.ID) (*domaindatabasenamespace.DatabaseNamespace, error)

	// FindAll retrieves all database namespaces from the data store.
	FindAll(ctx context.Context) ([]*domaindatabasenamespace.DatabaseNamespace, error)

	// Update updates a database namespace in the data store.
	// Returns ErrDatabaseNamespaceNotFound if the namespace does not exist.
	Update(ctx context.Context, ns *domaindatabasenamespace.DatabaseNamespace) error

	// Delete removes a database namespace from the data store by its identifier.
	// Returns ErrDatabaseNamespaceNotFound if the namespace does not exist.
	Delete(ctx context.Context, id domaindatabasenamespace.ID) error
}
