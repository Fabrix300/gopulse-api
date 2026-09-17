package http

import (
	"encoding/json"
	"net/http"

	"github.com/Fabrix300/gopulse-api/internal/application/monitor"
)

func (h *MonitorHandler) createMonitorHandler(
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

func (h *MonitorHandler) listMonitorsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	result, err := h.listMonitorsUseCase.Execute(r.Context())

	if err != nil {
		// TODO: Might need to handle errors somehow...
		return
	}

	response := make([]monitorResponse, 0, len(result))

	for _, monitor := range result {
		response = append(response, monitorResponse{
			ID:     monitor.ID,
			Name:   monitor.Name,
			URL:    monitor.URL,
			Active: monitor.Active,
		})
	}

	writeJSON(w, http.StatusOK, response)
}
