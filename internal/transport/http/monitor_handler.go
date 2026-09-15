package http

import (
	"encoding/json"
	"net/http"

	"github.com/Fabrix300/gopulse-api/internal/application/monitor"
)

func (h *Handler) createMonitorHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request createMonitorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	result, err := h.createMonitorUseCase.Execute(
		r.Context(),
		monitor.CreateMonitorCommand{
			Name: request.Name,
			URL:  request.URL,
		},
	)

	if err != nil {
		// TODO: map application errors to HTTP errors
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	response := monitorResponse{
		ID:     result.ID,
		Name:   result.Name,
		URL:    result.URL,
		Active: result.Active,
	}

	writeJSON(w, http.StatusCreated, response)
}

func (h *Handler) listMonitorsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// TODO: create and call ListMonitorsUseCase

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
