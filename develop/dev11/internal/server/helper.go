package server

import (
	"fmt"
	"net/http"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
)

func (s *Server) check(w http.ResponseWriter, r *http.Request, evn event.Event) error {
	if r.Method != http.MethodPost {
		err := fmt.Sprintf(`"the Post method was expected and not %s"`, r.Method)
		s.response(w, s.responseError(err), http.StatusMethodNotAllowed)
		return fmt.Errorf(err)
	}
	if err := evn.Decode(*r); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusBadRequest)
		return err
	}
	if err := evn.Validate(); err != nil {
		s.Logger.Error(err.Error())
		s.response(w, s.responseError(err.Error()), http.StatusServiceUnavailable)
		return err
	}
	return nil
}
