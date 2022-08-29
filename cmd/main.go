package main

import (
	"github.com/hensur/go-arp"
	"net"
	"os"
)

func main() {
	mode := os.Args[1]

	ip := net.ParseIP(os.Args[2])
	ipnet := net.IPNet{
		IP: ip,
	}

	if mode == "add" {
		err := arp.SetARP(os.Args[3], ipnet)
		if err != nil {
			panic(err)
		}
	} else {
		err := arp.DeleteARP(os.Args[3], ipnet)
		if err != nil {
			panic(err)
		}
	}

}
