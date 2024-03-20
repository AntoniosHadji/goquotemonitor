package sfox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// var baseURL = "https://api.sfox.com/v1"
var baseURL = "https://api.staging.sfox.com/v1"
var token string
var client = &http.Client{}

func init() {
	token = os.Getenv("SFOXAPI")
}

// BookResponse ...
type BookResponse struct {
	Bids          [][]interface{} `json:"bids"`
	Asks          [][]interface{} `json:"asks"`
	LastUpdated   int64           `json:"lastupdated"`
	LastPublished int64           `json:"lastpublished"`
}

type RFQRequest struct {
	Pair     string  `json:"pair"`
	Side     string  `json:"side"`
	Quantity float64 `json:"quantity"`
}

type RFQResponse struct {
	QuoteID     string    `json:"quote_id"`
	Pair        string    `json:"pair"`
	Side        string    `json:"side"`
	Date_expiry time.Time `json:"date_expiry"`
	Date_quote  time.Time `json:"date_quote"`
	Amount      float64   `json:"amount"`
	Quantity    float64   `json:"quantity"`
	Fees        float64   `json:"fees"`
	Buy_price   float64   `json:"buy_price,omitempty"`
	Sell_price  float64   `json:"sell_price,omitempty"`
	Error       string    `json:"error,omitempty"`
}

type APIError struct {
	Error string `json:"error"`
}

func GetBook(ticker string) (*BookResponse, error) {

	path := fmt.Sprintf("%s/markets/orderbook/%s", baseURL, ticker)
	req, err := http.NewRequest("GET", strings.ToLower(path), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", string(token)))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var r BookResponse
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

func GetRFQ(ticker string, side string, q float64) (*RFQResponse, error) {
	body := RFQRequest{ticker, side, q}
	pb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	postbody := bytes.NewBuffer(pb)
	path := fmt.Sprintf("%s/quote", baseURL)

	req, err := http.NewRequest("POST", path, postbody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", string(token)))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()


	var r RFQResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	if res.StatusCode == 201 {
		log.Printf(
			"Status: %s for %f %s\n",
			res.Status,
			r.Quantity,
			r.Pair,
		)
	} else {
		return nil, fmt.Errorf(
			"Error status: %s tradedesk: sFOX RFQ msg: %s",
			res.Status,
			r.Error,
		)
	}

	return &r, nil
}
