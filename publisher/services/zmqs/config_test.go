package zmqs_test

import (
	"testing"

	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/alikarimii/zmqph/publisher/services/zmqs"
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
