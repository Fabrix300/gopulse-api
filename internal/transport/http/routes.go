package http

import "net/http"

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /api/v1/monitors", h.listMonitorsHandler)
	mux.HandleFunc("POST /api/v1/monitors", h.createMonitorHandler)

	return mux
}
