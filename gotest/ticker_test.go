package main

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				t.Log("A: stopped")
				return
			case time := <-ticker.C:
				t.Log("A: Tick at", time)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				t.Log("B: stopped")
				return
			case time := <-ticker.C:
				t.Log("B: Tick at", time)
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond)
	done <- true
	done <- true
	ticker.Stop()
	t.Log("Ticker stopped")
}
