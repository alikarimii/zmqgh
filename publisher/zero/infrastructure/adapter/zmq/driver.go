package zmq

import (
	"context"
	"fmt"
	"time"

	"github.com/pebbe/zmq4"
)

type PubConfig struct {
	REQUEST_TIMEOUT     time.Duration `envconfig:"REQUEST_TIMEOUT"`
	REQUEST_RETRIES     int           `envconfig:"REQUEST_RETRIES"`
	SOURCE_REQ_ENDPOINT string        `envconfig:"SOURCE_REQ_ENDPOINT"`
}

func NewZmqDriver(config PubConfig, ctx *zmq4.Context) *ZmqDriver {
	return &ZmqDriver{config, ctx}
}

// (source [REQ:40001])  >-->  ([REP:40001] broker)
type ZmqDriver struct {
	config PubConfig
	ctx    *zmq4.Context
}

// https://zguide.zeromq.org/docs/chapter4/
// we must implement one of this pattern base on system situation
// this is Lazy Pirate
func (driver *ZmqDriver) ListenToDestination(ctx context.Context, fromGenerator <-chan []byte) chan struct{} {
	terminated := make(chan struct{})
	go func() {
		defer close(terminated)
		// socket not thread safe
		// @TODO find better way
		sok, _ := driver.ctx.NewSocket(zmq4.REQ)
		sok.Connect(driver.config.SOURCE_REQ_ENDPOINT)
		for {
			select {
			case dataForSend := <-fromGenerator:
				// @TODO use heartbeat
				for sequence, retriesLeft := 1, driver.config.REQUEST_RETRIES; retriesLeft > 0; sequence++ {
					fmt.Printf("I: Sending (%d)\n", sequence)
					sok.SendBytes(dataForSend, 0)
					for expectReply := true; expectReply; {
						//  Poll socket for a reply, with timeout
						poller := zmq4.NewPoller()
						poller.Add(sok, zmq4.POLLIN)
						sockets, err := poller.Poll(driver.config.REQUEST_TIMEOUT)
						if err != nil {
							panic(err) //  Interrupted
						}
						//  .split process server reply
						//  Here we process a server reply and exit our loop if the
						//  reply is valid. If we didn't a reply we close the client
						//  socket and resend the request. We try a number of times
						//  before finally abandoning:
						if item := sockets[0]; item.Events&zmq4.POLLIN != 0 {
							//  We got a reply from the server
							reply, err := item.Socket.Recv(0)
							if err != nil {
								panic(err) //  Interrupted
							}

							if reply == "ok" { // @TODO proper condition
								fmt.Printf("I: Server replied (%s)\n", reply)
								retriesLeft = 0
								expectReply = false
							} else {
								fmt.Printf("E: Malformed reply from server: %s", reply)
							}
						} else if retriesLeft--; retriesLeft == 0 {
							fmt.Println("E: Server seems to be offline, abandoning")
							sok.SetLinger(0)
							sok.Close()
							break
						} else {
							fmt.Println("W: No response from server, retrying...")
							//  Old socket is confused; close it and open a new one
							sok.SetLinger(0)
							sok.Close()
							// @TODO check thread safety
							// use mutex if needed
							sok, _ = zmq4.NewSocket(zmq4.REQ)
							sok.Connect(driver.config.SOURCE_REQ_ENDPOINT)
							fmt.Printf("I: Resending (%d)\n", sequence)
							//  Send request again, on new socket
							sok.SendBytes(dataForSend, 0)
						}
					}
				}
			case <-ctx.Done():
				return
				// default: @TODO check block select or not
			}
		}

		// @TODO if terminate, read next item from fileQ
		// so put proper REQUEST_RETRIES and REQUEST_TIMEOUT

	}()
	return terminated
}
