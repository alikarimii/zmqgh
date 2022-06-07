package zmq

import "context"

type forListenToSource func(ctx context.Context, toDB chan<- []byte) chan struct{}
type forListenToDestination func(ctx context.Context, fromDB <-chan []byte) chan struct{}

func NewZMQ(
	listenToSource forListenToSource,
	listenToDestination forListenToDestination,
) *ZMQ {
	return &ZMQ{
		listenToSource,
		listenToDestination,
	}
}

type ZMQ struct {
	listenToSource      forListenToSource
	listenToDestination forListenToDestination
}

func (outputAdapter *ZMQ) BroadcastMessageToDestination(ctx context.Context, data <-chan []byte) chan struct{} {
	return outputAdapter.listenToDestination(ctx, data)
}
func (outputAdapter *ZMQ) ReadingMessageFromSource(ctx context.Context, data chan<- []byte) chan struct{} {
	return outputAdapter.listenToSource(ctx, data)
}
