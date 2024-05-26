package server

import "net/http"

// POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/create_event", logMiddle(http.HandlerFunc(HandleCreate)))
	mux.Handle("/update_event", logMiddle(http.HandlerFunc(HandleUpdate)))
	mux.Handle("/delete_event", logMiddle(http.HandlerFunc(HandleDelete)))
	mux.Handle("/events_for_day", logMiddle(http.HandlerFunc(HandleEventsDay)))
	mux.Handle("/events_for_week", logMiddle(http.HandlerFunc(HandleEventsWeek)))
	mux.Handle("/events_for_month", logMiddle(http.HandlerFunc(HandleEventsMonth)))

	return mux
}
