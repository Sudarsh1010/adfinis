package workspace

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewWorkspaceID() ID {
	return ID(ulid.Make().String())
}

// Workspace is a pure domain entity (NO Bun dependencies).
type Workspace struct {
	id        ID
	name      string
	isDefault bool
	createdAt time.Time
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		id:        NewWorkspaceID(),
		name:      name,
		isDefault: false,
		createdAt: time.Now(),
	}
}

func (w *Workspace) ID() ID {
	return w.id
}

func (w *Workspace) Name() string {
	return w.name
}

func (w *Workspace) IsDefault() bool {
	return w.isDefault
}

func (w *Workspace) CreatedAt() time.Time {
	return w.createdAt
}

// Domain behavior

func (w *Workspace) ActivateAsDefault() {
	w.isDefault = true
}

func (w *Workspace) Rename(newName string) error {
	if newName == "" {
		return ErrInvalidName
	}
	w.name = newName
	return nil
}

func (w *Workspace) SetID(id ID) {
	w.id = id
}

func (w *Workspace) SetName(name string) {
	w.name = name
}

func (w *Workspace) SetIsDefault(isDefault bool) {
	w.isDefault = isDefault
}

func (w *Workspace) SetCreatedAt(createdAt time.Time) {
	w.createdAt = createdAt
}
