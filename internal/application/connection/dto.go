package connection

// CreateConnectionRequest represents the data needed to create a new connection.
type CreateConnectionRequest struct {
	WorkspaceID string
	Name        string
	Provider    string
	Endpoint    string
	APIKey      string
	Region      string
}

// UpdateConnectionRequest represents the data needed to update an existing connection.
type UpdateConnectionRequest struct {
	WorkspaceID string
	Name        string
	Provider    string
	Endpoint    string
	Region      string
}

// GetConnectionResponse represents a connection in the response.
type GetConnectionResponse struct {
	ID              string
	WorkspaceID     string
	Name            string
	Provider        string
	Endpoint        string
	APIKeyEncrypted string
	Region          string
	CreatedAt       string
	UpdatedAt       string
	LastConnectedAt string
}

// ListConnectionsResponse represents a paginated list of connections.
type ListConnectionsResponse struct {
	Connections []*GetConnectionResponse
	Total       int64
}
