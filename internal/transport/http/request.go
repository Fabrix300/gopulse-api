package http

type createMonitorRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
