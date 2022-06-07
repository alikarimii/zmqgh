package zmqs

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alikarimii/zmqph/pkg/zerologger"
)

func InitService(
	config *Config,
	logger *zerologger.Logger,
	exitFn func(),
	diContainter *DIContainer,
) *Service {
	return &Service{
		config:      config,
		logger:      logger,
		exitFn:      exitFn,
		diContainer: diContainter,
	}
}

type Service struct {
	config      *Config
	logger      *zerologger.Logger
	exitFn      func()
	diContainer *DIContainer
}

func (s *Service) StartZmq(ctx context.Context) {
	s.logger.Info().Msg("configuring zmq server ...")
	go func() {
		s.diContainer.GetCommandHandler().GettingMessageProcess(ctx)
	}()
}

func (s *Service) WaitForStopSignal() {
	s.logger.Info().Msg("start waiting for stop signal ...")

	stopSignalChannel := make(chan os.Signal, 1)
	signal.Notify(stopSignalChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-stopSignalChannel

	if _, ok := sig.(os.Signal); ok {
		s.logger.Info().Msgf("received '%s'", sig)
		close(stopSignalChannel)
		s.shutdown()
	}
}

func (s *Service) shutdown() {
	s.logger.Info().Msg("shutdown: stopping services ...")
	// stop service properly
	// like db connection

	s.logger.Info().Msg("shutdown: all services stopped!")

	s.exitFn()
}
