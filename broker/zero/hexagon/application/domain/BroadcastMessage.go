package domain

import "github.com/alikarimii/zmqph/broker/zero/hexagon/application/domain/broker/value"

func BuildBroadcastMessage(
	message []byte,
	qname value.QueueName,
) BroadcastMessage {
	command := BroadcastMessage{
		message,
		qname,
	}
	return command
}

type BroadcastMessage struct {
	message []byte
	qname   value.QueueName
}

func (command BroadcastMessage) GetQueueName() value.QueueName {
	return command.qname
}

func (command BroadcastMessage) GetMessage() []byte {
	return command.message
}
