package memory

import (
	"context"
	"sync"

	"github.com/Fabrix300/gopulse-api/internal/domain/model"
)

type MonitorRepository struct {
	mu       sync.RWMutex
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

func (r *MonitorRepository) List(
	ctx context.Context,
) ([]model.Monitor, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	monitors := make([]model.Monitor, len(r.monitors))
	copy(monitors, r.monitors)

	return monitors, nil
}

// A more idiomatic version of "List"?
// func (r *MonitorRepository) List(
// 	ctx context.Context,
// ) ([]model.Monitor, error) {

// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	return append([]model.Monitor(nil), r.monitors...), nil
// }
