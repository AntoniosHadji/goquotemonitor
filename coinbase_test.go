package main

import (
	"fmt"
	"testing"
)

func TestCoinbaseQuote(t *testing.T) {
	res, err := getBook("BTC")
	if err != nil {
		t.Errorf("FAIL: %v", err)
	}

	fmt.Println(res.Bids[0])
	fmt.Println(res.Asks[0])
}
