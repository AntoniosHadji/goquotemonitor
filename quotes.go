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

var tradeDesk = map[string]string{
	"DV":     "2fafbc29-d5a1-4a68-a8c6-cf93b043ad1b",
	"Enigma": "4c248890-703d-4ac3-9ce7-9de6465f328a",
}

var assets = map[string]string{
	"AVAX": "883a179b-91df-4b42-a31f-02b6c9736537",
	"BTC":  "43454116-7026-4b3d-a9de-eeaada500d4c",
	"ETH":  "efd4e846-99ec-4362-b21f-7982529bf570",
	"LTC":  "b2a5a8e2-d9d5-4854-8517-5a5608ffbe7d",
	"SOL":  "086d6588-6dec-4807-97c2-82e97ef6ff10",
	"USDC": "70c21c75-9362-4182-a984-daa63e30ee52",
	"USDT": "25311c35-26d0-4cf1-b916-80f377a7e468",
}

var account = "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b"

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
	Errors []map[string]interface{} `json:"errors,omitempty"`
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
	defer res.Body.Close()

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
		fmt.Printf("%#v\n", r.Errors)
		return nil, fmt.Errorf("Bad status: %s\n%#v", res.Status, r.Errors)
	}

	return &r, nil

}
