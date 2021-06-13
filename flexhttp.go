package flexhttp

import (
	"context"
	"log"
	"net/http"
)

type Server struct{ *http.Server }

func NewHTTPServer(s *http.Server) *Server {
	return &Server{Server: s}
}

func (s *Server) Run(_ context.Context) error {
	log.Printf("serving on: http://localhost%s\n", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) Halt(ctx context.Context) error {
	return s.Shutdown(ctx)
}
