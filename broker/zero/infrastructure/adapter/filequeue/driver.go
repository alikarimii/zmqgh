package filequeue

import (
	"context"

	"github.com/alikarimii/zmqph/pkg/zerologger"
	"github.com/nsqio/go-diskqueue"
)

func NewFileQDriver(
	logger *zerologger.Logger,
	queue diskqueue.Interface) *FileQDriver {

	return &FileQDriver{
		logger,
		queue,
	}
}

type FileQDriver struct {
	logger *zerologger.Logger
	queue  diskqueue.Interface
	// marshal   shared.Marshaler
	// unmarshal shared.UnMarshaler
}

func (q *FileQDriver) MessageCount() int64 {
	return q.queue.Depth()
}
func (q *FileQDriver) Save(data []byte) error {

	if e := q.queue.Put(data); e != nil {
		return e
	}
	return nil
}

func (q *FileQDriver) Retrieve(ctx context.Context) <-chan []byte {
	data := make(chan []byte)
	go func() {
		for {
			select {
			case d := <-q.queue.ReadChan():
				data <- d
			case <-ctx.Done():
				close(data)
				return
			}
		}
	}()
	return data
}
