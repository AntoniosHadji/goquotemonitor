package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var p = "eyJhbGciOiJIUzI1NiJ9.eyJhdXRoX3NlY3JldCI6ImEzNTUzZDhiLWYyNTYtNGI0My04ZGNhLTM0N2Y0MzkyNTRhYiIsInVzZXJfZ3JvdXBzIjpbImJvdHMiLCJjYXNoX21hbmFnZXJzIiwiY29tcGxpYW5jZV9vZmZpY2VycyIsImN1c3RvbWVyX3NlcnZpY2VfcmVwcyIsImVuZ2luZWVycyIsImludmVzdG1lbnRfb2ZmaWNlcnMiLCJsaXF1aWRpdHlfbWFuYWdlcnMiLCJ0cmFkZV9kZXNrcyIsInNhbGVzX21hbmFnZXJzIiwidHJ1c3Rfb2ZmaWNlcnMiXSwibm93IjoxNjQzMzk2NDk3LCJleHAiOjE2NzI1MzEyMDB9.eLCS46U_4MEy7oKTvMm8X1uH4N2XXZ8T4NU5VLZgfNY"
var client = &http.Client{}
var baseurl = "https://api.primetrust.com"
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
	} `json:"data"`
}

// PTRequestBody payload for request to /v2/quotes API
type PTRequestBody struct {
	Data QRData `json:"data"`
}

// QRData ...
type QRData struct {
	Type       string            `json:"type"`
	Attributes QuoteRequestAttrs `json:"attributes"`
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

func ptQuoteRequest(payload io.Reader) (*QuoteResponse, error) {
	path := fmt.Sprintf("%s/%s/quotes", baseurl, version)
	req, err := http.NewRequest("POST", path, payload)
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
		fmt.Printf("%#v", r)
		return nil, fmt.Errorf("Bad status: %s\n%#v", res.Status, r)
	}

	return &r, nil

}
