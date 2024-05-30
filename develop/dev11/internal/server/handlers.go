package server

import (
	"net/http"

	"github.com/ElRAS1/wb_L2/develop/dev11/event"
)

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}

	if err := evn.Decode(*r); err != nil {
		s.Logger.Error(err.Error())
		http.Error(w, s.responseError(err), http.StatusBadRequest)
		return
	}

	if err := evn.Validate(); err != nil {
		s.Logger.Error(err.Error())
		http.Error(w, s.responseError(err), http.StatusServiceUnavailable)
		return
	}

	w.Write([]byte(evn.Event_name))
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
func (s *Server) handleEventsDay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsDay"))
}
func (s *Server) handleEventsWeek(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handleEventsWeek"))
}
func (s *Server) handleEventsMonth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("EventsMonth"))
}
