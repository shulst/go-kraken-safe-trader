package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey    string `env:"ENV_KRAKEN_API_KEY,required"`
	ApiSecret string `env:"ENV_KRAKEN_API_SECRET,required"`
}

var (
	readConfigOnce sync.Once
	cfg            *Config
)

func FromEnv(filenames ...string) *Config {
	readConfigOnce.Do(func() {
		godotenv.Load(filenames...)
		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
