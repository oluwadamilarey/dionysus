package p2p

import "net"

// message owns any arbitrry data that is sent over the
//
//	each transport between the 2 nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
