package database_namespace

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewDatabaseNamespaceID() ID {
	return ID(ulid.Make().String())
}

// DatabaseNamespace is a pure domain entity (NO Bun dependencies).
type DatabaseNamespace struct {
	id           ID
	connectionID string
	name         string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewDatabaseNamespace(connectionID, name string) *DatabaseNamespace {
	return &DatabaseNamespace{
		id:           NewDatabaseNamespaceID(),
		connectionID: connectionID,
		name:         name,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
}

func (d *DatabaseNamespace) ID() ID {
	return d.id
}

func (d *DatabaseNamespace) ConnectionID() string {
	return d.connectionID
}

func (d *DatabaseNamespace) Name() string {
	return d.name
}

func (d *DatabaseNamespace) CreatedAt() time.Time {
	return d.createdAt
}

func (d *DatabaseNamespace) UpdatedAt() time.Time {
	return d.updatedAt
}

func (d *DatabaseNamespace) SetID(id ID) {
	d.id = id
}

func (d *DatabaseNamespace) SetConnectionID(connectionID string) {
	d.connectionID = connectionID
}

func (d *DatabaseNamespace) SetName(name string) {
	d.name = name
}

func (d *DatabaseNamespace) SetCreatedAt(createdAt time.Time) {
	d.createdAt = createdAt
}

func (d *DatabaseNamespace) SetUpdatedAt(updatedAt time.Time) {
	d.updatedAt = updatedAt
}
