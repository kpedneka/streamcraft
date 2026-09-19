package server

import (
	"encoding/json"
	"net/http"
)

type Dependencies struct {
	// Add any dependencies required by the server here.
	// For example, you might include a database connection, configuration settings, etc.
}

func NewRouter(deps Dependencies) http.Handler {
	api := http.NewServeMux()

	// Define your routes and handlers here.
	api.HandleFunc("GET /healthz", healthz)
	return api
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
