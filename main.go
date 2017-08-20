package main

import (
	"flag"
	"mychat/lib"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go - listen <ip>
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		connIP := os.Args[1]
		// go run main.go <ip>
		lib.RunGuest(connIP)
	}
}
