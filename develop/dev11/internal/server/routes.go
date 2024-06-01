package server

import "net/http"

func (s *Server) Router() *http.ServeMux {
	mux := http.NewServeMux()
	// post
	mux.Handle("/create_event", s.logMiddle(http.HandlerFunc(s.handleCreate)))
	mux.Handle("/update_event", s.logMiddle(http.HandlerFunc(s.handleUpdate)))
	mux.Handle("/delete_event", s.logMiddle(http.HandlerFunc(s.handleDelete)))
	// get
	mux.Handle("/events_for_day", s.logMiddle(http.HandlerFunc(s.handleEventsDay)))
	mux.Handle("/events_for_week", s.logMiddle(http.HandlerFunc(s.handleEventsWeek)))
	mux.Handle("/events_for_month", s.logMiddle(http.HandlerFunc(s.handleEventsMonth)))

	return mux
}
