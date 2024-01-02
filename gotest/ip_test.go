package main

import (
	"fmt"
	"net"
	"testing"
)

func TestGetIP(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Println("IP Address:", ipNet.IP.String())
		}
	}
}
