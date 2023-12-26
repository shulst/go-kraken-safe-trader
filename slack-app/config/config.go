package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string `env:"ENV_SLACK_BOT_TOKEN,required"`
	UserID   string `env:"ENV_SLACK_USER_ID,required"`
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
