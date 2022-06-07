package broker

import (
	"github.com/alikarimii/zmqph/broker/zero/hexagon/application/domain"
	"github.com/alikarimii/zmqph/pkg/shared"
)

// read from broker driver

func Read(command domain.ReadMessage) shared.RecordedEvents {
	// @TODO
	event := domain.BuildMessageRead(
		command.GetMessage(),
	)

	return shared.RecordedEvents{event}
}
