package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
	fmt.Println(r.Errors[0]["status"])
	fmt.Println(r.Errors[0]["title"])
	fmt.Println(r.Errors[0]["detail"])
	fmt.Println(r.Errors[0]["source"])

}

// === RUN   TestError
// {{
//   { 0 0001-01-01 00:00:00 +0000 UTC false 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 false  false 0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC  0  0}}
//     [map[detail:cannot be over 5000.0 for ACH source:map[pointer:/data/attributes/amount] status:422 title:Invalid Attribute]]}
// --- PASS: TestError (0.00s)
// PASS
// ok  	command-line-arguments	0.004s
