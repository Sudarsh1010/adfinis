package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func OK(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Created(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func BadRequest(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusBadRequest, Response{
		Success: false,
		Error:   "Bad Request",
		Message: message,
	})
}

func NotFound(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusNotFound, Response{
		Success: false,
		Error:   "Not Found",
		Message: message,
	})
}

func InternalError(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusInternalServerError, Response{
		Success: false,
		Error:   "Internal Server Error",
		Message: message,
	})
}

func Unauthorized(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusUnauthorized, Response{
		Success: false,
		Error:   "Unauthorized",
		Message: message,
	})
}
