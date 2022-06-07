package zmqs

import (
	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/publisher/zero/hexagon/application"
	"github.com/alikarimii/zmqph/publisher/zero/infrastructure/adapter/bytegen"
	"github.com/alikarimii/zmqph/publisher/zero/infrastructure/adapter/zmq"
	"github.com/cockroachdb/errors"
	"github.com/pebbe/zmq4"
)

type DIOption func(container *DIContainer) error

func WithZmqSockets(ctx *zmq4.Context) DIOption {
	return func(container *DIContainer) error {
		if ctx == nil {
			return errors.New("must init context socket first")
		}
		container.infra.ctx = ctx
		return nil
	}
}
func MustBuildDIContainer(config *Config, logger *zerologger.Logger, opts ...DIOption) *DIContainer {
	container := &DIContainer{
		config: config,
		logger: logger,
	}

	for _, v := range opts {
		if err := v(container); err != nil {
			logger.Panic().Msgf("mustBuildDIContainer: %s", err)
		}
	}
	container.init()

	return container
}
func MustBuildZmqContext(config *Config, logger *zerologger.Logger) (*zmq4.Context, error) {
	ctx, er := zmq4.NewContext()
	if er != nil {
		return nil, errors.Wrap(er, "MustBuildZmqContext")
	}
	return ctx, nil
}
func MustBuildGenerator(logger *zerologger.Logger) (*bytegen.Generator, error) {
	g := bytegen.NewGenerator(logger)
	return g, nil
}

type DIContainer struct {
	config *Config
	logger *zerologger.Logger
	infra  struct {
		ctx *zmq4.Context
	}
	service struct {
		commandHandler *application.CommandHandler
		queryHandler   *application.QueryHanndler
		byteGen        *bytegen.ByteGenerator
		zmq            *zmq.ZMQ
	}
	serviceDriver struct {
		generator *bytegen.Generator
		zmqDriver *zmq.ZmqDriver
	}
}

func (container *DIContainer) init() {
	_ = container.GetCommandHandler()
	_ = container.GetQueryHandler()
	_ = container.GetZMQ()
	_ = container.GetByteGen()
}

func (container *DIContainer) GetCommandHandler() *application.CommandHandler {
	if container.service.commandHandler == nil {
		container.service.commandHandler = application.NewCommandHandler(
			container.GetZMQ().BroadcastMessageToDestination,
			container.GetByteGen().GetGeneratedMessage,
		)
	}
	return container.service.commandHandler
}

func (container *DIContainer) GetQueryHandler() *application.QueryHanndler {
	if container.service.queryHandler == nil {
		container.service.queryHandler = application.NewQueryHandler(
			container.GetByteGen().MessageCount,
		)
	}
	return container.service.queryHandler
}
func (container *DIContainer) GetZMQ() *zmq.ZMQ {
	if container.service.zmq == nil {
		container.service.zmq = zmq.NewZMQ(
			container.getZMQDriver().ListenToDestination,
		)
	}
	return container.service.zmq
}
func (container *DIContainer) GetByteGen() *bytegen.ByteGenerator {
	if container.service.byteGen == nil {
		container.service.byteGen = bytegen.NewByteGenerator(
			container.getBytegenDriver().Generate,
			container.getBytegenDriver().MessageCount,
		)
	}
	return container.service.byteGen
}
func (container *DIContainer) getZMQDriver() *zmq.ZmqDriver {
	if container.serviceDriver.zmqDriver == nil {
		container.serviceDriver.zmqDriver = zmq.NewZmqDriver(
			container.config.ZMQ,
			container.infra.ctx,
		)
	}
	return container.serviceDriver.zmqDriver
}
func (container *DIContainer) getBytegenDriver() *bytegen.Generator {
	if container.serviceDriver.generator == nil {
		container.serviceDriver.generator = bytegen.NewGenerator(container.logger)
	}
	return container.serviceDriver.generator
}
