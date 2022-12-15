package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var cbBaseURL = "https://api.exchange.coinbase.com"

// CoinbaseBookResponse ...
type CoinbaseBookResponse struct {
	Bids        [][]interface{} `json:"bids"`
	Asks        [][]interface{} `json:"asks"`
	Sequence    int64           `json:"sequence"`
	AuctionMode bool            `json:"auction_mode"`
	Auction     interface{}     `json:"auction"`
}

func getBook(ticker string) (*CoinbaseBookResponse, error) {
	path := fmt.Sprintf("%s/products/%s-USD/book", cbBaseURL, ticker)
	// query = {"level": 2}
	//     r = requests.get(f"{BASE_PRO}{path}", params=query).json()
	req, err := http.NewRequest("GET", path+"?level=2", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var r CoinbaseBookResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode == 200 {
		log.Printf("Status: %s\n", res.Status)
	} else {
		fmt.Println(res.Status)
		fmt.Println(res)
		fmt.Printf("%#v", r)
		return nil, fmt.Errorf("Bad status: %s\n%#v", res.Status, r)
	}

	return &r, nil
}
