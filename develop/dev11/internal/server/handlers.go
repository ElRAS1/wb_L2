package main

import "net/http"

func (app *Application) handleCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateHandle"))
}
func (app *Application) handleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}
func (app *Application) handleDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
func (app *Application) handleEventsDay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsDay"))
}
func (app *Application) handleEventsWeek(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handleEventsWeek"))
}
func (app *Application) handleEventsMonth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsMonth"))
}
