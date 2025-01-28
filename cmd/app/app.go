package main

// @title           Swagger xamops auth API
// @version         1.0.0
// @description     This is a sample server celler server.

import (
	"context"
	"gitlab.com/xamops/auth/internal/application"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/controller/http"
	"gitlab.com/xamops/auth/internal/controller/http/v1/handlers"
	"gitlab.com/xamops/auth/internal/mappers"
	"gitlab.com/xamops/auth/pkg/httpserver"
	"log"
	"log/slog"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error parse config file: %v", err)
		return
	}
	ctx := context.Background()
	app, err := application.NewWithContext(ctx, cfg)
	if err != nil {
		log.Fatalf("error init application: %s", err)
	}

	usecases, err := app.InitUsecases()
	if err != nil {
		log.Fatalf("error init usecases: %s", err)
	}

	access := app.InitAccessToken()
	refresh := app.InitRefreshToken()

	deps := handlers.Dependencies{
		Config:         cfg,
		Logger:         slog.Default(),
		Usecases:       usecases,
		Mappers:        mappers.New(),
		AccessHandler:  access,
		RefreshHandler: refresh,
	}

	httpServer := httpserver.New(http.NewRouter(deps),
		httpserver.WithPort(cfg.HTTP.Port),
	)

	app.RegisterHTTPServer(httpServer)

	app.Run()
}
