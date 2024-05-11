package health

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// Read godoc
//
//  @summary        Read health / status
//  @description    Read health / status
//  @tags           health status
//  @success        200
//  @router         /../status [get]

func Read(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{
		Status: "OK",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
