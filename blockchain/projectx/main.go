// https://youtu.be/kYJyzTkIZjg
// 10:00 - commit of the day. sorry guys...
package main

import (
	"projectx/network"
	"time"
)

// Server
// Transport layer -> tcp, udp
// Block
// TX -> Transaction

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello remote"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transport: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
