package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestCtrlC(t *testing.T) {
	for {
		t.Log("Logging")
		time.Sleep(1500 * time.Millisecond)
	}
}

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Caught CTRL-C")
		fmt.Println("Run Cleanup")
		os.Exit(1)
	}()
}
