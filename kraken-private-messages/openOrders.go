package kraken_private_messages

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Order example:
//
//	{
//	  "OI3VWZ-XVFKD-42NHKO": {
//		"avg_price": "0.00000",
//		"cost": "0.00000",
//		"descr": {
//		  "close": "",
//		  "leverage": "",
//		  "order": "sell 264.80262516 DYDX/EUR @ limit 3.66900",
//		  "ordertype": "limit",
//		  "pair": "DYDX/EUR",
//		  "price": "3.66900",
//		  "price2": "0.00000",
//		  "type": "sell"
//		},
//		"expiretm": "",
//		"fee": "0.00000",
//		"stopprice": "0.00000",
//		"limitprice": "0.00000",
//		"misc": "",
//		"oflags": "fciq",
//		"timeinforce": "GTC",
//		"cancel_reason": "",
//		"ratecount": ""
//	  }, ...
//	}

type OrderDesc struct {
	Close     price  `json:"close"`
	Leverage  string `json:"leverage"`
	Order     string `json:"order"`
	OrderType string `json:"ordertype"`
	Pair      pair   `json:"pair"`
	Price     price  `json:"price"`
	Price2    price  `json:"price2"`
	Type      string `json:"type"`
}

type order struct {
	AvgPrice     price        `json:"avg_price"`
	Cost         price        `json:"cost"`
	Descr        OrderDesc    `json:"descr"`
	ExpireUTime  unixTime     `json:"expiretm"`
	Fee          price        `json:"fee"`
	StopPrice    price        `json:"stopprice"`
	LimitPrice   price        `json:"limitprice"`
	Misc         string       `json:"misc"`
	Oflags       string       `json:"oflags"`
	TimeInforce  unixTime     `json:"timeinforce"`
	CancelReason cancelReason `json:"cancel_reason"`
	RateCount    string       `json:"ratecount"`
}
type orders map[string]order

func (os orders) Combine(oss ...orders) (combined orders, err error) {
	combined = make(orders)
	for key, o := range os {
		combined[key] = o
	}
	for _, oo := range oss {
		for key, o := range oo {
			combined[key] = o
		}
	}
	return
}

func (api *API) subscribeToOpenOrders(token token) error {
	api.openOrdersChannel = make(chan orders)
	payload := fmt.Sprintf("{\"event\": \"subscribe\", \"subscription\": {\"name\": \"openOrders\", \"token\": \"%s\"}}", token)
	err := api.client.WriteMessage(1, []byte(payload))
	if err != nil {
		log.Printf("Error subscribing: %s", err)
		return err
	}
	return nil
}

func (api *API) handleOpenOrdersMessage(msg []byte) error {
	received := time.Now()
	fmt.Printf("%v %s \n", received, msg)

	var resp []interface{}
	if err := json.Unmarshal(msg, &resp); err != nil {
		return nil
	}
	if len(resp) < 3 || resp[1] != "openOrders" {
		return nil
	}
	api.openOrdersMessageLock.Lock()
	defer api.openOrdersMessageLock.Unlock()
	res, err := json.Marshal(resp[0])
	if err != nil {
		return err
	}
	var os []orders
	if err := json.Unmarshal(res, &os); err != nil {
		return err
	}
	combined, err := os[0].Combine(os[1:]...)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", combined)
	api.openOrdersChannel <- combined
	return nil
}
