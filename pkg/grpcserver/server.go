package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	_defaultShutdownTimeout = 5 * time.Second
	_defaultListenerHost    = "0.0.0.0"
	_defaultListenerPort    = "9090"
)

type GRPCService interface {
	RegisterGRPC(*grpc.Server)
}

type Server struct {
	server          *grpc.Server
	uInterceptors   []grpc.UnaryServerInterceptor
	listenerHost    string
	listenerPort    string
	notify          chan error
	shutdownTimeout time.Duration
}

func New(service GRPCService, opts ...Option) *Server {
	srv := &Server{
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		listenerHost:    _defaultListenerHost,
		listenerPort:    _defaultListenerPort,
	}

	// Custom options
	for _, opt := range opts {
		opt(srv)
	}

	srv.uInterceptors = append(srv.uInterceptors, unaryRecoveryInterceptor)

	var serverOpts []grpc.ServerOption

	serverOpts = append(serverOpts, grpc.ChainUnaryInterceptor(srv.uInterceptors...))

	srv.server = grpc.NewServer(serverOpts...)

	service.RegisterGRPC(srv.server)

	srv.start()

	return srv
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	stopped := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	t := time.NewTimer(s.shutdownTimeout)

	select {
	case <-t.C:
		s.server.Stop()
		return errors.New("graceful shutdown timeout exceeded")
	case <-stopped:
		t.Stop()
	}

	return nil
}

func (s *Server) start() {
	go func() {
		defer close(s.notify)

		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.listenerHost, s.listenerPort))
		if err != nil {
			s.notify <- err
			return
		}

		s.notify <- s.server.Serve(lis)
	}()
}

func unaryRecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		panicInfo := recover()
		if panicInfo != nil {
			err = fmt.Errorf("PANIC RECOVER %+v", panicInfo)
			log.Printf("PANIC RECOVER FROM method: %q, req: %v, res_err: %v\n", info.FullMethod, req, err)
		}
	}()

	result, err := handler(ctx, req)

	return result, err
}
