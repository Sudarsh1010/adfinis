package handlers

import (
	"net/http"

	appcollection "github.com/sudarsh1010/adfinis/internal/application/collection"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
)

type CollectionHandler struct {
	service appcollection.Service
}

func NewCollectionHandler(service appcollection.Service) *CollectionHandler {
	return &CollectionHandler{service: service}
}

type CollectionResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	EmbeddingID    string `json:"embeddingId"`
	VectorCount    int    `json:"vectorCount"`
	MetadataSchema string `json:"metadataSchema"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	LastSyncedAt   string `json:"lastSyncedAt"`
}

func (h *CollectionHandler) ListByNamespace(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		response.BadRequest(w, "Namespace ID is required")
		return
	}

	resp, err := h.service.ListByNamespace(r.Context(), idParam)
	if err != nil {
		response.InternalError(w, "Failed to list collections")
		return
	}

	collections := make([]*CollectionResponse, 0, len(resp.Collections))
	for _, c := range resp.Collections {
		collections = append(collections, &CollectionResponse{
			ID:             c.ID,
			Name:           c.Name,
			EmbeddingID:    c.EmbeddingID,
			VectorCount:    c.VectorCount,
			MetadataSchema: c.MetadataSchema,
			CreatedAt:      c.CreatedAt,
			UpdatedAt:      c.UpdatedAt,
			LastSyncedAt:   c.LastSyncedAt,
		})
	}

	response.OK(w, map[string]interface{}{
		"collections": collections,
		"total":       resp.Total,
	})
}
