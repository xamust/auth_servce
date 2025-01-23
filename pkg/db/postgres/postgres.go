package postgres

import (
	"fmt"
	"gitlab.com/xamops/auth/internal/config"
	"gorm.io/gorm"
	"log"
	"time"

	_ "github.com/lib/pq"
	postgresdriver "gorm.io/driver/postgres"
	gormlogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DB, cfg.DB.SSL)

	loglevel := gormlogger.Silent
	if cfg.DB.Debug {
		loglevel = gormlogger.Info
	}

	newLogger := gormlogger.New(
		log.Default(),
		gormlogger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      loglevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)

	conn, err := gorm.Open(postgresdriver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create a new postgres (v2) connection: %s", err.Error())
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %s", err.Error())
	}

	maxIdleConn := 1
	if cfg.DB.MaxIdle > 0 {
		maxIdleConn = cfg.DB.MaxIdle
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)

	maxOpenConn := 1
	if cfg.DB.MaxOpen > 0 {
		maxOpenConn = cfg.DB.MaxOpen
	}
	sqlDB.SetMaxOpenConns(maxOpenConn)

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping to database: %s", err.Error())
	}

	db = conn

	return db, nil
}
