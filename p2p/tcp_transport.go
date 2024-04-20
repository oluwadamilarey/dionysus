package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

// a peer represents a node on extablished tcp network
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if conn is dial and retrieve aa connection => outbound => true
	// if we accept and retrieve a connection => outbound => false
	outbound bool
}

func newTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakeFunc HandShakeFunc
	decoder       Decoder

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := newTCPPeer(conn, true)
	if err := t.shakeHands(conn); err != nil {

	}
	buf := new(bytes.Buffer)
	for {
		n, _ := conn.Read(buf)
	}
	fmt.Printf("new incoming connection %+v\n", peer)
}