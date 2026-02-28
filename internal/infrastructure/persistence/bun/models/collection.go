package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// Collection is a Bun model (infrastructure concern ONLY).
type Collection struct {
	bun.BaseModel `bun:"table:collections"`

	ID                  string     `bun:"id,pk"`
	DatabaseNamespaceID string     `bun:"database_namespace_id,notnull,index,references:database_namespaces(id) on delete cascade"`
	Name                string     `bun:"name,notnull,unique:database_namespace_id"`
	EmbeddingProfileID  string     `bun:"embedding_profile_id,notnull,index,references:embedding_profiles(id) on delete cascade"`
	VectorCount         int64      `bun:"vector_count,notnull,default:0"`
	MetadataSchemaJSON  string     `bun:"metadata_schema_json"`
	CreatedAt           time.Time  `bun:"created_at,notnull,default:now()"`
	UpdatedAt           time.Time  `bun:"updated_at,notnull"`
	LastSyncedAt        *time.Time `bun:"last_synced_at"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (c *Collection) BeforeUpdate(ctx context.Context) error {
	c.UpdatedAt = time.Now()
	return nil
}
