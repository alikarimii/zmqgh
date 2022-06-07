package application

import "context"

// https://medium.com/idealo-tech-blog/hexagonal-ports-adapters-architecture-e3617bcf00a0
// An output port (driven port) is another type of interface
// that is used by the application core
// to reach things outside of itself (like getting some data from a database).

// out port
// implement(method of struct) in output adapter (driven adapter)
// type outputAdapter struct{}
// func (adapter outputAdapter) RetrievingMessage func() error {}
type ForRetrievingMessage func(ctx context.Context) <-chan []byte
