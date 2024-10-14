// to test - any connection peers
package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.Mutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC { // read the channel, sending is chan<-
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
