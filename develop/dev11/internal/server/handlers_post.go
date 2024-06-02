package server

import (
	"encoding/json"
	"net/http"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
)

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	err := s.check(w, r, evn)
	if err != nil {
		return
	}
	if _, ok := s.data[evn.Event_name]; ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[evn.Event_name] = evn
	data, err := json.MarshalIndent(evn, "", " ")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(data), http.StatusOK)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	err := s.check(w, r, evn)
	if err != nil {
		return
	}
	if _, ok := s.data[evn.Event_name]; !ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[evn.Event_name] = evn
	data, err := json.MarshalIndent(evn, "", " ")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(data), http.StatusOK)
}
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	err := s.check(w, r, evn)
	if err != nil {
		return
	}
	if _, ok := s.data[evn.Event_name]; !ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, evn.Event_name)
	data, err := json.MarshalIndent(s.data, "", " ")
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(data), http.StatusOK)
}
