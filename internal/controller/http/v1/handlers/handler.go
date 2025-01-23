package handlers

import (
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gitlab.com/xamops/auth/pkg/auth"
	"log/slog"
)

const (
	accessToken  = "access_token"
	refreshToken = "refresh_token"
)

type Dependencies struct {
	Config         *config.Config
	Logger         *slog.Logger
	Usecases       interfaces.Usecases
	Mappers        interfaces.Mappers
	AccessHandler  auth.TokenHandler
	RefreshHandler auth.TokenHandler
}

type apiHandlersV1 struct {
	config   *config.Config
	logger   *slog.Logger
	usecases interfaces.Usecases
	mappers  interfaces.Mappers
	access   auth.TokenHandler
	refresh  auth.TokenHandler
}

func NewRoutes(deps Dependencies) interfaces.Handlers {
	return &apiHandlersV1{
		config:   deps.Config,
		logger:   deps.Logger,
		usecases: deps.Usecases,
		mappers:  deps.Mappers,
		access:   deps.AccessHandler,
		refresh:  deps.RefreshHandler,
	}
}
