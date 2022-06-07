package domain

import (
	"github.com/alikarimii/zmqph/broker/zero/hexagon/application/domain/broker/value"
	"github.com/alikarimii/zmqph/pkg/shared"
)

func BuildMessageRead(
	message value.Message,
) MessageRead {
	event := MessageRead{
		message: message,
	}
	event.meta = shared.BuildEventMeta(event, "testId", 1)
	return event
}

type MessageRead struct {
	message value.Message
	meta    shared.EventMeta
}

func (event MessageRead) Meta() shared.EventMeta {
	return event.meta
}

func (event MessageRead) IsFailureEvent() bool {
	return false
}

func (event MessageRead) FailureReason() error {
	return nil
}
