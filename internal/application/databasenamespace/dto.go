package databasenamespace

import "time"

// CreateDatabaseNamespaceRequest defines the input for creating a new database namespace.
type CreateDatabaseNamespaceRequest struct {
	ConnectionID string `json:"connectionId"`
	Name         string `json:"name"`
}

// Response defines the output when returning database namespace data.
type Response struct {
	ID           string    `json:"id"`
	ConnectionID string    `json:"connectionId"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
