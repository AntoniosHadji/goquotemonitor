package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCoinbaseQuote(t *testing.T) {
	res, err := getBook("BTC")
	if err != nil {
		t.Errorf("FAIL: %v", err)
	}

	fmt.Printf("%s %s\n", res.Bids[0][0], res.Bids[0][1])
	fmt.Printf("%s %s\n", res.Asks[0][0], res.Asks[0][1])
	fmt.Printf("%T %T\n", res.Asks[0][0], res.Asks[0][1])

	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][0]), 32); err == nil {
		fmt.Println(s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][0]), 64); err == nil {
		fmt.Println(s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][1]), 32); err == nil {
		fmt.Println(s)
	}
	if s, err := strconv.ParseFloat(fmt.Sprintf("%v", res.Asks[0][1]), 64); err == nil {
		fmt.Println(s)
	}
}
