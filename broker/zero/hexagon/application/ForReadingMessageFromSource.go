package application

import "context"

// An output port (driven port) is another type of interface
// that is used by the application core
// to reach things outside of itself (like getting some data from a database).

// out port
// implement(method of struct) in output adapter (driven adapter)
// type outputAdapter struct{}
// func (adapter outputAdapter) ReadingMessage func(data chan []byte) error {}

type ForReadingMessageFromSource func(ctx context.Context, data chan<- []byte) chan struct{}
