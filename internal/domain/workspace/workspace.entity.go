package workspace

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// WorkspaceID is a ULID wrapper (value object)
type WorkspaceID string

func NewWorkspaceID() WorkspaceID {
	return WorkspaceID(ulid.Make().String())
}

// Workspace is a pure domain entity (NO Bun dependencies)
type Workspace struct {
	id        WorkspaceID
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

func (w *Workspace) ID() WorkspaceID {
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
