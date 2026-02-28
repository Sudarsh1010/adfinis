package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// Workspace is a Bun model (infrastructure concern ONLY).
type Workspace struct {
	bun.BaseModel `bun:"table:workspaces"`

	ID        string    `bun:"id,pk"`
	Name      string    `bun:"name,notnull"`
	IsDefault bool      `bun:"is_default,notnull,default:false"`
	CreatedAt time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (w *Workspace) BeforeUpdate(ctx context.Context) error {
	w.UpdatedAt = time.Now()
	return nil
}
