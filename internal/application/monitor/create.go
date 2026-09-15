package monitor

import (
	"context"

	"github.com/Fabrix300/gopulse-api/internal/domain/model"
)

// TODO: Refactor might be needed to separate the use case from the service layer.

type CreateMonitorCommand struct {
	Name string
	URL  string
}

type CreateMonitorUseCase interface {
	Execute(ctx context.Context, command CreateMonitorCommand) (CreateMonitorResult, error)
}

type CreateMonitorResult struct {
	ID     int64
	Name   string
	URL    string
	Active bool
}

type MonitorRepository interface {
	Save(ctx context.Context, monitor model.Monitor) (model.Monitor, error)
}

// impl

type CreateMonitorService struct {
	repository MonitorRepository
}

func NewCreateMonitorService(repository MonitorRepository) *CreateMonitorService {
	return &CreateMonitorService{
		repository: repository,
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

	savedMonitor, err := s.repository.Save(ctx, monitor)

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
