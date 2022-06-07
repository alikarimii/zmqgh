package zmq

import "context"

type forListenToSource func(ctx context.Context, toDB chan<- []byte) chan struct{}

func NewZMQ(
	listenToSource forListenToSource,
) *ZMQ {
	return &ZMQ{
		listenToSource,
	}
}

type ZMQ struct {
	listenToSource forListenToSource
}

func (outputAdapter *ZMQ) ReadingMessageFromSource(ctx context.Context, data chan<- []byte) chan struct{} {
	return outputAdapter.listenToSource(ctx, data)
}
