package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Workspace is a Bun model (infrastructure concern ONLY).
type Workspace struct {
	bun.BaseModel `bun:"table:workspaces"`

	ID        string    `bun:"id,pk,default:''"`
	Name      string    `bun:"name,notnull"`
	IsDefault bool      `bun:"is_default,notnull,default:false"`
	CreatedAt time.Time `bun:"created_at,notnull,default:now()"`
}
