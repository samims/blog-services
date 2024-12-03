package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// StandardResponse represents the standard structure for API responses.
type StandardResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SetJSONHeader sets the Content-Type header to application/json.
func SetJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// RespondWithJSON sends a JSON response with the given status code, data, and error message.
// It can handle both standard and paginated responses based on the data provided.
func RespondWithJSON(w http.ResponseWriter, code int, data interface{}, errorMessage string) {
	SetJSONHeader(w)
	w.WriteHeader(code)

	response := StandardResponse{
		Success: errorMessage == "",
		Data:    data,
		Error:   errorMessage,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("Unable to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// RespondWithError sends an error response.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, nil, message)
	return
}
