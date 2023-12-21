package kraken_private_messages

import (
	"fmt"
	"github.com/shulst/go-kraken-safe-trader/kraken-private-messages/config"
	"testing"
)

func TestKrakenPrivateToken(t *testing.T) {
	api := API{}
	cfg := config.FromEnv("../.env")
	err, token := api.getPrivateToken(cfg.ApiKey, cfg.ApiSecret)
	fmt.Printf("Error: %v\nToken: %s\n", err, token)
}
