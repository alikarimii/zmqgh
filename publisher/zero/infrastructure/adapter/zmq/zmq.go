package zmq

import "context"

type forListenToDestination func(ctx context.Context, fromDB <-chan []byte) chan struct{}

func NewZMQ(
	listenToDestination forListenToDestination,
) *ZMQ {
	return &ZMQ{
		listenToDestination,
	}
}

type ZMQ struct {
	listenToDestination forListenToDestination
}

func (outputAdapter *ZMQ) BroadcastMessageToDestination(ctx context.Context, data <-chan []byte) chan struct{} {
	return outputAdapter.listenToDestination(ctx, data)
}
