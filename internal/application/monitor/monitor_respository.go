package monitor

import (
	"context"

	"github.com/Fabrix300/gopulse-api/internal/domain/model"
)

type MonitorRepository interface {
	Save(
		ctx context.Context,
		monitor model.Monitor,
	) (model.Monitor, error)

	List(
		ctx context.Context,
	) ([]model.Monitor, error)
}
