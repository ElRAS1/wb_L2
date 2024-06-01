package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
	// "github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
)

func (s *Server) handleEventsDay(w http.ResponseWriter, r *http.Request) {
	// evn := event.Event{}
	if r.Method != http.MethodGet {
		err := fmt.Sprintf(`"the Get method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	dt := r.URL.Query().Get("data")
	tm, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	res := []event.Event{}
	for _, k := range s.data {
		tmp, err := time.Parse(time.DateOnly, k.Time)
		if err != nil {
			s.Logger.Error(err.Error())
		}
		if tmp.Day() == tm.Day() {
			res = append(res, k)
		}
	}
	rs, err := json.Marshal(res)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(rs), http.StatusOK)
}
func (s *Server) handleEventsWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Sprintf(`"the Get method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
}
func (s *Server) handleEventsMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Sprintf(`"the Get method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	dt := r.URL.Query().Get("data")
	tm, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	res := []event.Event{}
	for _, k := range s.data {
		tmp, err := time.Parse(time.DateOnly, k.Time)
		if err != nil {
			s.Logger.Error(err.Error())
		}
		if int(tmp.Month()) == int(tm.Month()) {
			res = append(res, k)
		}
	}
	rs, err := json.Marshal(res)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(rs), http.StatusOK)
}
