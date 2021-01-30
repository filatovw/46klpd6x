package api

import (
	"encoding/json"
	"net/http"
)

// HealthCheckResponse status of server availability
type HealthCheckResponse struct {
	Status string `json:"status"`
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal(HealthCheckResponse{Status: "OK"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})
	return mux
}
