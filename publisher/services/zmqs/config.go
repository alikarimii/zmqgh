package zmqs

import (
	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/publisher/zero/infrastructure/adapter/zmq"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ZMQ zmq.PubConfig
}

func MustBuildConfigFromEnv(logger *zerologger.Logger) *Config {
	var zmqConf zmq.PubConfig

	mapTo("ZMQ", &zmqConf, logger)

	if zmqConf.SOURCE_REQ_ENDPOINT == "" {
		logger.Panic().Msg("env load faild: zmqConf.SOURCE_REQ_ENDPOINT")
	}
	if zmqConf.REQUEST_RETRIES == 0 {
		logger.Panic().Msg("env load faild: zmqConf.REQUEST_RETRIES")
	}
	if zmqConf.REQUEST_TIMEOUT == 0 {
		logger.Panic().Msg("env load faild: zmqConf.REQUEST_TIMEOUT")
	}

	return &Config{
		ZMQ: zmqConf,
	}
}

// mapTo map section
func mapTo(env string, schema interface{}, logger *zerologger.Logger) {
	e := envconfig.Process(env, schema)
	if e != nil {
		logger.Panic().Msgf("can't load env: %s", e)
	}
}
