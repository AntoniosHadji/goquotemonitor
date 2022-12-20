package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var p string

func init() {
	val, ok := os.LookupEnv("TOKEN")
	if ok {
		p = val
	} else {
		fmt.Println("environment variable TOKEN is not Defined")
		os.Exit(1)
	}
}

var client = &http.Client{}
var ptBaseURL = "https://api.primetrust.com"
var version = "v2"

// QuoteResponse struct to store response from /v2/quotes API
type QuoteResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			AssetName         string    `json:"asset-name"`
			BaseAmount        float64   `json:"base-amount"`
			CreatedAt         time.Time `json:"created-at"`
			DelayedSettlement bool      `json:"delayed-settlement"`
			ExecutedAt        time.Time `json:"executed-at"`
			ExpiresAt         time.Time `json:"expires-at"`
			FeeAmount         float64   `json:"fee-amount"`
			Hot               bool      `json:"hot"`
			TradeID           string    `json:"trade-id"`
			IntegratorSettled bool      `json:"integrator-settled"`
			PricePerUnit      float64   `json:"price-per-unit"`
			RejectedAt        time.Time `json:"rejected-at"`
			SettledAt         time.Time `json:"settled-at"`
			Status            string    `json:"status"`
			TotalAmount       float64   `json:"total-amount"`
			TransactionType   string    `json:"transaction-type"`
			UnitCount         float64   `json:"unit-count"`
		} `json:"attributes"`
	} `json:"data,omitempty"`
}

// PTQuotesRequestBody payload for request to /v2/quotes API
type PTQuotesRequestBody struct {
	Data struct {
		Type       string            `json:"type"`
		Attributes QuoteRequestAttrs `json:"attributes"`
	} `json:"data"`
}

// QuoteRequestAttrs is attributes for Quote request
type QuoteRequestAttrs struct {
	AccountID         string  `json:"account-id"`
	AssetID           string  `json:"asset-id"`
	TransactionType   string  `json:"transaction-type"`
	UnitCount         float64 `json:"unit-count"`
	Hot               bool    `json:"hot,omitempty"`
	DelayedSettlement bool    `json:"delayed-settlement,omitempty"`
	TradeDeskID       string  `json:"trade-desk-id,omitempty"`
}

// ====================================================================

func ptQuoteRequest(payload *PTQuotesRequestBody) (*QuoteResponse, error) {
	pb, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	body := bytes.NewBuffer(pb)

	path := fmt.Sprintf("%s/%s/quotes", ptBaseURL, version)
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", string(p)))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var r QuoteResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println(err)
		return nil, err
	}
	if res.StatusCode == 201 {
		log.Printf(
			"Status: %s for %f %s\n",
			res.Status,
			r.Data.Attributes.UnitCount,
			r.Data.Attributes.AssetName,
		)
	} else {
		fmt.Println(res.Status)
		fmt.Println(res)
		fmt.Printf("%#v\n", r)
		return nil, fmt.Errorf("Bad status: %s\n%#v", res.Status, r)
	}

	return &r, nil

}
