package hexagon

// in port
// inject in driving adapter(input adapter) in infrastructure. like grpc,rest
// get message from source
type ForGettingMessageProcess func() chan struct{}
