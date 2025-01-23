package httpserver

import (
	"net"
	"time"
)

// Option -.
type Option func(*Server)

// WithPort set http server port
func WithPort(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// WithReadTimeout set read timeout
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WithWriteTimeout set write timeout
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// WithShutdownTimeout set shutdown timeout
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
