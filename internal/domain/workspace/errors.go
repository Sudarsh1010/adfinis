package workspace

import "errors"

var (
	ErrInvalidName       = errors.New("workspace name is invalid")
	ErrWorkspaceNotFound = errors.New("workspace not found")
)
