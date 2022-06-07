package bytegen

import (
	"context"
	"math/rand"

	"github.com/alikarimii/zmqph/pkg/zerologger"
)

func NewGenerator(
	logger *zerologger.Logger) *Generator {

	return &Generator{
		logger: logger,
	}
}

type Generator struct {
	logger *zerologger.Logger
	count  int64 // in memory counter
}

func (q *Generator) MessageCount() int64 {
	return q.count
}

func (q *Generator) Generate(ctx context.Context) <-chan []byte {
	data := make(chan []byte)
	go func() {
		for {
			bt := make([]byte, GenerateRandomSize())
			_, er := rand.Read(bt)
			if er != nil {
				continue
			}
			select {
			case <-ctx.Done():
				close(data)
				return
			default:
				data <- bt
				q.count++
			}
		}
	}()
	return data
}
