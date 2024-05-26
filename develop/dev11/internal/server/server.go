package main

import "net/http"

// POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
func (app *Application) router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", app.handleCreate)
	mux.HandleFunc("/update_event", app.handleUpdate)
	mux.HandleFunc("/delete_event", app.handleDelete)
	mux.HandleFunc("/events_for_day", app.handleEventsDay)
	mux.HandleFunc("/events_for_week", app.handleEventsWeek)
	mux.HandleFunc("/events_for_month", app.handleEventsMonth)

	http.ListenAndServe(":8080", mux)
}
