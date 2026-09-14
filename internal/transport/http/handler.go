package http

import (
	"encoding/json"
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("GET /api/v1/monitors", listMonitorsHandler)
	mux.HandleFunc("POST /api/v1/monitors", createMonitorHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{"status":"UP"}`))
}

func listMonitorsHandler(w http.ResponseWriter, r *http.Request) {
	monitors := []monitorResponse{
		{
			ID:     1,
			Name:   "Google",
			URL:    "https://google.com",
			Active: true,
		},
		{
			ID:     2,
			Name:   "GitHub",
			URL:    "https://github.com",
			Active: true,
		},
	}

	writeJSON(w, http.StatusOK, monitors)
}

func createMonitorHandler(w http.ResponseWriter, r *http.Request) {
	var request createMonitorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	response := monitorResponse{
		ID:     1,
		Name:   request.Name,
		URL:    request.URL,
		Active: true,
	}

	writeJSON(w, http.StatusCreated, response)
}
