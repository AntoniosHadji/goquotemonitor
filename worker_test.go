package main

import "testing"

func TestCopyReqBody(t *testing.T) {

	w := Work{"Enigma", "BTC", 1.0}

	dowork(w)
}
