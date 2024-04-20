package p2p

// Peer is an interface that represent the remote node
type Peer interface{}

// transport is anything that handles the communication
// between the nodes on the network
// this could be UDP, TCP, WebSockets
type Transport interface {
	ListenAndAccept() error
}
