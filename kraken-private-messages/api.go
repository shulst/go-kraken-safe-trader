package kraken_private_messages

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type API struct {
	client                *websocket.Conn
	ownTradesChannel      chan trades
	openOrdersChannel     chan orders
	ownTradesMessageLock  sync.RWMutex
	openOrdersMessageLock sync.RWMutex
}

type token string
type Nonce string

func (n Nonce) New() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

type KrakenTokenResult struct {
	Token   token `json:"token"`
	Expires int   `json:"expires"`
}

type KrakenTokenResponse struct {
	Error  []string          `json:"error"`
	Result KrakenTokenResult `json:"result"`
}
