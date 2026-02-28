package embedding_profile

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewEmbeddingProfileID() ID {
	return ID(ulid.Make().String())
}

// EmbeddingProfile is a pure domain entity (NO Bun dependencies).
type EmbeddingProfile struct {
	id                ID
	providerAccountID string
	modelName         string
	dimension         *int
	metric            string
	normalize         bool
	createdAt         time.Time
	updatedAt         time.Time
}

func NewEmbeddingProfile(providerAccountID, modelName, metric string, normalize bool) *EmbeddingProfile {
	return &EmbeddingProfile{
		id:                NewEmbeddingProfileID(),
		providerAccountID: providerAccountID,
		modelName:         modelName,
		dimension:         nil,
		metric:            metric,
		normalize:         normalize,
		createdAt:         time.Now(),
		updatedAt:         time.Now(),
	}
}

func (e *EmbeddingProfile) ID() ID {
	return e.id
}

func (e *EmbeddingProfile) ProviderAccountID() string {
	return e.providerAccountID
}

func (e *EmbeddingProfile) ModelName() string {
	return e.modelName
}

func (e *EmbeddingProfile) Dimension() *int {
	return e.dimension
}

func (e *EmbeddingProfile) Metric() string {
	return e.metric
}

func (e *EmbeddingProfile) Normalize() bool {
	return e.normalize
}

func (e *EmbeddingProfile) CreatedAt() time.Time {
	return e.createdAt
}

func (e *EmbeddingProfile) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *EmbeddingProfile) SetID(id ID) {
	e.id = id
}

func (e *EmbeddingProfile) SetProviderAccountID(providerAccountID string) {
	e.providerAccountID = providerAccountID
}

func (e *EmbeddingProfile) SetModelName(modelName string) {
	e.modelName = modelName
}

func (e *EmbeddingProfile) SetDimension(dimension *int) {
	e.dimension = dimension
}

func (e *EmbeddingProfile) SetMetric(metric string) {
	e.metric = metric
}

func (e *EmbeddingProfile) SetNormalize(normalize bool) {
	e.normalize = normalize
}

func (e *EmbeddingProfile) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}

func (e *EmbeddingProfile) SetUpdatedAt(updatedAt time.Time) {
	e.updatedAt = updatedAt
}
