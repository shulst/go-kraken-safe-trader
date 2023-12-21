package kraken_private_messages

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	REST_URI      = "https://api.kraken.com"
	REST_URI_PATH = "/0/private/GetWebSocketsToken"
	WS_URI        = "wss://ws-auth.kraken.com"
)

func (api *API) connect() error {
	log.Println("Connecting and praying ...")
	c, _, err := websocket.DefaultDialer.Dial(WS_URI, http.Header{})
	if err != nil {
		return err
	}
	api.client = c
	return nil
}

func (api *API) listen() {
	for api.client != nil {
		_, msg, err := api.client.ReadMessage()
		if err != nil {
			log.Printf("Received fatal error from kraken: %s\n", err)
			api.client.Close()
			panic(err)
		}
		if err := api.handleOwnTradeMessage(msg); err != nil {
			log.Printf("Error handling message: %s\nError: %s\n", msg, err)
		}
		if err := api.handleOpenOrdersMessage(msg); err != nil {
			log.Printf("Error handling message: %s\nError: %s\n", msg, err)
		}
	}
}
