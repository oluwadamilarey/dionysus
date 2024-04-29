package p2p

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestTcpTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandShakeFunc: NOPHandShakeFunc,
		Decoder:       DefaultDecoder{},
	}

	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, tr.ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
