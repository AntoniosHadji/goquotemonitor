package sfox

import (
	"log"
	"testing"
)

func TestMarket(t *testing.T) {
	response, err := GetBook("BTCUSD")
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}

	if response.Bids == nil {
		t.Fatalf(`No Bids in response`)
	}

	log.Println(response.Bids)

	if response.Asks == nil {
		t.Fatalf(`No Asks in response`)
	}

	log.Println(response.Asks)

	bid := GetPriceForSize(response.Bids, 1.0)

	log.Println(bid)

	ask := GetPriceForSize(response.Asks, 1.0)

	log.Println(ask)

	bid, ask, spread := CalcSpread(response, 1.0)

	log.Println(bid, ask, spread)
}

func TestRFQ(t *testing.T) {
	response, err := GetRFQ("BTCUSD", "BUY", 1.0)
	if err != nil {
		t.Fatalf("error:%v", err)
	}
	log.Printf("%#v", response)
}
