package main

import (
	"context"
	"gitlab.com/xamops/auth/internal/application"
	"gitlab.com/xamops/auth/internal/config"
	"log"
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
	// todo @ do nothing
	app.Run()
}
