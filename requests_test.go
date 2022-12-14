package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Not Used, only for example purposes ========================================

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

// QuoteRequest payload for request to /v2/quotes API
type QuoteRequest2 struct {
	Data struct {
		Type              string `json:"type"`
		QuoteRequestAttrs `json:"attributes"`
	} `json:"data"`
}

//=============================================================================

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

	// embedded attrs
	var APIReq2 = new(QuoteRequest2)
	APIReq2.Data.Type = "quotes"
	APIReq2.Data.AccountID = "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b"
	APIReq2.Data.AssetID = "43454116-7026-4b3d-a9de-eeaada500d4c"
	APIReq2.Data.Hot = false
	APIReq2.Data.TransactionType = "buy"
	APIReq2.Data.UnitCount = 1.0
	APIReq2.Data.TradeDeskID = "4c248890-703d-4ac3-9ce7-9de6465f328a"
	APIReq2.Data.DelayedSettlement = true
	fmt.Println(APIReq2)

	var reqdata = &PTRequestBody{
		QRData{
			Type: "quotes",
			Attributes: QuoteRequestAttrs{
				AccountID:         "0c7715e3-7cdd-4d49-88bb-f1ab3cb8803b",
				AssetID:           "43454116-7026-4b3d-a9de-eeaada500d4c",
				Hot:               false,
				TransactionType:   "sell",
				UnitCount:         1.0,
				TradeDeskID:       "4c248890-703d-4ac3-9ce7-9de6465f328a",
				DelayedSettlement: true,
			},
		},
	}
	fmt.Println(reqdata)

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan *QuoteResponse, 2)

	start := time.Now()
	// Code to measure
	go func() {
		defer wg.Done()

		pb, err := json.Marshal(APIReq)
		if err != nil {
			fmt.Println(err)
		}
		r, err := ptQuoteRequest(bytes.NewBuffer(pb))
		if err != nil {
			fmt.Println(err)
		}
		ch <- r
	}()

	go func() {
		defer wg.Done()

		pb, err := json.Marshal(reqdata)
		if err != nil {
			fmt.Println(err)
		}
		r, err := ptQuoteRequest(bytes.NewBuffer(pb))
		if err != nil {
			fmt.Println(err)
		}
		ch <- r
	}()

	fmt.Println("Waiting...")
	wg.Wait()
	duration := time.Since(start)
	close(ch)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Printf("duration: %v\n", duration)

	for r := range ch {
		fmt.Println(r)
		fmt.Println(r.Data.Attributes.PricePerUnit)
	}
	fmt.Println("Finished loop over channel")

}

func TestEmptyQuoteResponse(t *testing.T) {
	qr := new(QuoteResponse)
	fmt.Println(qr)

	if qr != nil {
		fmt.Println("qr is not nil")
	}
	fmt.Println(qr.Data.Attributes.PricePerUnit)

}
