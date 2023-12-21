package kraken_private_messages

import (
	"github.com/shulst/go-kraken-safe-trader/kraken-private-messages/config"
	"testing"
)

func TestKrakenOwnTrades(t *testing.T) {
	api := API{}
	cfg := config.FromEnv("../.env")
	err, token := api.getPrivateToken(cfg.ApiKey, cfg.ApiSecret)
	if err != nil {
		t.Fatal(err)
	}
	err = api.connect()
	if err != nil {
		t.Fatal(err)
	}
	err = api.subscribeToOwnTrades(token)
	if err != nil {
		t.Fatal(err)
	}
	api.listen()
}
