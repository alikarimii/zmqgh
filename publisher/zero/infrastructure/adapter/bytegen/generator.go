package bytegen

import "context"

type forGenerate func(ctx context.Context) <-chan []byte
type forMessageCount func() int64

func NewByteGenerator(
	forGenerate forGenerate,
	forMessageCount forMessageCount,
) *ByteGenerator {
	return &ByteGenerator{
		forGenerate,
		forMessageCount,
	}
}

type ByteGenerator struct {
	generate     forGenerate
	messageCount forMessageCount
}

func (ouuputAdapter ByteGenerator) MessageCount() int64 {
	return ouuputAdapter.messageCount()
}

func (outputAdapter ByteGenerator) GetGeneratedMessage(ctx context.Context) <-chan []byte {
	return outputAdapter.generate(ctx)
}
