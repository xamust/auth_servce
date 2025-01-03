package handlers

import (
	"gitlab.com/xamops/auth/internal/config"
	"log/slog"
)

type Dependencies struct {
	Config *config.Config
	Logger *slog.Logger
}

type Handlers struct {
}

func NewRoutes(deps Dependencies) Handlers {
	return Handlers{}
}
