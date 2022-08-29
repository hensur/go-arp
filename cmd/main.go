package main

import (
	"github.com/hensur/go-arp"
	"net"
	"os"
)

func main() {
	ip := net.ParseIP(os.Args[1])
	ipnet := net.IPNet{
		IP: ip,
	}

	err := arp.SetARP(os.Args[2], ipnet)
	if err != nil {
		panic(err)
	}
}
