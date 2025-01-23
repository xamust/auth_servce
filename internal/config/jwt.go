package config

import "time"

type JWT struct {
	AccessSecret    string        `yaml:"access_secret"`
	RefreshSecret   string        `yaml:"refresh_secret"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
}
