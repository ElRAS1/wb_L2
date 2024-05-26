package server

import (
	"net/http"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateHandle"))
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
func HandleEventsDay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsDay"))
}
func HandleEventsWeek(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handleEventsWeek"))
}
func HandleEventsMonth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsMonth"))
}
