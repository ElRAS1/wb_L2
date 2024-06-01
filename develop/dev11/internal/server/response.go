package server

import (
	"fmt"
	"net/http"
)

func (s *Server) responseError(err string) []byte {
	return []byte(fmt.Sprintf(`{"error": %s}`, err))
}

func (s *Server) responseResult(txt []byte) []byte {
	return []byte(fmt.Sprintf(`{"result": %s}`, txt))
}

func (s *Server) response(w http.ResponseWriter, res []byte, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err := w.Write(res)

	if err != nil {
		s.Logger.Error(err.Error())
	}
}
