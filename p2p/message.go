package p2p

import "net"

// message owns any arbitrary data that is sent over the
//
//	each transport between the 2 nodes in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
