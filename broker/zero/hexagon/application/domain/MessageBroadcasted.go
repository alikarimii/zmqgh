package domain

import "github.com/alikarimii/zmqph/pkg/shared"

func BuildMessageBroadcasted(
	messageID shared.MessageID,
	data []byte,
) MessageBroadcasted {
	event := MessageBroadcasted{
		messageID: messageID,
	}
	event.meta = shared.BuildEventMeta(event, "testId", 1)
	return event
}

type MessageBroadcasted struct {
	messageID shared.MessageID
	meta      shared.EventMeta
}

func (event MessageBroadcasted) Meta() shared.EventMeta {
	return event.meta
}

func (event MessageBroadcasted) IsFailureEvent() bool {
	return false
}

func (event MessageBroadcasted) FailureReason() error {
	return nil
}
