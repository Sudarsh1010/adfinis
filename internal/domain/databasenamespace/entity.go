package databasenamespace

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewID() ID {
	return ID(ulid.Make().String())
}

// DatabaseNamespace is a pure domain entity (NO Bun dependencies).
type DatabaseNamespace struct {
	id           ID
	name         string
	connectionID string
	createdAt    time.Time
}

func NewDatabaseNamespace(name, connectionID string) *DatabaseNamespace {
	return &DatabaseNamespace{
		id:           NewID(),
		name:         name,
		connectionID: connectionID,
		createdAt:    time.Now(),
	}
}

func (ns *DatabaseNamespace) ID() ID {
	return ns.id
}

func (ns *DatabaseNamespace) Name() string {
	return ns.name
}

func (ns *DatabaseNamespace) ConnectionID() string {
	return ns.connectionID
}

func (ns *DatabaseNamespace) CreatedAt() time.Time {
	return ns.createdAt
}

// Domain behavior

func (ns *DatabaseNamespace) SetName(name string) {
	ns.name = name
}

func (ns *DatabaseNamespace) SetConnectionID(connectionID string) {
	ns.connectionID = connectionID
}

func (ns *DatabaseNamespace) SetCreatedAt(createdAt time.Time) {
	ns.createdAt = createdAt
}

func (ns *DatabaseNamespace) SetID(id ID) {
	ns.id = id
}
