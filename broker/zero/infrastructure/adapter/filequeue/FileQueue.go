package filequeue

import "context"

type forSave func(data []byte) error
type forRetrieve func(ctx context.Context) <-chan []byte
type forMessageCount func() int64

func NewFileQueue(
	forSave forSave,
	forRetrieve forRetrieve,
	forMessageCount forMessageCount,
) *FileQueue {
	return &FileQueue{
		forSave,
		forRetrieve,
		forMessageCount,
	}
}

type FileQueue struct {
	save         forSave
	retrieve     forRetrieve
	messageCount forMessageCount
}

func (ouuputAdapter FileQueue) MessageCount() int64 {
	return ouuputAdapter.messageCount()
}
func (outputAdapter FileQueue) SavingMessage(data []byte) error {
	if e := outputAdapter.save(data); e != nil {
		return e
	}
	return nil
}

func (outputAdapter FileQueue) RetrievingMessage(ctx context.Context) <-chan []byte {
	return outputAdapter.retrieve(ctx)
}
