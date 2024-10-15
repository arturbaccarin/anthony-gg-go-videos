// https://youtu.be/oCm46sUILcs
// 28:00
package main

import (
	"projectx/network"
)

// Server
// Transport layer -> tcp, udp
// Block
// TX -> Transaction

func main() {
	trLocal := network.NewLocalTransport("LOCAL")

	opts := network.ServerOpts{
		Transport: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
}
