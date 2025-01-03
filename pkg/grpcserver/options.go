package grpcserver

import (
	"time"
)

// Option -.
type Option func(*Server)

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// Host -.
func Host(host string) Option {
	return func(s *Server) {
		s.listenerHost = host
	}
}

// Port -.
func Port(port string) Option {
	return func(s *Server) {
		s.listenerPort = port
	}
}
