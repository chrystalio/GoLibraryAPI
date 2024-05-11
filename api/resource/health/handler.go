package health

import (
	"encoding/json"
	"net/http"
)

// HealthCheckResponse represents the response format for the health check endpoint
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func Read(w http.ResponseWriter, r *http.Request) {
	// Create a HealthCheckResponse struct
	response := HealthCheckResponse{
		Status: "OK",
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	// Write the JSON response
	w.Write(responseJSON)
}
