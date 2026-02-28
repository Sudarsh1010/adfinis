package connection

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewConnectionID() ID {
	return ID(ulid.Make().String())
}

// Connection is a pure domain entity (NO Bun dependencies).
type Connection struct {
	id              ID
	workspaceID     string
	name            string
	provider        string
	endpoint        string
	apiKeyEncrypted string
	region          string
	createdAt       time.Time
	updatedAt       time.Time
	lastConnectedAt time.Time
}

func NewConnection(workspaceID, name, provider, endpoint string) *Connection {
	return &Connection{
		id:              NewConnectionID(),
		workspaceID:     workspaceID,
		name:            name,
		provider:        provider,
		endpoint:        endpoint,
		apiKeyEncrypted: "",
		region:          "",
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
		lastConnectedAt: time.Time{},
	}
}

func (c *Connection) ID() ID {
	return c.id
}

func (c *Connection) WorkspaceID() string {
	return c.workspaceID
}

func (c *Connection) Name() string {
	return c.name
}

func (c *Connection) Provider() string {
	return c.provider
}

func (c *Connection) Endpoint() string {
	return c.endpoint
}

func (c *Connection) APIKeyEncrypted() string {
	return c.apiKeyEncrypted
}

func (c *Connection) Region() string {
	return c.region
}

func (c *Connection) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Connection) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Connection) LastConnectedAt() time.Time {
	return c.lastConnectedAt
}

func (c *Connection) SetID(id ID) {
	c.id = id
}

func (c *Connection) SetWorkspaceID(workspaceID string) {
	c.workspaceID = workspaceID
}

func (c *Connection) SetName(name string) {
	c.name = name
}

func (c *Connection) SetProvider(provider string) {
	c.provider = provider
}

func (c *Connection) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *Connection) SetAPIKeyEncrypted(apiKeyEncrypted string) {
	c.apiKeyEncrypted = apiKeyEncrypted
}

func (c *Connection) SetRegion(region string) {
	c.region = region
}

func (c *Connection) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Connection) SetUpdatedAt(updatedAt time.Time) {
	c.updatedAt = updatedAt
}

func (c *Connection) SetLastConnectedAt(lastConnectedAt time.Time) {
	c.lastConnectedAt = lastConnectedAt
}
