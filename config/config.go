package config

import (
	"time"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

type Configs struct {
	App      *AppConfig `json:"app"`
	Firebase *Firebase  `json:"firebase"`
}

type AppConfig struct {
	Timeout time.Duration `json:"timeout" default:"60s"`
	Port    string        `json:"port"`
	Host    string        `json:"host"`
}

type Firebase struct {
	ProjectID string `json:"project_id"`
}

func New() (*Configs, error) {
	configFile := "config/config.yaml"

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Configs{}

	if err := defaults.Set(cfg); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
