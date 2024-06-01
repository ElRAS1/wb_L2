package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
)

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	if r.Method != http.MethodPost {
		err := fmt.Sprintf(`"the Post method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	if err := evn.Decode(*r); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	if err := evn.Validate(); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusServiceUnavailable)
		return
	}
	if _, ok := s.data[evn.Event_name]; ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[evn.Event_name] = evn
	// потом убрать
	data, err := json.MarshalIndent(evn, "", " ")

	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(data), http.StatusOK)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	if r.Method != http.MethodPost {
		err := fmt.Sprintf(`"the Post method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	if err := evn.Decode(*r); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	if err := evn.Validate(); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusServiceUnavailable)
		return
	}
	if _, ok := s.data[evn.Event_name]; !ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[evn.Event_name] = evn
	// потом убрать
	data, err := json.MarshalIndent(evn, "", " ")

	if err != nil {
		s.Logger.Error(err.Error())
	}

	s.response(w, s.responseResult(data), http.StatusOK)
}
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	evn := event.Event{}
	if r.Method != http.MethodPost {
		err := fmt.Sprintf(`"the Post method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	if err := evn.Decode(*r); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	if err := evn.Validate(); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusServiceUnavailable)
		return
	}
	if _, ok := s.data[evn.Event_name]; !ok {
		s.response(w, s.responseError("the event does not exist"), http.StatusConflict)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, evn.Event_name)
	// потом убрать
	data, err := json.MarshalIndent(s.data, "", " ")

	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(data), http.StatusOK)
}
