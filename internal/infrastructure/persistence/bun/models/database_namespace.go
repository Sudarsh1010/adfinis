package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// DatabaseNamespace is a Bun model (infrastructure concern ONLY).
type DatabaseNamespace struct {
	bun.BaseModel `bun:"table:database_namespaces"`

	ID           string    `bun:"id,pk"`
	ConnectionID string    `bun:"connection_id,notnull,index,references:connections(id) on delete cascade"`
	Name         string    `bun:"name,notnull"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time `bun:"updated_at,notnull"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (d *DatabaseNamespace) BeforeUpdate(ctx context.Context) error {
	d.UpdatedAt = time.Now()
	return nil
}
