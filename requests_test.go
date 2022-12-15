package main

import (
	"fmt"
	"testing"
)

func TestPrepareQuoteRequest(t *testing.T) {

	var reqdata = PTRequestBody{
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

	var reqdata2 = reqdata
	reqdata2.Data.Attributes.TransactionType = "buy"

	fmt.Println(reqdata)
	fmt.Println(reqdata2)

	if reqdata2.Data.Attributes.TransactionType == reqdata.Data.Attributes.TransactionType {
		t.Fatalf("Failed because two requests are not different.")
	}

}

func TestEmptyQuoteResponse(t *testing.T) {
	qr := new(QuoteResponse)
	fmt.Println(qr)

	if qr != nil {
		fmt.Println("qr is not nil")
	}
	fmt.Println(qr.Data.Attributes.PricePerUnit)

}
