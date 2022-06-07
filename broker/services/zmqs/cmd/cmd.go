package cmd

import (
	"os"

	"github.com/alikarimii/zmqph/broker/services/zmqs"
	"github.com/alikarimii/zmqph/pkg/zerologger"
)

func RunZmqs() *zmqs.Service {
	stdLogger := zerologger.NewStandardLogger()
	config := zmqs.MustBuildConfigFromEnv(stdLogger)
	exitFn := func() { os.Exit(1) }
	zmqCOntext, er := zmqs.MustBuildZmqContext(config, stdLogger)
	if er != nil {
		stdLogger.Panic().Msgf("RunZmqs, %v", er)
	}
	que := zmqs.MustBuildDiskQueue(config, stdLogger)
	diContainer := zmqs.MustBuildDIContainer(
		config,
		stdLogger,
		zmqs.WithDiskQueue(que),
		zmqs.WithZmqSockets(zmqCOntext),
	)

	s := zmqs.InitService(config, stdLogger, exitFn, diContainer)
	return s
}
