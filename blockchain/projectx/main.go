// https://youtu.be/oCm46sUILcs
// 39:40
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
	trRemote := network.NewLocalTransport("REMOTE")

	opts := network.ServerOpts{
		Transport: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
