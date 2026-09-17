package http

import (
	"github.com/Fabrix300/gopulse-api/internal/application/monitor"
)

type MonitorHandler struct {
	createMonitorUseCase monitor.CreateMonitorUseCase
}

func NewHandler(createMonitorUseCase monitor.CreateMonitorUseCase) *MonitorHandler {
	return &MonitorHandler{
		createMonitorUseCase: createMonitorUseCase,
	}
}
