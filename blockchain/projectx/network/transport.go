package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte // could be any
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
