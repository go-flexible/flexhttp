package flexhttp

import (
	"context"
	"net"
	"net/http"
	"time"
)

// Default timeout values for http servers.
const (
	DefaultIdleTimeout       = 5 * time.Second
	DefaultWriteTimeout      = 5 * time.Second
	DefaultReadTimeout       = 5 * time.Second
	DefaultReadHeaderTimeout = 1 * time.Second
)

// DefaultHTTPServer provides a default http server.
var DefaultHTTPServer = &http.Server{
	WriteTimeout:      DefaultWriteTimeout,
	ReadTimeout:       DefaultReadTimeout,
	IdleTimeout:       DefaultIdleTimeout,
	ReadHeaderTimeout: DefaultReadHeaderTimeout,
}

// Logger defines any logger able to call Printf.
type Logger interface {
	Printf(format string, v ...interface{})
}

// Option is a type of func that allows you change defaults of the *Server
type Option func(s *Server)

// WithLogger allows you to set a logger for the server.
func WithLogger(logger Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// Server defines the flexhttp Server.
type Server struct {
	logger Logger
	*http.Server
}

// New returns a new flexhttp server, using the provided http.Server.
// NOTE: If no values are provided, defaults will be used the following fields:
//  - ReadTimeout
//  - ReadHeaderTimeout
//  - WriteTimeout
//  - IdleTimeout
func New(httpServer *http.Server, options ...Option) *Server {
	if httpServer == nil {
		httpServer = DefaultHTTPServer
	}
	if httpServer.ReadTimeout == 0 {
		httpServer.ReadTimeout = DefaultHTTPServer.ReadTimeout
	}
	if httpServer.ReadHeaderTimeout == 0 {
		httpServer.ReadHeaderTimeout = DefaultHTTPServer.ReadHeaderTimeout
	}
	if httpServer.WriteTimeout == 0 {
		httpServer.WriteTimeout = DefaultHTTPServer.WriteTimeout
	}
	if httpServer.IdleTimeout == 0 {
		httpServer.IdleTimeout = DefaultHTTPServer.IdleTimeout
	}

	server := &Server{Server: httpServer}

	for _, opt := range options {
		opt(server)
	}

	return server
}

// Run satisfies the flex Runner interface.
func (s *Server) Run(_ context.Context) error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	if address, ok := listener.Addr().(*net.TCPAddr); ok {
		s.logger.Printf("serving on http://%s", address)
	}
	return s.Serve(listener)
}

// Halt satisfies the flex Halter interface.
func (s *Server) Halt(ctx context.Context) error {
	s.logger.Printf("shutting down http server...")
	return s.Shutdown(ctx)
}
