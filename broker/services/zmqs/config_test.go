package zmqs_test

import (
	"testing"

	"github.com/alikarimii/zmqph/broker/services/zmqs"
	"github.com/alikarimii/zmqph/pkg/zerologger"
)

func TestConfig(t *testing.T) {
	lg := zerologger.NewNilLogger()
	buildConfig := func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("[buildConfig]: %s", err)
			}
		}()
		zmqs.MustBuildConfigFromEnv(lg)
	}
	buildConfig()
}
