package zmq

import (
	"context"

	"github.com/pebbe/zmq4"
)

type SubZmqCondig struct {
	SUBSCRIBER_REP_ENDPOINT string `envconfig:"SUBSCRIBER_REP_ENDPOINT"`
}

func NewZmqDriver(ctx *zmq4.Context, config SubZmqCondig) *ZmqDriver {
	return &ZmqDriver{ctx, config}
}

// ([REP:40001] broker [REQ:40002])  >-->  ([REP:40002] destination)
type ZmqDriver struct {
	ctx    *zmq4.Context
	config SubZmqCondig
}

func (driver *ZmqDriver) ListenToSource(ctx context.Context, toDB chan<- []byte) chan struct{} {
	terminated := make(chan struct{})
	go func() {
		defer close(terminated)
		// socket not thread safe
		// @TODO find better way
		rep, err := driver.ctx.NewSocket(zmq4.REP)
		if err != nil {
			//
		}
		defer rep.Close()
		if e := rep.Bind(driver.config.SUBSCRIBER_REP_ENDPOINT); e != nil {
			//
		}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				request, _ := rep.RecvBytes(0)
				toDB <- request // send for save
				rep.Send("ok", 0)
			}
		}
	}()
	return terminated
}
