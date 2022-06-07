package zmqs

import (
	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/subscriber/zero/infrastructure/adapter/zmq"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ZMQ zmq.SubZmqCondig
}

func MustBuildConfigFromEnv(logger *zerologger.Logger) *Config {
	var zmqConf zmq.SubZmqCondig

	mapTo("ZMQ", &zmqConf, logger)

	if zmqConf.SUBSCRIBER_REP_ENDPOINT == "" {
		logger.Panic().Msg("env load faild: zmqConf.SUBSCRIBER_REP_ENDPOINT")
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
