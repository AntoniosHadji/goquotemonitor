package coinbase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var cbBaseURL = "https://api.exchange.coinbase.com"
var client = &http.Client{}

// CoinbaseBookResponse ...
type CoinbaseBookResponse struct {
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
	Sequence int64           `json:"sequence"`
}

func getBook(ticker string) (*CoinbaseBookResponse, error) {
	path := fmt.Sprintf("%s/products/%s/book?level=2", cbBaseURL, ticker)
	req, err := http.NewRequest("GET", path, nil)
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
	defer res.Body.Close()

	var r CoinbaseBookResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode == 200 {
		log.Printf("Status: %s for ticker: %s", res.Status, ticker)
	} else {
		log.Printf("Status: %s for ticker: %s", res.Status, ticker)
		log.Println(res)
		return nil, fmt.Errorf("Bad status: %s", res.Status)
	}

	return &r, nil
}

func calcSpread(data *CoinbaseBookResponse, size float64) (float64, float64, float64) {

	bid := getPriceForSize(data.Bids, size)
	ask := getPriceForSize(data.Asks, size)
	// debug
	//fmt.Println(bid, ask)

	return bid, ask, 10000 * (ask - bid) / bid

}

func getPriceForSize(data [][]interface{}, size float64) float64 {

	// fmt.Println(len(data), ",psum", ",ssum", ",psum/ssum", ",p", ",s")
	var p, s, psum, ssum, z float64
	var err error
	for i := range data {
		if s, err = strconv.ParseFloat(fmt.Sprintf("%v", data[i][1]), 64); err == nil {
			// fmt.Printf("%T: %f -> ", s, s)
			ssum += s
		}
		if p, err = strconv.ParseFloat(fmt.Sprintf("%v", data[i][0]), 64); err == nil {
			if ssum >= size {
				z = size - (ssum - s)
				ssum = ssum - s + z
				psum += p * z
				break
			}
			psum += p * s
		}
		// debug
		// fmt.Printf("%d, %.8f,%.8f,%.8f,%.8f,%.8f\n", i, psum, ssum, psum/ssum, p, s)
	}
	// fmt.Printf("99, %.8f,%.8f,%.8f,%.8f,%.8f\n", psum, ssum, psum/ssum, p, z)

	return psum / ssum
}
