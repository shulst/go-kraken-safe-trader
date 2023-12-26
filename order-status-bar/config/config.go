package config

import (
	"image/color"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// ENV_BUY_BG_COLOR="80;160;80;1"
// ENV_BUY_FG_COLOR="0;160;0;1"
// ENV_SELL_BG_COLOR="160;80;80;1"
// ENV_SELL_FG_COLOR="160;0;0;1"
// ENV_OVERLAY_COLOR="0;0;0;0.34"
type EnvColor string

func (envColor EnvColor) RGBA() color.RGBA {
	strColors := strings.Split(string(envColor), ";")
	return color.RGBA{
		R: str2Uint8(strColors[0]),
		G: str2Uint8(strColors[1]),
		B: str2Uint8(strColors[2]),
		A: str2Uint8(strColors[3]),
	}
}
func str2Uint8(str string) uint8 {
	uint64Value, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		log.Fatal(err)
	}

	return uint8(uint64Value)
}

type Config struct {
	BuyBgColor   EnvColor `env:"ENV_BUY_BG_COLOR,required"`
	BuyFgColor   EnvColor `env:"ENV_BUY_FG_COLOR,required"`
	SellBgColor  EnvColor `env:"ENV_SELL_BG_COLOR,required"`
	SellFgColor  EnvColor `env:"ENV_SELL_FG_COLOR,required"`
	OverlayColor EnvColor `env:"ENV_OVERLAY_COLOR,required"`
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
