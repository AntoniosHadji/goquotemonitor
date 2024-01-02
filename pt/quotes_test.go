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

func TestError(t *testing.T) {
	//Simple Employee JSON array which we will parse
	errors := `
{
  "errors": [
    {
      "status": 422,
      "title": "Invalid Attribute",
      "source": {
        "pointer": "/data/attributes/amount"
      },
      "detail": "cannot be over 5000.0 for ACH"
    }
  ]
}`
	// Declared an empty interface of type Array
	//var results []map[string]interface{}
	var r QuoteResponse

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(errors), &r)
	fmt.Println(r)
	fmt.Println(r.Errors)
	fmt.Println(r.Errors[0])
	// TODO: error resolve
	// fmt.Println(r.Errors[0]["status"])

	// fmt.Println(r.Errors[0]["title"])
	// fmt.Println(r.Errors[0]["detail"])
	// fmt.Println(r.Errors[0]["source"])

}

// === RUN   TestError
// {{
//   { 0 0001-01-01 00:00:00 +0000 UTC false 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 false  false 0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC  0  0}}
//     [map[detail:cannot be over 5000.0 for ACH source:map[pointer:/data/attributes/amount] status:422 title:Invalid Attribute]]}
// --- PASS: TestError (0.00s)
// PASS
// ok  	command-line-arguments	0.004s
