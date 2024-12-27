package config

import (
	"fmt"
	"time"
)

type DB struct {
	Url                  string        `yaml:"url"`
	Type                 string        `yaml:"db_type"`
	User                 string        `yaml:"db_user"`
	Password             string        `yaml:"db_password"`
	Host                 string        `yaml:"db_host"`
	Port                 string        `yaml:"db_port"`
	DB                   string        `yaml:"db_db"`
	SSL                  string        `yaml:"db_ssl_mode"`
	ConnectionTimeout    time.Duration `yaml:"db_connection_timeout"`
	MaxConnectionRetries int           `yaml:"db_max_connection_retries"`
	Debug                bool          `yaml:"db_debug"`
	MaxIdle              int           `yaml:"db_max_idle"`
	MaxOpen              int           `yaml:"db_max_open"`
}

func (d *DB) ParseURL() {
	d.Url = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%v",
		d.Type,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DB,
		d.SSL)
}
