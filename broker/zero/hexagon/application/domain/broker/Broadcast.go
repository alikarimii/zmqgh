package broker

//

import (
	"github.com/alikarimii/zmqph/broker/zero/hexagon/application/domain"
	"github.com/alikarimii/zmqph/pkg/shared"
)

func Broadcast(command domain.BroadcastMessage) shared.RecordedEvents {

	messageId := shared.GenerateMessageID()
	event := domain.BuildMessageBroadcasted(
		messageId,
		command.GetMessage(),
	)

	return shared.RecordedEvents{event}
}
