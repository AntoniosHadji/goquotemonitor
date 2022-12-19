package main

import "testing"

func TestWorker(t *testing.T) {
	// this code does not end
	t.SkipNow()

	w := Work{"Enigma", "BTC", 1.0}
	dowork(w)
}
