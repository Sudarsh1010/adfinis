package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	appconnection "github.com/sudarsh1010/adfinis/internal/application/connection"
	domainconnection "github.com/sudarsh1010/adfinis/internal/domain/connection"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
)

type ConnectionHandler struct {
	service appconnection.Service
}

func NewConnectionHandler(service appconnection.Service) *ConnectionHandler {
	return &ConnectionHandler{service: service}
}

type CreateConnectionRequest struct {
	WorkspaceID string `json:"workspaceId"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Endpoint    string `json:"endpoint"`
	APIKey      string `json:"apiKey"`
	Region      string `json:"region"`
}

type ConnectionResponse struct {
	ID              string `json:"id"`
	WorkspaceID     string `json:"workspaceId"`
	Name            string `json:"name"`
	Provider        string `json:"provider"`
	Endpoint        string `json:"endpoint"`
	APIKeyEncrypted string `json:"apiKeyEncrypted"`
	Region          string `json:"region"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	LastConnectedAt string `json:"lastConnectedAt"`
}

func (h *ConnectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get workspaceId from URL path
	workspaceID := chi.URLParam(r, "workspaceId")
	if workspaceID == "" {
		response.BadRequest(w, "Workspace ID is required")
		return
	}

	var req CreateConnectionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	// Use workspaceId from URL path
	resp, err := h.service.CreateConnection(
		r.Context(),
		&appconnection.CreateConnectionRequest{
			WorkspaceID: workspaceID,
			Name:        req.Name,
			Provider:    req.Provider,
			Endpoint:    req.Endpoint,
			APIKey:      req.APIKey,
			Region:      req.Region,
		},
	)
	if err != nil {
		if errors.Is(err, domainconnection.ErrInvalidProvider) {
			response.BadRequest(
				w,
				"Unsupported provider. Only 'milvus' is supported",
			)
			return
		}
		if errors.Is(err, domainconnection.ErrInvalidEndpoint) {
			response.BadRequest(w, "Invalid endpoint URL")
			return
		}
		if errors.Is(err, domainconnection.ErrConnectionNotFound) {
			response.NotFound(w, "Connection not found")
			return
		}
		log.Printf("CreateConnection error: %v", err)
		response.InternalError(w, "Failed to create connection")
		return

	}

	response.Created(w, ConnectionResponse{
		ID:              resp.ID,
		WorkspaceID:     resp.WorkspaceID,
		Name:            resp.Name,
		Provider:        resp.Provider,
		Endpoint:        resp.Endpoint,
		APIKeyEncrypted: resp.APIKeyEncrypted,
		Region:          resp.Region,
		CreatedAt:       resp.CreatedAt,
		UpdatedAt:       resp.UpdatedAt,
		LastConnectedAt: resp.LastConnectedAt,
	})
}

func (h *ConnectionHandler) List(w http.ResponseWriter, r *http.Request) {
	workspacesParam := r.URL.Query().Get("workspaceId")

	connections, err := h.service.ListConnections(r.Context())
	if err != nil {
		response.InternalError(w, "Failed to list connections")
		return
	}

	// Filter by workspace ID if provided
	if workspacesParam != "" {
		filtered := make([]*appconnection.GetConnectionResponse, 0)
		for _, conn := range connections.Connections {
			if conn.WorkspaceID == workspacesParam {
				filtered = append(filtered, conn)
			}
		}
		connections.Connections = filtered
		connections.Total = int64(len(filtered))
	}

	response.OK(w, connections)
}

func (h *ConnectionHandler) Test(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id := domainconnection.ID(idParam)

	if idParam == "" {
		response.BadRequest(w, "Connection ID is required")
		return
	}

	if err := h.service.TestConnection(r.Context(), id); err != nil {
		if errors.Is(err, domainconnection.ErrConnectionNotFound) {
			response.NotFound(w, "Connection not found")
			return
		}
		response.InternalError(w, "Failed to test connection")
		return
	}

	response.OK(
		w,
		map[string]string{"message": "Connection tested successfully"},
	)
}
