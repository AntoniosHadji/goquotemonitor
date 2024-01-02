package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/antonioshadji/goquotemonitor/db"
)

func TestPTRequestBody(t *testing.T) {

	var reqdata = QuoteRequest{}
	reqdata.Data.Type = "quotes"
	reqdata.Data.Attributes.AccountID = db.Account
	reqdata.Data.Attributes.AssetID = db.Assets["BTC"]
	reqdata.Data.Attributes.TransactionType = "sell"
	reqdata.Data.Attributes.UnitCount = 1.0
	reqdata.Data.Attributes.Hot = false
	reqdata.Data.Attributes.TradeDeskID = db.TradeDesk["Enigma"]
	reqdata.Data.Attributes.DelayedSettlement = true

	pb, err := json.Marshal(reqdata)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pb)

	// pb2, err := json.Marshal(reqdataOrig)
	// if err != nil {
	// 	t.Error(err)
	// }

	// if bytes.Compare(pb, pb2) != 0 {
	// 	fmt.Println(pb)
	// 	fmt.Println(pb2)
	// }
}

func TestPrepareQuoteRequest(t *testing.T) {

	var reqdata = QuoteRequest{}
	reqdata.Data.Type = "quotes"
	reqdata.Data.Attributes.AccountID = db.Account
	reqdata.Data.Attributes.AssetID = db.Assets["BTC"]
	reqdata.Data.Attributes.TransactionType = "sell"
	reqdata.Data.Attributes.UnitCount = 1.0
	reqdata.Data.Attributes.Hot = false
	reqdata.Data.Attributes.TradeDeskID = db.TradeDesk["Enigma"]
	reqdata.Data.Attributes.DelayedSettlement = true

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
