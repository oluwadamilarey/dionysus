package p2p

type HandShakeFunc func(*TCPPeer) error

func NOPHandShakeFunc(*TCPPeer) error {
	return nil
}
