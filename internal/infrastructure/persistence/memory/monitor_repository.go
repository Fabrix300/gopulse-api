package memory

import (
	"context"
	"sync"

	"github.com/Fabrix300/gopulse-api/internal/domain/model"
)

type MonitorRepository struct {
	mu       sync.Mutex
	monitors []model.Monitor
	nextID   int64
}

func NewMonitorRepository() *MonitorRepository {
	return &MonitorRepository{
		nextID: 1,
	}
}

func (r *MonitorRepository) Save(
	ctx context.Context,
	monitor model.Monitor,
) (model.Monitor, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	monitor.ID = r.nextID
	r.nextID++

	r.monitors = append(r.monitors, monitor)

	return monitor, nil
}
