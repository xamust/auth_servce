package application

import (
	"context"
	"gitlab.com/xamops/auth/internal/config"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type server interface {
	Notify() <-chan error
	Shutdown() error
}

type Application struct {
	ctx        context.Context
	config     *config.Config
	logger     *slog.Logger
	db         *gorm.DB
	shutdown   chan os.Signal
	httpServer server
	grpcServer server
}

func NewWithContext(ctx context.Context, cfg *config.Config) (*Application, error) {

}

func (a *Application) Run() error {

	return nil
}
