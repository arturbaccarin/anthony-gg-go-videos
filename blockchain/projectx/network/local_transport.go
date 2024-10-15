// to test - any connection peers
package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.Mutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC { // read the channel, sending is chan<-
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

/*************  ✨ Codeium Command ⭐  *************/
/******  d1436587-08ce-4c2b-8bc2-0ee4d10b26ab  *******/
func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
