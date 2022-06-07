package messagedb

type forSave func(data []byte) error
type forMessageCount func() int64

func NewMessagedb(
	forSave forSave,
	forMessageCount forMessageCount,
) *Messagedb {
	return &Messagedb{
		forSave,
		forMessageCount,
	}
}

type Messagedb struct {
	save         forSave
	messageCount forMessageCount
}

func (ouuputAdapter Messagedb) MessageCount() int64 {
	return ouuputAdapter.messageCount()
}
func (outputAdapter Messagedb) SavingMessage(data []byte) error {
	if e := outputAdapter.save(data); e != nil {
		return e
	}
	return nil
}
