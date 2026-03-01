package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	databasenamespace "github.com/sudarsh1010/adfinis/internal/domain/databasenamespace"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
)

type DatabaseNamespaceHandler struct{}

func NewDatabaseNamespaceHandler() *DatabaseNamespaceHandler {
	return &DatabaseNamespaceHandler{}
}

type CreateDatabaseNamespaceRequest struct {
	Name string `json:"name"`
}

type DatabaseNamespaceResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ConnectionID string `json:"connectionId"`
	CreatedAt    string `json:"createdAt"`
}

func (h *DatabaseNamespaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse connection ID from URL parameter
	connectionID := chi.URLParam(r, "id")
	if connectionID == "" {
		response.BadRequest(w, "Connection ID is required")
		return
	}

	// Parse JSON body
	var req CreateDatabaseNamespaceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if req.Name == "" {
		response.BadRequest(w, "Namespace name is required")
		return
	}

	// Create namespace using domain entity
	namespace := databasenamespace.NewDatabaseNamespace(req.Name, connectionID)

	response.Created(w, DatabaseNamespaceResponse{
		ID:           string(namespace.ID()),
		Name:         namespace.Name(),
		ConnectionID: connectionID,
		CreatedAt:    namespace.CreatedAt().Format("2006-01-02T15:04:05Z"),
	})
}

func (h *DatabaseNamespaceHandler) List(w http.ResponseWriter, r *http.Request) {
	// Parse connection ID from URL parameter
	connectionID := chi.URLParam(r, "id")
	if connectionID == "" {
		response.BadRequest(w, "Connection ID is required")
		return
	}

	// List namespaces - for now return empty list since there's no persistence
	// TODO: Implement actual list logic when persistence is added
	response.OK(w, []DatabaseNamespaceResponse{})
}

func (h *DatabaseNamespaceHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequest(w, "Namespace ID is required")
		return
	}

	// Get namespace - for now return mock data since there's no persistence
	// TODO: Implement actual get logic when persistence is added
	connectionID := chi.URLParam(r, "id")
	response.OK(w, DatabaseNamespaceResponse{
		ID:           idParam,
		Name:         "Demo Database Namespace",
		ConnectionID: connectionID,
		CreatedAt:    time.Now().Format("2006-01-02T15:04:05Z"),
	})
}

func (h *DatabaseNamespaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequest(w, "Namespace ID is required")
		return
	}

	// Delete namespace - for now return success since there's no persistence
	// TODO: Implement actual delete logic when persistence is added
	response.OK(w, map[string]string{"message": "Namespace deleted"})
}
