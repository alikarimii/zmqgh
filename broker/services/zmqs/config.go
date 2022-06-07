package zmqs

import (
	"time"

	"github.com/alikarimii/zmqph/broker/zero/infrastructure/adapter/zmq"

	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/kelseyhightower/envconfig"
)

type filequeueConf struct {
	DBDir           string        `envconfig:"FILEQUEUE_DB_DIR"`
	MaxBytesPerFile int64         `envconfig:"FILEQUEUE_MAX_BYTES_PER_FILE"`
	MinMsgSize      int32         `envconfig:"FILEQUEUE_MIN_MSG_SIZE"`
	MaxMsgSize      int32         `envconfig:"FILEQUEUE_MAX_MSG_SIZE"`
	SyncEvery       int64         `envconfig:"FILEQUEUE_SYNC_EVERY"`
	SyncTimeout     time.Duration `envconfig:"FILEQUEUE_SYNC_TIMEOUT"`
}

type Config struct {
	ZMQ       zmq.BrokerZmqCondig
	FileQueue filequeueConf
}

func MustBuildConfigFromEnv(logger *zerologger.Logger) *Config {
	var zmqConf zmq.BrokerZmqCondig
	var fileqConf filequeueConf

	mapTo("ZMQ", &zmqConf, logger)
	mapTo("FILEQUEUE", &fileqConf, logger)
	if zmqConf.BROKER_REP_ENDPOINT == "" {
		logger.Panic().Msg("env load faild: zmqConf.BROKER_REP_ENDPOINT")
	}
	if zmqConf.BROKER_DESTINATION_REQ_ENDPOINT == "" {
		logger.Panic().Msg("env load faild: zmqConf.BROKER_DESTINATION_REQ_ENDPOINT")
	}
	if zmqConf.REQUEST_RETRIES == 0 {
		logger.Panic().Msg("env load faild: zmqConf.REQUEST_RETRIES")
	}
	if fileqConf.MaxBytesPerFile == 0 {
		logger.Panic().Msg("env load faild: fileqConf.MaxBytesPerFile")
	}
	if fileqConf.MaxMsgSize == 0 {
		logger.Panic().Msg("env load faild: fileqConf.MaxMsgSize")
	}
	if fileqConf.MinMsgSize == 0 {
		logger.Panic().Msg("env load faild: fileqConf.MinMsgSize")
	}
	if fileqConf.SyncEvery == 0 {
		logger.Panic().Msg("env load faild: fileqConf.SyncEvery")
	}
	if fileqConf.SyncTimeout == 0 {
		logger.Panic().Msg("env load faild: fileqConf.SyncTimeout")
	}
	if fileqConf.DBDir == "" {
		logger.Panic().Msg("env load faild: fileqConf.DBDir")
	}
	return &Config{
		ZMQ:       zmqConf,
		FileQueue: fileqConf,
	}
}

// mapTo map section
func mapTo(env string, schema interface{}, logger *zerologger.Logger) {
	e := envconfig.Process(env, schema)
	if e != nil {
		logger.Panic().Msgf("can't load env: %s", e)
	}
}
