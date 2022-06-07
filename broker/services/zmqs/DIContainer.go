package zmqs

import (
	"fmt"

	"github.com/alikarimii/zmqph/broker/zero/hexagon/application"
	"github.com/alikarimii/zmqph/broker/zero/infrastructure/adapter/filequeue"
	"github.com/alikarimii/zmqph/broker/zero/infrastructure/adapter/zmq"
	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/cockroachdb/errors"
	"github.com/nsqio/go-diskqueue"
	"github.com/pebbe/zmq4"
)

type DIOption func(container *DIContainer) error

func WithZmqSockets(ctx *zmq4.Context) DIOption {
	return func(container *DIContainer) error {
		if ctx == nil {
			return errors.New("must init all socket first")
		}
		container.infra.ctx = ctx
		return nil
	}
}
func WithDiskQueue(queue diskqueue.Interface) DIOption {
	return func(container *DIContainer) error {
		if queue == nil {
			return errors.New("file queue must be init")
		}
		container.infra.queue = queue
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

func MustBuildDiskQueue(config *Config, logger *zerologger.Logger) diskqueue.Interface {
	fn := func(lvl diskqueue.LogLevel, f string, args ...interface{}) {
		logger.Log().Msg(fmt.Sprintf(lvl.String()+": "+f, args...))
	}
	return diskqueue.New(
		"zero_disk_queue",
		config.FileQueue.DBDir,
		config.FileQueue.MaxBytesPerFile,
		config.FileQueue.MinMsgSize,
		config.FileQueue.MaxMsgSize,
		config.FileQueue.SyncEvery,
		config.FileQueue.SyncTimeout, fn)
}

type DIContainer struct {
	config *Config
	logger *zerologger.Logger
	infra  struct {
		queue diskqueue.Interface
		ctx   *zmq4.Context
	}
	service struct {
		commandHandler *application.CommandHandler
		queryHandler   *application.QueryHanndler
		fileQ          *filequeue.FileQueue
		zmq            *zmq.ZMQ
	}
	serviceDriver struct {
		fileQDriver *filequeue.FileQDriver
		zmqDriver   *zmq.ZmqDriver
	}
}

func (container *DIContainer) init() {
	_ = container.GetCommandHandler()
	_ = container.GetQueryHandler()
	_ = container.GetFileQ()
	_ = container.GetZMQ()
}

func (container *DIContainer) GetCommandHandler() *application.CommandHandler {
	if container.service.commandHandler == nil {
		container.service.commandHandler = application.NewCommandHandler(
			container.GetZMQ().BroadcastMessageToDestination,
			container.GetZMQ().ReadingMessageFromSource,
			container.GetFileQ().SavingMessage,
			container.GetFileQ().RetrievingMessage,
		)
	}
	return container.service.commandHandler
}
func (container *DIContainer) GetQueryHandler() *application.QueryHanndler {
	if container.service.queryHandler == nil {
		container.service.queryHandler = application.NewQueryHandler(
			container.GetFileQ().MessageCount,
		)
	}
	return container.service.queryHandler
}

func (container *DIContainer) GetFileQ() *filequeue.FileQueue {
	if container.service.fileQ == nil {
		container.service.fileQ = filequeue.NewFileQueue(
			container.getFileQDriver().Save,
			container.getFileQDriver().Retrieve,
			container.getFileQDriver().MessageCount,
		)
	}
	return container.service.fileQ
}
func (container *DIContainer) GetZMQ() *zmq.ZMQ {
	if container.service.zmq == nil {
		container.service.zmq = zmq.NewZMQ(
			container.getZMQDriver().ListenToSource,
			container.getZMQDriver().ListenToDestination,
		)
	}
	return container.service.zmq
}
func (container *DIContainer) getFileQDriver() *filequeue.FileQDriver {
	if container.serviceDriver.fileQDriver == nil {
		container.serviceDriver.fileQDriver = filequeue.NewFileQDriver(
			container.logger,
			container.infra.queue,
		)
	}
	return container.serviceDriver.fileQDriver
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

// func (container *DIContainer) () {}
