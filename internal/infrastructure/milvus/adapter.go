// Package milvus provides the Milvus infrastructure adapter implementation.
package milvus

import (
	"context"
	"fmt"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
	domainmilvus "github.com/sudarsh1010/adfinis/internal/domain/milvus"
)

// Adapter implements the MilvusClient interface using the Milvus SDK v2.
type Adapter struct {
	endpoint string
	apiKey   string
}

// NewAdapter creates a new Milvus adapter.
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Connect establishes a connection to the Milvus server.
// It stores the connection parameters and validates the connection.
func (a *Adapter) Connect(ctx context.Context, endpoint, apiKey string) error {
	a.endpoint = endpoint
	a.apiKey = apiKey

	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: endpoint,
		APIKey:  apiKey,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}
	defer cli.Close(ctx)

	// Verify connection by listing collections
	_, err = cli.ListCollections(ctx, milvusclient.NewListCollectionOption())
	if err != nil {
		return fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}

	return nil
}

// Ping checks if the Milvus server is reachable.
func (a *Adapter) Ping(ctx context.Context) error {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: a.endpoint,
		APIKey:  a.apiKey,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}
	defer cli.Close(ctx)

	// Verify connection by listing collections
	_, err = cli.ListCollections(ctx, milvusclient.NewListCollectionOption())
	if err != nil {
		return fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}

	return nil
}

// ListCollections retrieves all collections in the database.
func (a *Adapter) ListCollections(
	ctx context.Context,
) ([]domainmilvus.CollectionInfo, error) {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: a.endpoint,
		APIKey:  a.apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}
	defer cli.Close(ctx)

	collectionNames, err := cli.ListCollections(
		ctx,
		milvusclient.NewListCollectionOption(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list collections: %w", err)
	}

	collections := make([]domainmilvus.CollectionInfo, len(collectionNames))
	for i, name := range collectionNames {
		collections[i] = domainmilvus.CollectionInfo{
			Name: name,
		}
	}

	return collections, nil
}

// DescribeCollection retrieves the schema of a specific collection.
func (a *Adapter) DescribeCollection(
	ctx context.Context, name string,
) (*domainmilvus.CollectionSchema, error) {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: a.endpoint,
		APIKey:  a.apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}
	defer cli.Close(ctx)

	collection, err := cli.DescribeCollection(
		ctx, milvusclient.NewDescribeCollectionOption(name),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to describe collection %s: %w",
			name,
			err,
		)
	}

	if collection.Schema == nil {
		return &domainmilvus.CollectionSchema{
			Name:   collection.Name,
			Fields: []domainmilvus.FieldSchema{},
		}, nil
	}

	fields := make([]domainmilvus.FieldSchema, len(collection.Schema.Fields))
	for i, field := range collection.Schema.Fields {
		fields[i] = domainmilvus.FieldSchema{
			Name:     field.Name,
			Type:     field.DataType.String(),
			DataType: field.DataType.Name(),
		}
	}

	return &domainmilvus.CollectionSchema{
		Name:   collection.Schema.CollectionName,
		Fields: fields,
	}, nil
}

// ListDatabases retrieves all databases in the Milvus server.
func (a *Adapter) ListDatabases(
	ctx context.Context,
) ([]domainmilvus.DatabaseInfo, error) {
	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: a.endpoint,
		APIKey:  a.apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domainmilvus.ErrConnectionFailed, err)
	}
	defer cli.Close(ctx)

	databaseNames, err := cli.ListDatabase(
		ctx, milvusclient.NewListDatabaseOption(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list databases: %w", err)
	}

	databases := make([]domainmilvus.DatabaseInfo, len(databaseNames))
	for i, name := range databaseNames {
		databases[i] = domainmilvus.DatabaseInfo{
			Name: name,
		}
	}

	return databases, nil
}
