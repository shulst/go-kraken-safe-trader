package kraken_private_messages

import (
	"github.com/shulst/go-kraken-safe-trader/kraken-private-messages/config"
	"log"
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

	ownTrades := api.ownTradesChannel
	go listenOwnTrades(ownTrades)

	api.listen()
}

func listenOwnTrades(ch chan trades) {
	for ownTrades := range ch {
		log.Printf("%v\n", ownTrades)
	}
}
