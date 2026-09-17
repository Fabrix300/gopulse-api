package monitor

import (
	"context"

	"github.com/Fabrix300/gopulse-api/internal/domain/model"
)

type CreateMonitorService struct {
	monitorRepository MonitorRepository
}

func NewCreateMonitorService(
	monitorRepository MonitorRepository,
) *CreateMonitorService {
	return &CreateMonitorService{
		monitorRepository: monitorRepository,
	}
}

func (s *CreateMonitorService) Execute(
	ctx context.Context,
	command CreateMonitorCommand,
) (CreateMonitorResult, error) {

	monitor, err := model.NewMonitor(
		command.Name,
		command.URL,
	)
	if err != nil {
		return CreateMonitorResult{}, err
	}

	savedMonitor, err := s.monitorRepository.Save(ctx, monitor)
	if err != nil {
		return CreateMonitorResult{}, err
	}

	return CreateMonitorResult{
		ID:     savedMonitor.ID,
		Name:   savedMonitor.Name,
		URL:    savedMonitor.URL,
		Active: savedMonitor.Active,
	}, nil
}
