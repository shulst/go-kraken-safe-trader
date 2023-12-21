package kraken_private_messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type unixTime string

func (u unixTime) Time() time.Time {
	i, _ := strconv.ParseFloat(string(u), 64)
	return time.Unix(int64(i), 0)
}

type price string
type vol string
type margin string

type trade struct {
	PostXId   string   `json:"postxid"`
	OrderTXId string   `json:"ordertxid"`
	OrderType string   `json:"ordertype"`
	UTime     unixTime `json:"time"`
	Pair      string   `json:"pair"`
	Type      string   `json:"type"`
	Price     price    `json:"price"`
	Cost      price    `json:"cost"`
	Fee       price    `json:"fee"`
	Vol       vol      `json:"vol"`
	Margin    margin   `json:"margin"`
}
type trades map[string]trade

func (ts trades) Combine(cTss ...trades) (combined trades, err error) {
	combined = make(trades)
	for key, ct := range ts {
		combined[key] = ct
	}
	for _, cts := range cTss {
		for key, ct := range cts {
			if _, ok := combined[key]; !ok {
				combined[key] = ct
			} else {
				return nil, errors.New(fmt.Sprintf("duplicate key %s", key))
			}
		}
	}
	return
}

func (api *API) subscribeToOwnTrades(token token) error {
	payload := fmt.Sprintf("{\"event\": \"subscribe\", \"subscription\": {\"name\": \"ownTrades\", \"token\": \"%s\"}}", token)
	err := api.client.WriteMessage(1, []byte(payload))
	if err != nil {
		log.Printf("Error subscribing: %s", err)
		return err
	}
	return nil
}

func (api *API) handleOwnTradeMessage(msg []byte) error {
	received := time.Now()
	fmt.Printf("%v %s \n", received, msg)

	var resp []interface{}
	if err := json.Unmarshal(msg, &resp); err != nil {
		return err
	}
	if len(resp) < 3 || resp[1] != "ownTrades" {
		return errors.New("not ownTrades")
	}
	api.ownTradesMessageLock.Lock()
	defer api.ownTradesMessageLock.Unlock()
	res, err := json.Marshal(resp[0])
	if err != nil {
		return err
	}
	var ts []trades
	if err := json.Unmarshal(res, &ts); err != nil {
		return err
	}
	allTs, err := ts[0].Combine(ts[1:]...)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", allTs)
	return nil
}
