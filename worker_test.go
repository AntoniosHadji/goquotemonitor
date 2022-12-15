package main

import "testing"

func TestWorker(t *testing.T) {

	w := Work{"Enigma", "BTC", 1.0}

	dowork(w)
}
