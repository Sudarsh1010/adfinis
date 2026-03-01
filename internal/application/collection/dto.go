package collection

// Response defines the output when returning collection data.
type Response struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	EmbeddingID    string `json:"embeddingProfileId"`
	VectorCount    int    `json:"vectorCount"`
	MetadataSchema string `json:"metadataSchema"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	LastSyncedAt   string `json:"lastSyncedAt"`
}

// ListByNamespaceResponse represents a paginated list of collections.
type ListByNamespaceResponse struct {
	Collections []*Response
	Total       int64
}
