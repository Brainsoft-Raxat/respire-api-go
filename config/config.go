package config

import (
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

type Configs struct {
	App         *AppConfig   `mapstructure:"app" json:"app"`
	Firebase    *Firebase    `mapstructure:"firebase" json:"firebase"`
	AIAssistant *AIAssistant `mapstructure:"ai_assistant" json:"ai_assistant"`
}

type AppConfig struct {
	Timeout time.Duration `mapstructure:"timeout" json:"timeout" default:"60s"`
	Port    string        `mapstructure:"port" json:"port"`
	Host    string        `mapstructure:"host" json:"host"`
	Env     string        `mapstructure:"env" json:"env"`
}

type Firebase struct {
	ProjectID string `mapstructure:"project_id" json:"project_id"`
}

type AIAssistant struct {
	BaseURL string `mapstructure:"base_url" json:"base_url"`
}

func New() (*Configs, error) {
	configFile := "config/config.yaml"

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

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
