package model

import "errors"

var (
	ErrInvalidMonitorName = errors.New("monitor name is required")
	ErrInvalidMonitorURL  = errors.New("monitor URL is required")
)

type Monitor struct {
	ID     int64
	Name   string
	URL    string
	Active bool
}

func NewMonitor(name string, url string) (Monitor, error) {
	if name == "" {
		return Monitor{}, ErrInvalidMonitorName
	}

	if url == "" {
		return Monitor{}, ErrInvalidMonitorURL
	}

	return Monitor{
		Name:   name,
		URL:    url,
		Active: true,
	}, nil
}
