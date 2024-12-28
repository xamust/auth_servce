package main

import (
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/pkg/db/postgres"
	migrator "gitlab.com/xamops/auth/pkg/migrator/goose"
	"log"
	"os"
)

/*
go run migrator status
go run migrator create filename sql
go run migrator up
go run migrator down
*/

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error parse config file: %v", err)
		return
	}
	db, err := postgres.Database(cfg).DB()
	if err != nil {
		panic(err)
	}
	migr := migrator.NewGooseMigrator(db)
	if err := migr.Commands(os.Args[1], os.Args[2:]...); err != nil {
		panic(err)
	}
}
