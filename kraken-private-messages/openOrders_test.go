package kraken_private_messages

import (
	"encoding/json"
	"github.com/shulst/go-kraken-safe-trader/kraken-private-messages/config"
	"log"
	"testing"
)

func TestKrakenOpenOrders(t *testing.T) {
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
	err = api.subscribeToOpenOrders(token)
	if err != nil {
		t.Fatal(err)
	}

	openOrdersChan := api.openOrdersChannel
	go listenOrders(openOrdersChan)

	api.listen()
}

func listenOrders(ch chan orders) {
	for openOrders := range ch {
		msg, _ := json.MarshalIndent(openOrders, "", "  ")
		log.Printf("%s\n", msg)
	}
}
