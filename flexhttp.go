package flexhttp

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// Default timeout values for http servers.
const (
	DefaultIdleTimeout       = 5 * time.Second
	DefaultWriteTimeout      = 5 * time.Second
	DefaultReadTimeout       = 5 * time.Second
	DefaultReadHeaderTimeout = 1 * time.Second
)

var (
	// DefaultHTTPServer provides a default http server.
	DefaultHTTPServer = &http.Server{
		WriteTimeout:      DefaultWriteTimeout,
		ReadTimeout:       DefaultReadTimeout,
		IdleTimeout:       DefaultIdleTimeout,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
	}

	// logger defines a logger with a prefix.
	logger = log.New(os.Stderr, "flexhttp: ", 0)
)

// Server defines the flexhttp Server.
type Server struct{ *http.Server }

// New returns a new flexhttp server, using the provided http.Server.
// NOTE: If no values are provided, defaults will be used the following fields:
//  - ReadTimeout
//  - ReadHeaderTimeout
//  - WriteTimeout
//  - IdleTimeout
func New(server *http.Server) *Server {
	if server == nil {
		server = DefaultHTTPServer
	}
	if server.ReadTimeout == 0 {
		server.ReadTimeout = DefaultHTTPServer.ReadTimeout
	}
	if server.ReadHeaderTimeout == 0 {
		server.ReadHeaderTimeout = DefaultHTTPServer.ReadHeaderTimeout
	}
	if server.WriteTimeout == 0 {
		server.WriteTimeout = DefaultHTTPServer.WriteTimeout
	}
	if server.IdleTimeout == 0 {
		server.IdleTimeout = DefaultHTTPServer.IdleTimeout
	}
	return &Server{Server: server}
}

// Run satisfies the flex Runner interface.
func (s *Server) Run(_ context.Context) error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	if address, ok := listener.Addr().(*net.TCPAddr); ok {
		logger.Printf("serving on http://%s", address)
	}
	return s.Serve(listener)
}

// Halt satisfies the flex Halter interface.
func (s *Server) Halt(ctx context.Context) error {
	logger.Printf("shutting down http server...")
	return s.Shutdown(ctx)
}
