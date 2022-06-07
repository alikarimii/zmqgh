package main

import (
	"context"

	"github.com/alikarimii/zmqph/publisher/services/zmqs/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	s := cmd.RunZmqs()
	s.StartZmq(ctx)
	s.WaitForStopSignal()
	cancel()
}
