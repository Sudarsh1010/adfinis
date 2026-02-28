package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// ProviderAccount is a Bun model (infrastructure concern ONLY).
type ProviderAccount struct {
	bun.BaseModel `bun:"table:provider_accounts"`

	ID              string    `bun:"id,pk"`
	WorkspaceID     string    `bun:"workspace_id,notnull,index,references:workspaces(id) on delete cascade"`
	Provider        string    `bun:"provider,notnull"`
	Name            string    `bun:"name,notnull"`
	APIKeyEncrypted string    `bun:"api_key_encrypted,notnull"`
	ExtraConfigJSON string    `bun:"extra_config_json"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt       time.Time `bun:"updated_at,notnull"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (p *ProviderAccount) BeforeUpdate(ctx context.Context) error {
	p.UpdatedAt = time.Now()
	return nil
}
