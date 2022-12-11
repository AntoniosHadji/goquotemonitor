package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

var p = "eyJhbGciOiJIUzI1NiJ9.eyJhdXRoX3NlY3JldCI6ImEzNTUzZDhiLWYyNTYtNGI0My04ZGNhLTM0N2Y0MzkyNTRhYiIsInVzZXJfZ3JvdXBzIjpbImJvdHMiLCJjYXNoX21hbmFnZXJzIiwiY29tcGxpYW5jZV9vZmZpY2VycyIsImN1c3RvbWVyX3NlcnZpY2VfcmVwcyIsImVuZ2luZWVycyIsImludmVzdG1lbnRfb2ZmaWNlcnMiLCJsaXF1aWRpdHlfbWFuYWdlcnMiLCJ0cmFkZV9kZXNrcyIsInNhbGVzX21hbmFnZXJzIiwidHJ1c3Rfb2ZmaWNlcnMiXSwibm93IjoxNjQzMzk2NDk3LCJleHAiOjE2NzI1MzEyMDB9.eLCS46U_4MEy7oKTvMm8X1uH4N2XXZ8T4NU5VLZgfNY"
var client = &http.Client{}
var baseurl = "https://api.primetrust.com"

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

// QuoteRequest payload for request to /v2/quotes API
type QuoteRequest struct {
	Data struct {
		Type  string `json:"type"`
		Attrs struct {
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

func TestSendQuoteRequest(t *testing.T) {
	var APIReq = new(QuoteRequest)
	APIReq.Data.Type = "quotes"
	APIReq.Data.Attrs.AccountID = "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b"
	APIReq.Data.Attrs.AssetID = "43454116-7026-4b3d-a9de-eeaada500d4c"
	APIReq.Data.Attrs.Hot = false
	APIReq.Data.Attrs.TransactionType = "buy"
	APIReq.Data.Attrs.UnitCount = 1.0
	APIReq.Data.Attrs.TradeDeskID = "4c248890-703d-4ac3-9ce7-9de6465f328a"
	APIReq.Data.Attrs.DelayedSettlement = true
	fmt.Println(APIReq)

	var reqdata = &PTRequestBody{
		QRData{
			Type: "quotes",
			Attributes: QuoteRequestAttrs{
				AccountID:         "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b",
				AssetID:           "43454116-7026-4b3d-a9de-eeaada500d4c",
				Hot:               false,
				TransactionType:   "buy",
				UnitCount:         1.0,
				TradeDeskID:       "4c248890-703d-4ac3-9ce7-9de6465f328a",
				DelayedSettlement: true,
			},
		},
	}
	fmt.Println(reqdata)

	pb, err := json.Marshal(APIReq)
	fmt.Println(string(pb))
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	send("POST", fmt.Sprintf("%s/v2/quotes", baseurl), bytes.NewBuffer(pb))

	pb, err = json.Marshal(reqdata)
	fmt.Println(string(pb))
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	send("POST", fmt.Sprintf("%s/v2/quotes", baseurl), bytes.NewBuffer(pb))
}

func send(method, url string, payload io.Reader) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("request error")
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", string(p)))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("http client error")
		fmt.Println(err)
		return
	}
	fmt.Println(res.Status)

	if res.StatusCode == 201 {
		var r QuoteResponse
		if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
			fmt.Println("unmarshal error")
			fmt.Println(err)
		}
		fmt.Println(r)
		fmt.Println(r.Data.Attributes.PricePerUnit)
	} else {
		fmt.Println(res.Status)
		fmt.Println(res)
		fmt.Println(res.Body)
	}

}
