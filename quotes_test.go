package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

// QRData ... original struct used before upgrade above
type QRData struct {
	Type       string            `json:"type"`
	Attributes QuoteRequestAttrs `json:"attributes"`
}

func TestPTRequestBody(t *testing.T) {

	var reqdata = PTQuotesRequestBody{}
	reqdata.Data.Type = "quotes"
	reqdata.Data.Attributes = QuoteRequestAttrs{
		AccountID:         account,
		AssetID:           assets["BTC"],
		TransactionType:   "sell",
		UnitCount:         1.0,
		Hot:               false,
		TradeDeskID:       tradeDesk["Enigma"],
		DelayedSettlement: true,
	}

	var reqdataOrig = PTQuotesRequestBody{
		QRData{
			Type: "quotes",
			Attributes: QuoteRequestAttrs{
				AccountID:         account,
				AssetID:           assets["BTC"],
				TransactionType:   "sell",
				UnitCount:         1.0,
				Hot:               false,
				TradeDeskID:       tradeDesk["Enigma"],
				DelayedSettlement: true,
			},
		},
	}

	pb, err := json.Marshal(reqdata)
	if err != nil {
		t.Error(err)
	}
	pb2, err := json.Marshal(reqdataOrig)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(pb, pb2) != 0 {
		fmt.Println(pb)
		fmt.Println(pb2)
	}
}

func TestPrepareQuoteRequest(t *testing.T) {

	var reqdata = PTQuotesRequestBody{
		QRData{
			Type: "quotes",
			Attributes: QuoteRequestAttrs{
				AccountID:         account,
				AssetID:           assets["BTC"],
				Hot:               false,
				TransactionType:   "sell",
				UnitCount:         1.0,
				TradeDeskID:       tradeDesk["Enigma"],
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
