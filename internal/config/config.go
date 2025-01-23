package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Config struct {
	//Version string `yaml:"version"`
	//Env     string `yaml:"env"`
	//Debug   bool   `yaml:"debug"`
	//Logger  Logger `yaml:"logger"`
	JWT  JWT  `yaml:"jwt"`
	DB   DB   `yaml:"db"`
	HTTP HTTP `yaml:"http"`
	//GRPC    GRPC   `yaml:"grpc"`
}

func New() (*Config, error) {
	data, err := preprocess()
	if err != nil {
		return &Config{}, fmt.Errorf("couldn't process config file: %s", err.Error())
	}
	var cfg Config

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return &Config{}, fmt.Errorf("couldn't unmarshall config file: %s", err.Error())
	}
	cfg.DB.ParseURL()
	return &cfg, nil
}
