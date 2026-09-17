package monitor

import "context"

type CreateMonitorCommand struct {
	Name string
	URL  string
}

type CreateMonitorResult struct {
	ID     int64
	Name   string
	URL    string
	Active bool
}

type CreateMonitorUseCase interface {
	Execute(
		ctx context.Context,
		command CreateMonitorCommand,
	) (CreateMonitorResult, error)
}
