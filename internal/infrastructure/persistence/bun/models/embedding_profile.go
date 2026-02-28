package models

import (
	"context"
	"time"

	"github.com/sudarsh1010/adfinis/internal/domain/embedding_profile"
	"github.com/uptrace/bun"
)

// EmbeddingProfile is a Bun model (infrastructure concern ONLY).
type EmbeddingProfile struct {
	bun.BaseModel `bun:"table:embedding_profiles"`

	ID                string    `bun:"id,pk"`
	ProviderAccountID string    `bun:"provider_account_id,notnull,index,references:provider_accounts(id) on delete cascade"`
	ModelName         string    `bun:"model_name,notnull"`
	Dimension         *int      `bun:"dimension"`
	Metric            string    `bun:"metric,notnull"`
	Normalize         bool      `bun:"normalize,notnull,default:false"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:now()"`
	UpdatedAt         time.Time `bun:"updated_at,notnull"`
}

// BeforeUpdate is a Bun hook that sets UpdatedAt before each update.
func (e *EmbeddingProfile) BeforeUpdate(ctx context.Context) error {
	e.UpdatedAt = time.Now()
	return nil
}

// toDomainEntity converts a models.EmbeddingProfile to a domain.EmbeddingProfile entity.
func toDomainEntity(m *EmbeddingProfile) *embedding_profile.EmbeddingProfile {
	if m == nil {
		return nil
	}
	ef := &embedding_profile.EmbeddingProfile{}
	ef.SetID(embedding_profile.ID(m.ID))
	ef.SetProviderAccountID(m.ProviderAccountID)
	ef.SetModelName(m.ModelName)
	ef.SetDimension(m.Dimension)
	ef.SetMetric(m.Metric)
	ef.SetNormalize(m.Normalize)
	ef.SetCreatedAt(m.CreatedAt)
	ef.SetUpdatedAt(m.UpdatedAt)
	return ef
}
