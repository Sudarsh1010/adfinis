package connection

import (
	"context"

	"github.com/sudarsh1010/adfinis/internal/domain/connection"
)

// Repository defines the contract for connection data access.
// This interface isolates the application layer from infrastructure concerns.
// The implementation (in infrastructure/persistence/bun/repositories) provides
// actual database operations using Bun ORM.
type Repository interface {
	// Save persists a new connection to the data store.
	Save(ctx context.Context, c *connection.Connection) error

	// FindByID retrieves a connection by its unique identifier.
	// Returns ErrConnectionNotFound if the connection does not exist.
	FindByID(ctx context.Context, id connection.ID) (*connection.Connection, error)

	// FindAll retrieves all connections from the data store.
	FindAll(ctx context.Context) ([]*connection.Connection, error)

	// Update updates an existing connection in the data store.
	// Returns ErrConnectionNotFound if the connection does not exist.
	Update(ctx context.Context, c *connection.Connection) error

	// Delete removes a connection from the data store by its identifier.
	// Returns ErrConnectionNotFound if the connection does not exist.
	Delete(ctx context.Context, id connection.ID) error
}
