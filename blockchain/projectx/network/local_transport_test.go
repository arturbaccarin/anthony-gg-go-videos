package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestConnect(t *testing.T) {
// 	tra := NewLocalTransport("A")
// 	trb := NewLocalTransport("B")

// 	tra.Connect(trb)
// 	trb.Connect(tra)

// 	assert.Equal(t, tra.peers[trb.Addr()], trb)
// 	assert.Equal(t, trb.peers[tra.Addr()], tra)
// }

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.Consume()
	assert.Equal(t, msg, rpc.Payload)
	assert.Equal(t, tra.Addr(), rpc.From)
}
