package repositories_test

import (
	"context"
	"sync"
	"testing"

	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
)

// mockWorkspaceRepository is a simple in-memory implementation for testing.
// It avoids database dependencies and is suitable for unit tests.
type mockWorkspaceRepository struct {
	mu         sync.RWMutex
	workspaces []*workspace.Workspace
}

// newMockWorkspaceRepository creates a new mock repository with empty storage.
func newMockWorkspaceRepository() *mockWorkspaceRepository {
	return &mockWorkspaceRepository{
		workspaces: make([]*workspace.Workspace, 0),
	}
}

// Save adds a workspace to the in-memory storage.
func (m *mockWorkspaceRepository) Save(ctx context.Context, ws *workspace.Workspace) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.workspaces = append(m.workspaces, ws)
	return nil
}

// FindByID retrieves a workspace by its ID from the in-memory storage.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (m *mockWorkspaceRepository) FindByID(ctx context.Context, id workspace.WorkspaceID) (*workspace.Workspace, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, ws := range m.workspaces {
		if ws.ID() == id {
			return ws, nil
		}
	}

	return nil, workspace.ErrWorkspaceNotFound
}

// FindAll retrieves all workspaces from the in-memory storage.
func (m *mockWorkspaceRepository) FindAll(ctx context.Context) ([]*workspace.Workspace, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*workspace.Workspace, len(m.workspaces))
	copy(result, m.workspaces)
	return result, nil
}

// UpdateName changes the name of a workspace in the in-memory storage.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (m *mockWorkspaceRepository) UpdateName(ctx context.Context, id workspace.WorkspaceID, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, ws := range m.workspaces {
		if ws.ID() == id {
			if err := ws.Rename(name); err != nil {
				return err
			}
			return nil
		}
	}

	return workspace.ErrWorkspaceNotFound
}

// Delete removes a workspace from the in-memory storage by its ID.
// Returns ErrWorkspaceNotFound if the workspace does not exist.
func (m *mockWorkspaceRepository) Delete(ctx context.Context, id workspace.WorkspaceID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ws := range m.workspaces {
		if ws.ID() == id {
			m.workspaces = append(m.workspaces[:i], m.workspaces[i+1:]...)
			return nil
		}
	}

	return workspace.ErrWorkspaceNotFound
}

// TestSave_Success tests the Save method with a valid workspace.
func TestSave_Success(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Create a new workspace
	ws := workspace.NewWorkspace("Test Workspace")

	// Save the workspace
	err := mock.Save(ctx, ws)
	if err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	// Verify workspace was saved
	found, err := mock.FindByID(ctx, ws.ID())
	if err != nil {
		t.Fatalf("FindByID() unexpected error: %v", err)
	}

	if found.ID() != ws.ID() {
		t.Errorf("expected ID %s, got %s", ws.ID(), found.ID())
	}

	if found.Name() != ws.Name() {
		t.Errorf("expected name %s, got %s", ws.Name(), found.Name())
	}
}

// TestFindByID_Success tests the FindByID method with an existing workspace.
func TestFindByID_Success(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Create and save a workspace
	ws := workspace.NewWorkspace("Test Workspace")
	mock.Save(ctx, ws)

	// Find the workspace by ID
	found, err := mock.FindByID(ctx, ws.ID())
	if err != nil {
		t.Fatalf("FindByID() unexpected error: %v", err)
	}

	// Verify the found workspace
	if found.ID() != ws.ID() {
		t.Errorf("expected ID %s, got %s", ws.ID(), found.ID())
	}

	if found.Name() != ws.Name() {
		t.Errorf("expected name %s, got %s", ws.Name(), found.Name())
	}
}

// TestFindByID_NotFound tests the FindByID method with a non-existent workspace.
func TestFindByID_NotFound(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Try to find a non-existent workspace
	id := workspace.NewWorkspaceID()
	_, err := mock.FindByID(ctx, id)

	if err != workspace.ErrWorkspaceNotFound {
		t.Errorf("expected ErrWorkspaceNotFound, got %v", err)
	}
}

// TestFindAll_Success tests the FindAll method with multiple workspaces.
func TestFindAll_Success(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Create and save multiple workspaces
	ws1 := workspace.NewWorkspace("Workspace 1")
	ws2 := workspace.NewWorkspace("Workspace 2")
	ws3 := workspace.NewWorkspace("Workspace 3")

	mock.Save(ctx, ws1)
	mock.Save(ctx, ws2)
	mock.Save(ctx, ws3)

	// Find all workspaces
	found, err := mock.FindAll(ctx)
	if err != nil {
		t.Fatalf("FindAll() unexpected error: %v", err)
	}

	// Verify all workspaces were found
	if len(found) != 3 {
		t.Errorf("expected 3 workspaces, got %d", len(found))
	}

	// Verify we can find each workspace by ID
	for _, ws := range found {
		byID, err := mock.FindByID(ctx, ws.ID())
		if err != nil {
			t.Errorf("FindByID(%s) unexpected error: %v", ws.ID(), err)
		}
		if byID.Name() != ws.Name() {
			t.Errorf("expected name %s, got %s", ws.Name(), byID.Name())
		}
	}
}

// TestUpdateName_Success tests the UpdateName method with an existing workspace.
func TestUpdateName_Success(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Create and save a workspace
	ws := workspace.NewWorkspace("Original Name")
	mock.Save(ctx, ws)

	// Update the name
	newName := "Updated Name"
	err := mock.UpdateName(ctx, ws.ID(), newName)
	if err != nil {
		t.Fatalf("UpdateName() unexpected error: %v", err)
	}

	// Verify the name was updated
	found, err := mock.FindByID(ctx, ws.ID())
	if err != nil {
		t.Fatalf("FindByID() unexpected error: %v", err)
	}

	if found.Name() != newName {
		t.Errorf("expected name %s, got %s", newName, found.Name())
	}
}

// TestUpdateName_NotFound tests the UpdateName method with a non-existent workspace.
func TestUpdateName_NotFound(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Try to update a non-existent workspace
	id := workspace.NewWorkspaceID()
	err := mock.UpdateName(ctx, id, "New Name")

	if err != workspace.ErrWorkspaceNotFound {
		t.Errorf("expected ErrWorkspaceNotFound, got %v", err)
	}
}

// TestDelete_Success tests the Delete method with an existing workspace.
func TestDelete_Success(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Create and save a workspace
	ws := workspace.NewWorkspace("To Be Deleted")
	mock.Save(ctx, ws)

	// Delete the workspace
	err := mock.Delete(ctx, ws.ID())
	if err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}

	// Verify the workspace was deleted
	_, err = mock.FindByID(ctx, ws.ID())
	if err != workspace.ErrWorkspaceNotFound {
		t.Errorf("expected ErrWorkspaceNotFound after delete, got %v", err)
	}
}

// TestDelete_NotFound tests the Delete method with a non-existent workspace.
func TestDelete_NotFound(t *testing.T) {
	mock := newMockWorkspaceRepository()
	ctx := context.Background()

	// Try to delete a non-existent workspace
	id := workspace.NewWorkspaceID()
	err := mock.Delete(ctx, id)

	if err != workspace.ErrWorkspaceNotFound {
		t.Errorf("expected ErrWorkspaceNotFound, got %v", err)
	}
}
