package server

import "fmt"

func (s *Server) responseError(err error) string {
	return fmt.Sprintf("{error: %v}", err)
}

func (s *Server) responseResult(txt string) string {
	return fmt.Sprintf("{result: %s}", txt)
}
