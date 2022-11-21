package server

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App          `yaml:"app"`
		GRPC         `yaml:"grpc"`
		FilesStorage `yaml:"files_storage"`
		Log          `yaml:"logger"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// FRPC -.
	GRPC struct {
		Port string `yaml:"port" env:"GRPC_PORT"`
	}

	FilesStorage struct {
		Location string `yaml:"location" env:"FILES_LOCATION"`
	}
	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/server/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadConfig(".env", cfg); err != nil {
		log.Println(err.Error())
		return cfg, nil
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
