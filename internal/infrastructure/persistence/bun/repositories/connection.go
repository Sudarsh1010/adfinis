package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/connection"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"
	"github.com/uptrace/bun"
)

// ConnectionRepository implements persistence for Connection domain entity.
type ConnectionRepository struct {
	db *bun.DB
}

// NewConnectionRepository creates a new ConnectionRepository with the given database connection.
func NewConnectionRepository(db *bun.DB) *ConnectionRepository {
	return &ConnectionRepository{db: db}
}

// toDomainEntity converts a models.Connection to a domain.Connection entity.
func toConnectionDomainEntity(m *models.Connection) *connection.Connection {
	c := &connection.Connection{}
	c.SetID(connection.ID(m.ID))
	c.SetWorkspaceID(m.WorkspaceID)
	c.SetName(m.Name)
	c.SetProvider(m.Provider)
	c.SetEndpoint(m.Endpoint)
	c.SetAPIKeyEncrypted(m.APIKeyEncrypted)
	c.SetRegion(m.Region)
	c.SetCreatedAt(m.CreatedAt)
	c.SetUpdatedAt(m.UpdatedAt)
	if m.LastConnectedAt != nil {
		c.SetLastConnectedAt(*m.LastConnectedAt)
	}
	return c
}

// toConnectionModel converts a domain.Connection entity to a database model.
func toConnectionModel(c *connection.Connection) *models.Connection {
	m := &models.Connection{
		ID:              string(c.ID()),
		WorkspaceID:     c.WorkspaceID(),
		Name:            c.Name(),
		Provider:        c.Provider(),
		Endpoint:        c.Endpoint(),
		APIKeyEncrypted: c.APIKeyEncrypted(),
		Region:          c.Region(),
		CreatedAt:       c.CreatedAt(),
		UpdatedAt:       c.UpdatedAt(),
	}
	if !c.LastConnectedAt().IsZero() {
		t := c.LastConnectedAt()
		m.LastConnectedAt = &t
	}
	return m
}

// Save persists a connection entity to the database.
func (r *ConnectionRepository) Save(ctx context.Context, c *connection.Connection) error {
	model := toConnectionModel(c)
	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves a connection by its ID from the database.
func (r *ConnectionRepository) FindByID(ctx context.Context, id connection.ID) (*connection.Connection, error) {
	var model models.Connection
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection not found: %w", err)
	}
	return toConnectionDomainEntity(&model), nil
}

// FindAll retrieves all connections from the database.
func (r *ConnectionRepository) FindAll(ctx context.Context) ([]*connection.Connection, error) {
	var modelList []models.Connection
	err := r.db.NewSelect().Model(&modelList).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all connections: %w", err)
	}

	connections := make([]*connection.Connection, len(modelList))
	for i := range modelList {
		connections[i] = toConnectionDomainEntity(&modelList[i])
	}
	return connections, nil
}

// Update updates a connection in the database.
func (r *ConnectionRepository) Update(ctx context.Context, c *connection.Connection) error {
	model := toConnectionModel(c)
	result, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", string(c.ID())).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update connection: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return connection.ErrConnectionNotFound
	}

	return nil
}

// Delete removes a connection from the database by its ID.
func (r *ConnectionRepository) Delete(ctx context.Context, id connection.ID) error {
	result, err := r.db.NewDelete().
		Model(&models.Connection{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete connection: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return connection.ErrConnectionNotFound
	}

	return nil
}
