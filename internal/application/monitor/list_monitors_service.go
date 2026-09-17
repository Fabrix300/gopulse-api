package monitor

import (
	"context"
)

type ListMonitorsService struct {
	monitorRepository MonitorRepository
}

func NewListMonitorsService(
	monitorRepository MonitorRepository,
) *ListMonitorsService {
	return &ListMonitorsService{
		monitorRepository: monitorRepository,
	}
}

func (s *ListMonitorsService) Execute(
	ctx context.Context,
) ([]MonitorListItem, error) {

	savedMonitors, err := s.monitorRepository.List(ctx)

	if err != nil {
		// TODO: Enhance?
		return make([]MonitorListItem, 0), err
	}

	monitors := make([]MonitorListItem, 0, len(savedMonitors))

	for _, monitor := range savedMonitors {
		monitors = append(monitors, MonitorListItem{
			ID:     monitor.ID,
			Name:   monitor.Name,
			URL:    monitor.URL,
			Active: monitor.Active,
		})
	}

	return monitors, nil
}
