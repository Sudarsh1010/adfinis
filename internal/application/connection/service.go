package connection

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/sudarsh1010/adfinis/internal/domain/connection"
	"github.com/sudarsh1010/adfinis/internal/domain/milvus"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/crypto"
)

// Service defines the application layer contract for connection operations.
// This service orchestrates business logic and coordinates with the repository layer.
// It enforces validation rules and transforms domain entities into DTOs.
type Service interface {
	// CreateConnection creates a new connection with validation.
	// Returns ErrInvalidName if the name is empty.
	// Returns ErrUnsupportedProvider if the provider is not milvus.
	CreateConnection(
		ctx context.Context,
		req *CreateConnectionRequest,
	) (*GetConnectionResponse, error)

	// GetConnectionByID retrieves a connection by its unique identifier.
	// Returns ErrConnectionNotFound if the connection does not exist.
	GetConnectionByID(
		ctx context.Context,
		id connection.ID,
	) (*GetConnectionResponse, error)

	// ListConnections retrieves all connections.
	// Returns an empty list if no connections exist.
	ListConnections(ctx context.Context) (*ListConnectionsResponse, error)

	// UpdateConnection updates an existing connection.
	// Returns ErrConnectionNotFound if the connection does not exist.
	UpdateConnection(
		ctx context.Context,
		id connection.ID,
		req *UpdateConnectionRequest,
	) (*GetConnectionResponse, error)

	// DeleteConnection removes a connection from the data store.
	// Returns ErrConnectionNotFound if the connection does not exist.
	DeleteConnection(ctx context.Context, id connection.ID) error

	// TestConnection tests the connection by decrypting the API key
	// and calling Milvus Ping().
	TestConnection(
		ctx context.Context,
		id connection.ID,
	) error
}

// connectionService is the concrete implementation of ConnectionService.
// It uses ConnectionRepository for data persistence and transforms between
// domain entities and DTOs.
type connectionService struct {
	repo         Repository
	encryptor    *crypto.Service
	milvusClient milvus.MilvusClient
}

// NewService creates a new ConnectionService instance.
func NewService(
	repo Repository,
	encryptor *crypto.Service,
	milvusClient milvus.MilvusClient,
) Service {
	return &connectionService{
		repo:         repo,
		encryptor:    encryptor,
		milvusClient: milvusClient,
	}
}

// CreateConnection creates a new connection with validation.
func (s *connectionService) CreateConnection(
	ctx context.Context,
	req *CreateConnectionRequest,
) (*GetConnectionResponse, error) {
	// Validate provider is milvus
	if req.Provider != "milvus" {
		return nil, connection.ErrInvalidProvider
	}
	if req.Provider != "milvus" {
		return nil, errors.New("unsupported provider: " + req.Provider)
	}

	// Validate endpoint is a valid URL
	if _, err := url.ParseRequestURI(req.Endpoint); err != nil {
		return nil, connection.ErrInvalidEndpoint
	}
	if _, err := url.ParseRequestURI(req.Endpoint); err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	// Encrypt API key before saving
	encryptedKey, err := s.encryptor.Encrypt(req.APIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt API key: %w", err)
	}

	// Create domain entity
	entity := connection.NewConnection(
		req.WorkspaceID,
		req.Name,
		req.Provider,
		req.Endpoint,
	)

	// Set encrypted API key and region
	entity.SetAPIKeyEncrypted(encryptedKey)
	entity.SetRegion(req.Region)

	// Save to repository
	err = s.repo.Save(ctx, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to save connection: %w", err)
	}

	// Return response
	return &GetConnectionResponse{
		ID:              string(entity.ID()),
		WorkspaceID:     entity.WorkspaceID(),
		Name:            entity.Name(),
		Provider:        entity.Provider(),
		Endpoint:        entity.Endpoint(),
		APIKeyEncrypted: entity.APIKeyEncrypted(),
		Region:          entity.Region(),
		CreatedAt:       entity.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:       entity.UpdatedAt().Format("2006-01-02 15:04:05"),
		LastConnectedAt: entity.LastConnectedAt().Format("2006-01-02 15:04:05"),
	}, nil
}

// GetConnectionByID retrieves a connection by its unique identifier.
func (s *connectionService) GetConnectionByID(
	ctx context.Context,
	id connection.ID,
) (*GetConnectionResponse, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, connection.ErrConnectionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	return &GetConnectionResponse{
		ID:              string(entity.ID()),
		WorkspaceID:     entity.WorkspaceID(),
		Name:            entity.Name(),
		Provider:        entity.Provider(),
		Endpoint:        entity.Endpoint(),
		APIKeyEncrypted: entity.APIKeyEncrypted(),
		Region:          entity.Region(),
		CreatedAt:       entity.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:       entity.UpdatedAt().Format("2006-01-02 15:04:05"),
		LastConnectedAt: entity.LastConnectedAt().Format("2006-01-02 15:04:05"),
	}, nil
}

// ListConnections retrieves all connections.
func (s *connectionService) ListConnections(ctx context.Context) (*ListConnectionsResponse, error) {
	entities, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list connections: %w", err)
	}

	responses := make([]*GetConnectionResponse, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, &GetConnectionResponse{
			ID:              string(entity.ID()),
			WorkspaceID:     entity.WorkspaceID(),
			Name:            entity.Name(),
			Provider:        entity.Provider(),
			Endpoint:        entity.Endpoint(),
			APIKeyEncrypted: entity.APIKeyEncrypted(),
			Region:          entity.Region(),
			CreatedAt:       entity.CreatedAt().Format("2006-01-02 15:04:05"),
			UpdatedAt:       entity.UpdatedAt().Format("2006-01-02 15:04:05"),
			LastConnectedAt: entity.LastConnectedAt().Format("2006-01-02 15:04:05"),
		})
	}

	return &ListConnectionsResponse{
		Connections: responses,
		Total:       int64(len(responses)),
	}, nil
}

// UpdateConnection updates an existing connection.
func (s *connectionService) UpdateConnection(
	ctx context.Context,
	id connection.ID,
	req *UpdateConnectionRequest,
) (*GetConnectionResponse, error) {
	// Get existing entity
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return nil, connection.ErrConnectionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	// Update fields
	entity.SetWorkspaceID(req.WorkspaceID)
	entity.SetName(req.Name)
	entity.SetProvider(req.Provider)
	entity.SetEndpoint(req.Endpoint)
	entity.SetRegion(req.Region)

	// Save to repository
	err = s.repo.Update(ctx, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to update connection: %w", err)
	}

	return &GetConnectionResponse{
		ID:              string(entity.ID()),
		WorkspaceID:     entity.WorkspaceID(),
		Name:            entity.Name(),
		Provider:        entity.Provider(),
		Endpoint:        entity.Endpoint(),
		APIKeyEncrypted: entity.APIKeyEncrypted(),
		Region:          entity.Region(),
		CreatedAt:       entity.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:       entity.UpdatedAt().Format("2006-01-02 15:04:05"),
		LastConnectedAt: entity.LastConnectedAt().Format("2006-01-02 15:04:05"),
	}, nil
}

// DeleteConnection removes a connection from the data store.
func (s *connectionService) DeleteConnection(ctx context.Context, id connection.ID) error {
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return connection.ErrConnectionNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete connection: %w", err)
	}

	return nil
}

// TestConnection tests the connection by decrypting the API key
// and calling Milvus Ping().
func (s *connectionService) TestConnection(
	ctx context.Context,
	id connection.ID,
) error {
	// Get existing entity
	entity, err := s.repo.FindByID(ctx, id)
	if entity == nil {
		return connection.ErrConnectionNotFound
	}
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}

	// Decrypt API key
	apiKey, err := s.encryptor.Decrypt(entity.APIKeyEncrypted())
	if err != nil {
		return fmt.Errorf("failed to decrypt API key: %w", err)
	}

	// Connect to Milvus using decrypted API key
	err = s.milvusClient.Connect(ctx, entity.Endpoint(), apiKey)
	if err != nil {
		return fmt.Errorf("failed to connect to Milvus: %w", err)
	}

	// Ping the Milvus server
	err = s.milvusClient.Ping(ctx)
	if err != nil {
		return fmt.Errorf("Milvus Ping failed: %w", err)
	}

	// Update last connected at time
	now := entity.UpdatedAt()
	entity.SetUpdatedAt(now)
	entity.SetLastConnectedAt(now)
	s.repo.Update(ctx, entity)

	return nil
}
