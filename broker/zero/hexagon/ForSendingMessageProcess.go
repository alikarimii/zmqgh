package hexagon

// in port
// inject in driving adapter(input adapter) in infrastructure. like grpc,rest
// send to destination
type ForSendingMessageProcess func() chan struct{}
