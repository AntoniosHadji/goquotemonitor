package main

import (
	"testing"

	"github.com/antonioshadji/goquotemonitor/db"
)

func TestWorker(t *testing.T) {
	// this code does not end
	t.SkipNow()

	w := db.Work{
		LP:     "Enigma",
		Ticker: "BTC",
		Size:   1.0,
	}
	dowork(w)
}
