package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// Connection is a Bun model (infrastructure concern ONLY).
type Connection struct {
	bun.BaseModel `bun:"table:connections"`

	ID              string     `bun:"id,pk"`
	WorkspaceID     string     `bun:"workspace_id,notnull,index,references:workspaces(id) on delete cascade"`
	Name            string     `bun:"name,notnull"`
	Provider        string     `bun:"provider,notnull"`
	Endpoint        string     `bun:"endpoint,notnull"`
	APIKeyEncrypted string     `bun:"api_key_encrypted"`
	Region          string     `bun:"region"`
	CreatedAt       time.Time  `bun:"created_at,notnull,default:now()"`
	UpdatedAt       time.Time  `bun:"updated_at,notnull"`
	LastConnectedAt *time.Time `bun:"last_connected_at"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (c *Connection) BeforeUpdate(ctx context.Context) error {
	c.UpdatedAt = time.Now()
	return nil
}
