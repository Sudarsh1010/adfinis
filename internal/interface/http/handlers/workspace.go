package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	appworkspace "github.com/sudarsh1010/adfinis/internal/application/workspace"
	domainworkspace "github.com/sudarsh1010/adfinis/internal/domain/workspace"
	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
)

type WorkspaceHandler struct {
	service appworkspace.Service
}

func NewWorkspaceHandler(service appworkspace.Service) *WorkspaceHandler {
	return &WorkspaceHandler{service: service}
}

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

type WorkspaceResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
	CreatedAt string `json:"createdAt"`
}

func (h *WorkspaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateWorkspaceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	resp, err := h.service.CreateWorkspace(r.Context(), &appworkspace.CreateWorkspaceRequest{Name: req.Name})
	if err != nil {
		if err == domainworkspace.ErrInvalidName {
			response.BadRequest(w, "Name cannot be empty")
			return
		}
		response.InternalError(w, "Failed to create workspace")
		return
	}

	response.Created(w, WorkspaceResponse{
		ID:        resp.ID,
		Name:      resp.Name,
		IsDefault: resp.IsDefault,
			CreatedAt: resp.CreatedAt.Format(time.RFC3339),
	})
}

func (h *WorkspaceHandler) List(w http.ResponseWriter, r *http.Request) {
	workspaces, err := h.service.ListWorkspaces(r.Context())
	if err != nil {
		response.InternalError(w, "Failed to list workspaces")
		return
	}

	response.OK(w, workspaces)
}

func (h *WorkspaceHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id := domainworkspace.ID(idParam)

	resp, err := h.service.GetWorkspaceByID(r.Context(), id)
	if err != nil {
		if err == domainworkspace.ErrWorkspaceNotFound {
			response.NotFound(w, "Workspace not found")
			return
		}
		response.InternalError(w, "Failed to get workspace")
		return
	}

	response.OK(w, WorkspaceResponse{
		ID:        resp.ID,
		Name:      resp.Name,
		IsDefault: resp.IsDefault,
			CreatedAt: resp.CreatedAt.Format(time.RFC3339),
	})
}

func (h *WorkspaceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id := domainworkspace.ID(idParam)

	if err := h.service.DeleteWorkspace(r.Context(), id); err != nil {
		if err == domainworkspace.ErrWorkspaceNotFound {
			response.NotFound(w, "Workspace not found")
			return
		}
		response.InternalError(w, "Failed to delete workspace")
		return
	}

	response.OK(w, map[string]string{"message": "Workspace deleted"})
}
