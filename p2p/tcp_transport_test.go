package p2p

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestTcpTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOpts{ListenAddr: listenAddr})

	assert.Equal(t, tr.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
