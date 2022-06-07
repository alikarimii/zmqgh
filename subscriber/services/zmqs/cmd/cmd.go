package cmd

import (
	"os"

	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/subscriber/services/zmqs"
)

func RunZmqs() *zmqs.Service {
	lg := zerologger.NewStandardLogger()
	config := zmqs.MustBuildConfigFromEnv(lg)
	fn := func() { os.Exit(1) }
	zmqCOntext, er := zmqs.MustBuildZmqContext(config, lg)
	if er != nil {
		lg.Panic().Msgf("RunZmqs, %v", er)
	}
	diContainer := zmqs.MustBuildDIContainer(
		config,
		lg,
		zmqs.WithZmqSockets(zmqCOntext),
	)
	s := zmqs.InitService(config, lg, fn, diContainer)
	return s
}
