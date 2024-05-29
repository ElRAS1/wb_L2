package server

import "net/http"

// POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
func (s *Server) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/create_event", logMiddle(http.HandlerFunc(s.handleCreate)))
	mux.Handle("/update_event", logMiddle(http.HandlerFunc(s.handleUpdate)))
	mux.Handle("/delete_event", logMiddle(http.HandlerFunc(s.handleDelete)))
	mux.Handle("/events_for_day", logMiddle(http.HandlerFunc(s.handleEventsDay)))
	mux.Handle("/events_for_week", logMiddle(http.HandlerFunc(s.handleEventsWeek)))
	mux.Handle("/events_for_month", logMiddle(http.HandlerFunc(s.handleEventsMonth)))

	return mux
}
