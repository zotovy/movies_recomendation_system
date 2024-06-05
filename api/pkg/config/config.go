package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"log"
	"os"
)

type Config struct {
	LogLevel    string `yaml:"logLevel" env-required:"true"`
	Environment string `yaml:"environment" env-required:"true"`

	Api struct {
		IP   string `yaml:"ip" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
		Host string `yaml:"host" env-required:"true"`
	} `yaml:"api" env-required:"true"`

	PostgreSQL struct {
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Host     string `yaml:"host" env-required:"true"`
		Port     int64  `yaml:"port" env-required:"true"`
		Database string `yaml:"database" env-required:"true"`
	} `yaml:"postgres"`

	Redis struct {
		Address  string `yaml:"address" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Database int    `yaml:"database"`
	} `yaml:"redis" env-required:"true"`

	Auth struct {
		Access struct {
			Lifetime int64  `yaml:"lifetime" env-required:"true"`
			Secret   string `yaml:"secret" env-required:"true"`
		} `yaml:"access" env-required:"true"`
		Refresh struct {
			Lifetime int64  `yaml:"lifetime" env-required:"true"`
			Secret   string `yaml:"secret" env-required:"true"`
		} `yaml:"refresh"`
	} `yaml:"auth" env-required:"true"`

	Services struct {
		RecommendationsAPI struct {
			BaseUrl string `yaml:"base-url" env-required:"true"`
		} `yaml:"recommendations-api"`
	} `yaml:"services"`
}

func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatal("error opening config file", zap.Error(err))
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal("error reading config file", zap.Error(err))
	}

	return &cfg
}
