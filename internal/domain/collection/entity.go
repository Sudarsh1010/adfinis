package collection

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewCollectionID() ID {
	return ID(ulid.Make().String())
}

// Collection is a pure domain entity (NO Bun dependencies).
type Collection struct {
	id                  ID
	databaseNamespaceID string
	name                string
	embeddingProfileID  string
	vectorCount         int
	metadataSchemaJSON  string
	createdAt           time.Time
	updatedAt           time.Time
	lastSyncedAt        time.Time
}

func NewCollection(databaseNamespaceID, name, embeddingProfileID string) *Collection {
	return &Collection{
		id:                  NewCollectionID(),
		databaseNamespaceID: databaseNamespaceID,
		name:                name,
		embeddingProfileID:  embeddingProfileID,
		vectorCount:         0,
		metadataSchemaJSON:  "",
		createdAt:           time.Now(),
		updatedAt:           time.Now(),
		lastSyncedAt:        time.Time{},
	}
}

func (c *Collection) ID() ID {
	return c.id
}

func (c *Collection) DatabaseNamespaceID() string {
	return c.databaseNamespaceID
}

func (c *Collection) Name() string {
	return c.name
}

func (c *Collection) EmbeddingProfileID() string {
	return c.embeddingProfileID
}

func (c *Collection) VectorCount() int {
	return c.vectorCount
}

func (c *Collection) MetadataSchemaJSON() string {
	return c.metadataSchemaJSON
}

func (c *Collection) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Collection) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Collection) LastSyncedAt() time.Time {
	return c.lastSyncedAt
}

func (c *Collection) SetID(id ID) {
	c.id = id
}

func (c *Collection) SetDatabaseNamespaceID(databaseNamespaceID string) {
	c.databaseNamespaceID = databaseNamespaceID
}

func (c *Collection) SetName(name string) {
	c.name = name
}

func (c *Collection) SetEmbeddingProfileID(embeddingProfileID string) {
	c.embeddingProfileID = embeddingProfileID
}

func (c *Collection) SetVectorCount(vectorCount int) {
	c.vectorCount = vectorCount
}

func (c *Collection) SetMetadataSchemaJSON(metadataSchemaJSON string) {
	c.metadataSchemaJSON = metadataSchemaJSON
}

func (c *Collection) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Collection) SetUpdatedAt(updatedAt time.Time) {
	c.updatedAt = updatedAt
}

func (c *Collection) SetLastSyncedAt(lastSyncedAt time.Time) {
	c.lastSyncedAt = lastSyncedAt
}
