package handlers

import (
	"net/http"
	"time"

	"github.com/sudarsh1010/adfinis/internal/interface/http/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	response.OK(w, map[string]any{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "adfinis",
		"version":   "0.1.0",
	})
}
