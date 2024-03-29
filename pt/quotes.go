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
		ID         string `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
		Attributes struct {
			AssetName         string    `json:"asset-name,omitempty"`
			BaseAmount        float64   `json:"base-amount,omitempty"`
			CreatedAt         time.Time `json:"created-at,omitempty"`
			DelayedSettlement bool      `json:"delayed-settlement,omitempty"`
			ExecutedAt        time.Time `json:"executed-at,omitempty"`
			ExpiresAt         time.Time `json:"expires-at,omitempty"`
			FeeAmount         float64   `json:"fee-amount,omitempty"`
			Hot               bool      `json:"hot,omitempty"`
			TradeID           string    `json:"trade-id,omitempty"`
			IntegratorSettled bool      `json:"integrator-settled,omitempty"`
			PricePerUnit      float64   `json:"price-per-unit,omitempty"`
			RejectedAt        time.Time `json:"rejected-at,omitempty"`
			SettledAt         time.Time `json:"settled-at,omitempty"`
			Status            string    `json:"status,omitempty"`
			TotalAmount       float64   `json:"total-amount,omitempty"`
			TransactionType   string    `json:"transaction-type,omitempty"`
			UnitCount         float64   `json:"unit-count,omitempty"`
		} `json:"attributes,omitempty"`
	} `json:"data,omitempty"`
	Errors Errors `json:"errors,omitempty"`
}

// Errors API Error response
type Errors []struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Source struct {
		Pointer string `json:"pointer,omitempty"`
	} `json:"source,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// QuoteRequest payload for request to /v2/quotes API
type QuoteRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			AccountID         string  `json:"account-id"`
			AssetID           string  `json:"asset-id"`
			TransactionType   string  `json:"transaction-type"`
			UnitCount         float64 `json:"unit-count"`
			Hot               bool    `json:"hot,omitempty"`
			DelayedSettlement bool    `json:"delayed-settlement,omitempty"`
			TradeDeskID       string  `json:"trade-desk-id,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

// ====================================================================

func ptQuoteRequest(payload *QuoteRequest) (*QuoteResponse, error) {
	pb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(pb)

	path := fmt.Sprintf("%s/%s/quotes", ptBaseURL, version)
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", string(p)))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r QuoteResponse
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
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
		return nil, fmt.Errorf(
			"Error status: %s asset: %s tradedesk: %s\n%#v",
			res.Status,
			payload.Data.Attributes.AssetID,
			payload.Data.Attributes.TradeDeskID,
			r.Errors[0].Detail)
	}

	return &r, nil

}
