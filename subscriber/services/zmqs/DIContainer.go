package zmqs

import (
	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/subscriber/zero/hexagon/application"
	"github.com/alikarimii/zmqph/subscriber/zero/infrastructure/adapter/messagedb"
	"github.com/alikarimii/zmqph/subscriber/zero/infrastructure/adapter/zmq"
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

type DIContainer struct {
	config *Config
	logger *zerologger.Logger
	infra  struct {
		ctx *zmq4.Context
	}
	service struct {
		commandHandler *application.CommandHandler
		queryHandler   *application.QueryHanndler
		messagedb      *messagedb.Messagedb
		zmq            *zmq.ZMQ
	}
	serviceDriver struct {
		mdbDriver *messagedb.MdbDriver
		zmqDriver *zmq.ZmqDriver
	}
}

func (container *DIContainer) init() {
	_ = container.GetCommandHandler()
	_ = container.GetQueryHandler()
	_ = container.GetZMQ()
	_ = container.GetMessagedb()
}

func (container *DIContainer) GetCommandHandler() *application.CommandHandler {
	if container.service.commandHandler == nil {
		container.service.commandHandler = application.NewCommandHandler(
			container.GetZMQ().ReadingMessageFromSource,
			container.GetMessagedb().SavingMessage,
		)
	}
	return container.service.commandHandler
}

func (container *DIContainer) GetQueryHandler() *application.QueryHanndler {
	if container.service.queryHandler == nil {
		container.service.queryHandler = application.NewQueryHandler(
			container.GetMessagedb().MessageCount,
		)
	}
	return container.service.queryHandler
}
func (container *DIContainer) GetZMQ() *zmq.ZMQ {
	if container.service.zmq == nil {
		container.service.zmq = zmq.NewZMQ(
			container.getZMQDriver().ListenToSource,
		)
	}
	return container.service.zmq
}
func (container *DIContainer) GetMessagedb() *messagedb.Messagedb {
	if container.service.messagedb == nil {
		container.service.messagedb = messagedb.NewMessagedb(
			container.getMdbDriver().Save,
			container.getMdbDriver().MessageCount,
		)
	}
	return container.service.messagedb
}
func (container *DIContainer) getZMQDriver() *zmq.ZmqDriver {
	if container.serviceDriver.zmqDriver == nil {
		container.serviceDriver.zmqDriver = zmq.NewZmqDriver(
			container.infra.ctx,
			container.config.ZMQ,
		)
	}
	return container.serviceDriver.zmqDriver
}
func (container *DIContainer) getMdbDriver() *messagedb.MdbDriver {
	if container.serviceDriver.mdbDriver == nil {
		container.serviceDriver.mdbDriver = messagedb.NewMdbDriver(container.logger)
	}
	return container.serviceDriver.mdbDriver
}
