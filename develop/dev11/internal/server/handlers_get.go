package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
)

func (s *Server) handleEventsDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Sprintf(`"the Get method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	dt := r.URL.Query().Get("data")
	tm, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	res := []event.Event{}
	for _, k := range s.data {
		tmp, err := time.Parse(time.DateOnly, k.Time)
		if err != nil {
			s.Logger.Error(err.Error())
			continue
		}
		if (tmp.Day() == tm.Day()) && (tmp.Year() == tm.Year()) && (tmp.Month() == tm.Month()) {
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	dt := r.URL.Query().Get("data")
	tm, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	res := []event.Event{}
	for _, evn := range s.data {
		tmp, err := time.Parse(time.DateOnly, evn.Time)
		if err != nil {
			s.Logger.Error(err.Error())
			continue
		}
		subTime := int(tm.AddDate(0, 0, 7).Sub(time.Time(tmp)).Hours()) / 24
		if subTime >= 0 && subTime <= 7 {
			res = append(res, evn)
		}
	}
	rs, err := json.Marshal(res)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(rs), http.StatusOK)
}
func (s *Server) handleEventsMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Sprintf(`"the Get method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	dt := r.URL.Query().Get("data")
	tm, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return
	}
	res := []event.Event{}
	for _, k := range s.data {
		tmp, err := time.Parse(time.DateOnly, k.Time)
		if err != nil {
			s.Logger.Error(err.Error())
			continue
		}
		if int(tmp.Month()) == int(tm.Month()) && (tmp.Year() == tm.Year()) {
			res = append(res, k)
		}
	}
	rs, err := json.Marshal(res)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	s.response(w, s.responseResult(rs), http.StatusOK)
}
