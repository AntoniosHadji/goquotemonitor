package sfox

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var baseURL = "https://api.sfox.com/v1"
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

func GetBook(ticker string) (*BookResponse, error) {

	path := fmt.Sprintf("%s/markets/orderbook/%susd", baseURL, ticker)
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
