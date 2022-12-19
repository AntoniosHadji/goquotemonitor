package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestCoinbaseQuote(t *testing.T) {
	res, err := getBook("BTC")
	if err != nil {
		t.Errorf("FAIL: %#v", err)
	}

	fmt.Printf("%s %s\n", res.Bids[0][0], res.Bids[0][1])
	fmt.Printf("%s %s\n", res.Asks[0][0], res.Asks[0][1])
	fmt.Printf("%T %T\n", res.Asks[0][0], res.Asks[0][1])

	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][0]), 32); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][0]), 64); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][1]), 32); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][1]), 64); err == nil {
		fmt.Printf("%T: %f\n", s, s)
	}
}

func TestCalcSpread(t *testing.T) {
	f, e := os.ReadFile("./testdata/testcb.json")
	if e != nil {
		t.Errorf("file read failed: %#v\n", e)
	}

	var data CoinbaseBookResponse

	e = json.Unmarshal([]byte(f), &data)
	if e != nil {
		t.Errorf("Unmarshal failed: %#v\n", e)
	}

	// fmt.Printf("bids: %v\n", data.Bids)
	// fmt.Printf("asks: %v\n", data.Asks)

	bid, ask, result := calcSpread(&data, 1)
	fmt.Println(bid, ask, result, "bps")
	fmt.Printf("%.3f bps\n", result)

}

func TestCBWork(t *testing.T) {
	t.SkipNow()
	// this code does not return
	cbwork("BTC", 1)
}
