package p2p

import (
	"fmt"
	"net"
	"sync"
)

// a peer represents a node on extablished tcp network
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if conn is dial and retrieve a connection => outbound => true
	// if we accept and retrieve a connection => outbound => false
	outbound bool
}

func newTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
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

		fmt.Printf("new imcoming connection %s\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := newTCPPeer(conn, true)
	if err := t.HandShakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake error %s\n", err)
		return
	}

	//read loop
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %v", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		fmt.Printf("message: %+v\n", msg)
	}
}
