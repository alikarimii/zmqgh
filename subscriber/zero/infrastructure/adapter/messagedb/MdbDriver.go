package messagedb

import (
	"fmt"

	"github.com/alikarimii/zmqph/pkg/zerologger"
)

func NewMdbDriver(
	logger *zerologger.Logger) *MdbDriver {

	return &MdbDriver{
		logger,
	}
}

type MdbDriver struct {
	logger *zerologger.Logger
}

func (q *MdbDriver) MessageCount() int64 {
	return 0
}
func (q *MdbDriver) Save(data []byte) error {
	// @TODO save to file
	fmt.Println("%s", data)
	return nil
}
