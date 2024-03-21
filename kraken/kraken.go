package kraken

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	baseURL = "https://api.kraken.com/0"
	client  = &http.Client{}
)

// OrderBookResponse ...
type OrderBookResponse struct {
	Error  interface{} `json:"error"`
  Result map[string]Prices `json:"result"`
}

type Prices struct{
  Bids [][]interface{} `json:"bids"`
  Asks [][]interface{} `json:"asks"`
}

func getBook(ticker string) (*OrderBookResponse, error) {
	path := fmt.Sprintf("%s/public/Depth?pair=%s", baseURL, ticker)
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

	var r OrderBookResponse
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

func calcSpread(data *OrderBookResponse, size float64) (float64, float64, float64) {
  var bid, ask float64
    for k := range data.Result {
     bid = getPriceForSize(data.Result[k].Bids, size)
     ask = getPriceForSize(data.Result[k].Asks, size)
  }
	// debug
	// fmt.Println(bid, ask)

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
