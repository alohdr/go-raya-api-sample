package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Port       string `envconfig:"PORT" default:"8080"`
		DbName     string `envconfig:"DB_NAME" default:"bank"`
		DbHost     string `envconfig:"DB_HOST" default:"localhost"`
		DbPort     int    `envconfig:"DB_PORT" default:"5432"`
		DbUsername string `envconfig:"DB_USER" default:"postgres"`
		DbPassword string `envconfig:"DB_PASS" default:"password"`
	}
)

var cfg *Config

func Init() {
	cfg = new(Config)
	envconfig.MustProcess("", cfg)
}

func Get() *Config {
	return cfg
}
