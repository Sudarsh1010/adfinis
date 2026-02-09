package workspace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sudarsh1010/adfinis/internal/domain/workspace"
)

// mockWorkspaceRepository is a mock implementation of WorkspaceRepository for testing.
type mockWorkspaceRepository struct {
	mock.Mock
}

func (m *mockWorkspaceRepository) Save(ctx context.Context, ws *workspace.Workspace) error {
	args := m.Called(ctx, ws)
	return args.Error(0)
}

func (m *mockWorkspaceRepository) FindByID(ctx context.Context, id workspace.WorkspaceID) (*workspace.Workspace, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*workspace.Workspace), args.Error(1)
}

func (m *mockWorkspaceRepository) FindAll(ctx context.Context) ([]*workspace.Workspace, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*workspace.Workspace), args.Error(1)
}

func (m *mockWorkspaceRepository) UpdateName(ctx context.Context, id workspace.WorkspaceID, name string) error {
	args := m.Called(ctx, id, name)
	return args.Error(0)
}

func (m *mockWorkspaceRepository) Delete(ctx context.Context, id workspace.WorkspaceID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// TestCreateWorkspace_Success tests successful workspace creation.
func TestCreateWorkspace_Success(t *testing.T) {
	// Arrange
	req := &CreateWorkspaceRequest{Name: "Test Workspace"}
	mockRepo := new(mockWorkspaceRepository)

	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*workspace.Workspace")).Return(nil)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.CreateWorkspace(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ID)
	assert.Equal(t, "Test Workspace", resp.Name)
	assert.False(t, resp.IsDefault)
	mockRepo.AssertExpectations(t)
}

// TestCreateWorkspace_EmptyName tests validation for empty workspace name.
func TestCreateWorkspace_EmptyName(t *testing.T) {
	// Arrange
	req := &CreateWorkspaceRequest{Name: ""}
	mockRepo := new(mockWorkspaceRepository)
	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.CreateWorkspace(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.ErrorIs(t, err, workspace.ErrInvalidName)
	mockRepo.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
}

// TestGetWorkspaceByID_Success tests successful workspace retrieval.
func TestGetWorkspaceByID_Success(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")
	expectedWorkspace := workspace.NewWorkspace("Test Workspace")
	expectedWorkspace.SetID(id)
	expectedWorkspace.SetIsDefault(false)

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("FindByID", mock.Anything, id).Return(expectedWorkspace, nil)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.GetWorkspaceByID(context.Background(), id)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, string(id), resp.ID)
	assert.Equal(t, "Test Workspace", resp.Name)
	assert.False(t, resp.IsDefault)
	mockRepo.AssertExpectations(t)
}

// TestGetWorkspaceByID_NotFound tests retrieval of non-existent workspace.
func TestGetWorkspaceByID_NotFound(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("FindByID", mock.Anything, id).Return(nil, workspace.ErrWorkspaceNotFound)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.GetWorkspaceByID(context.Background(), id)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.ErrorIs(t, err, workspace.ErrWorkspaceNotFound)
	mockRepo.AssertExpectations(t)
}

// TestListWorkspaces_Success tests successful listing of workspaces.
func TestListWorkspaces_Success(t *testing.T) {
	// Arrange
	ws1 := workspace.NewWorkspace("Workspace 1")
	ws1.SetID(workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1"))
	ws1.SetIsDefault(false)

	ws2 := workspace.NewWorkspace("Workspace 2")
	ws2.SetID(workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S2"))
	ws2.SetIsDefault(true)

	expectedWorkspaces := []*workspace.Workspace{ws1, ws2}

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("FindAll", mock.Anything).Return(expectedWorkspaces, nil)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.ListWorkspaces(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Workspace 1", resp[0].Name)
	assert.Equal(t, "Workspace 2", resp[1].Name)
	assert.False(t, resp[0].IsDefault)
	assert.True(t, resp[1].IsDefault)
	mockRepo.AssertExpectations(t)
}

// TestUpdateWorkspaceName_Success tests successful workspace name update.
func TestUpdateWorkspaceName_Success(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")
	newName := "New Name"

	expectedWorkspace := workspace.NewWorkspace(newName)
	expectedWorkspace.SetID(id)
	expectedWorkspace.SetIsDefault(false)

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("UpdateName", mock.Anything, id, newName).Return(nil)
	mockRepo.On("FindByID", mock.Anything, id).Return(expectedWorkspace, nil)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.UpdateWorkspaceName(context.Background(), id, newName)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, string(id), resp.ID)
	assert.Equal(t, newName, resp.Name)
	mockRepo.AssertExpectations(t)
}

// TestUpdateWorkspaceName_NotFound tests updating non-existent workspace name.
func TestUpdateWorkspaceName_NotFound(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")
	newName := "New Name"

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("UpdateName", mock.Anything, id, newName).Return(workspace.ErrWorkspaceNotFound)

	service := NewWorkspaceService(mockRepo)

	// Act
	resp, err := service.UpdateWorkspaceName(context.Background(), id, newName)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.ErrorIs(t, err, workspace.ErrWorkspaceNotFound)
	mockRepo.AssertExpectations(t)
}

// TestDeleteWorkspace_Success tests successful workspace deletion.
func TestDeleteWorkspace_Success(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")

	mockRepo := new(mockWorkspaceRepository)
	expectedWorkspace := workspace.NewWorkspace("Test Workspace")
	expectedWorkspace.SetID(id)
	mockRepo.On("FindByID", mock.Anything, id).Return(expectedWorkspace, nil)
	mockRepo.On("Delete", mock.Anything, id).Return(nil)

	service := NewWorkspaceService(mockRepo)

	// Act
	err := service.DeleteWorkspace(context.Background(), id)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestDeleteWorkspace_NotFound tests deletion of non-existent workspace.
func TestDeleteWorkspace_NotFound(t *testing.T) {
	// Arrange
	id := workspace.WorkspaceID("01H7V8F0G1RJZ7B0G1S1G1S1")

	mockRepo := new(mockWorkspaceRepository)
	mockRepo.On("FindByID", mock.Anything, id).Return(nil, workspace.ErrWorkspaceNotFound)

	service := NewWorkspaceService(mockRepo)

	// Act
	err := service.DeleteWorkspace(context.Background(), id)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, workspace.ErrWorkspaceNotFound)
	mockRepo.AssertExpectations(t)
}
