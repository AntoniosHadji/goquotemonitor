package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamps(t *testing.T) {
	fmt.Println(time.Now().UTC())
	// 2022-12-14 20:28:42.72878716 +0000 UTC
}

func TestMultipleDuration(t *testing.T) {
	d := 60
	fmt.Println(time.Duration(d) * time.Second)
}
