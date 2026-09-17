package monitor

import (
	"context"
)

type MonitorListItem struct {
	ID     int64
	Name   string
	URL    string
	Active bool
}

type ListMonitorsUseCase interface {
	Execute(
		ctx context.Context,
	) ([]MonitorListItem, error)
}
