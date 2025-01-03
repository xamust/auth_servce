package application

import "dbfhub.gitlab.yandexcloud.net/plano-dev/backend/plano-auth.git/internal/interfaces"

type defaultHTTPServer struct{}

var _ interfaces.HTTPServer = defaultHTTPServer{}

// Notify implements interfaces.HTTPServer.
func (defaultHTTPServer) Notify() <-chan error {
	return nil
}

// Shutdown implements interfaces.HTTPServer.
func (defaultHTTPServer) Shutdown() error {
	return nil
}

type defaultGRPCServer struct{}

var _ interfaces.GRPCServer = defaultGRPCServer{}

// Notify implements interfaces.GRPCServer.
func (defaultGRPCServer) Notify() <-chan error {
	return nil
}

// Shutdown implements interfaces.GRPCServer.
func (defaultGRPCServer) Shutdown() error {
	return nil
}
