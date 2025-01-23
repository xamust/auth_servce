package application

import (
	"context"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gitlab.com/xamops/auth/internal/mappers"
	"gitlab.com/xamops/auth/internal/repositories"
	"gitlab.com/xamops/auth/internal/usecases"
	"gitlab.com/xamops/auth/pkg/auth"
	"gitlab.com/xamops/auth/pkg/db/postgres"
	"gitlab.com/xamops/auth/pkg/grpcserver"
	"gitlab.com/xamops/auth/pkg/httpserver"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	ctx        context.Context
	config     *config.Config
	logger     *slog.Logger
	db         *gorm.DB
	shutdown   chan os.Signal
	httpServer interfaces.HTTPServer
	grpcServer interfaces.GRPCServer
}

func NewWithContext(ctx context.Context, cfg *config.Config) (*Application, error) {
	gormDB, err := postgres.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}

	app := &Application{
		ctx:        ctx,
		config:     cfg,
		logger:     slog.Default(),
		db:         gormDB,
		shutdown:   make(chan os.Signal, 1),
		httpServer: defaultHTTPServer{},
		grpcServer: defaultGRPCServer{},
	}

	signal.Notify(app.shutdown, os.Interrupt, syscall.SIGTERM)
	return app, nil
}

func (a *Application) Run() {
	const op = "application.Run"
	_, cancel := context.WithCancel(a.ctx)

	a.logger.Info("Application started ...")

	select {
	case err := <-a.httpServer.Notify():
		a.logger.Error(op, "http server error", err.Error())
	case err := <-a.grpcServer.Notify():
		a.logger.Error(op, "grpc server error", err.Error())
	case s := <-a.shutdown:
		a.logger.Info(op, "Signal received", s)
	}

	// Graceful shutdown
	cancel()

	a.stop()
}

func (a *Application) Logger() *slog.Logger {
	return a.logger
}

func (a *Application) stop() {
	const op = "application.stop"
	close(a.shutdown)

	if err := a.httpServer.Shutdown(); err != nil {
		a.logger.Error(op, "http server shutdown", err.Error())
	}
	if err := a.grpcServer.Shutdown(); err != nil {
		a.logger.Error(op, "grpc server shutdown", err.Error())
	}

	a.logger.Info(op, "Info", "Application stopped")
}

func (a *Application) RegisterHTTPServer(httpServer *httpserver.Server) {
	a.httpServer = httpServer
}

func (a *Application) RegisterGRPCServer(grpcServer *grpcserver.Server) {
	a.grpcServer = grpcServer
}

func (a *Application) InitUsecases() (interfaces.Usecases, error) {
	return usecases.New(usecases.Dependencies{
		Config:       a.config,
		Repositories: repositories.New(a.db, a.logger),
		Mappers:      mappers.New(),
		Logger:       a.logger,
	}), nil
}

func (a *Application) InitAccessToken() auth.TokenHandler {
	return auth.NewJWTHandler(&auth.TokenConfig{
		Secret: a.config.JWT.AccessSecret,
		TTL:    a.config.JWT.AccessTokenTTL,
	})
}

func (a *Application) InitRefreshToken() auth.TokenHandler {
	return auth.NewJWTHandler(&auth.TokenConfig{
		Secret: a.config.JWT.RefreshSecret,
		TTL:    a.config.JWT.RefreshTokenTTL,
	})
}
