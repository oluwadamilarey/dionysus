package p2p

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestTcpTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
