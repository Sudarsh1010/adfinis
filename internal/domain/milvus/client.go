package milvus

import "context"

// MilvusClient defines the read-only interface for interacting with Milvus.
// This is a pure domain interface with no external dependencies.
type MilvusClient interface {
	// Connect establishes a connection to the Milvus server.
	Connect(ctx context.Context, endpoint, apiKey string) error

	// Ping checks if the Milvus server is reachable.
	Ping(ctx context.Context) error

	// ListCollections retrieves all collections in the database.
	ListCollections(ctx context.Context) ([]CollectionInfo, error)

	// DescribeCollection retrieves the schema of a specific collection.
	DescribeCollection(ctx context.Context, name string) (*CollectionSchema, error)

	// ListDatabases retrieves all databases in the Milvus server.
	ListDatabases(ctx context.Context) ([]DatabaseInfo, error)
}
