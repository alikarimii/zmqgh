package domain

import "github.com/alikarimii/zmqph/broker/zero/hexagon/application/domain/broker/value"

func BuildReadMessage(message value.Message) ReadMessage {
	command := ReadMessage{
		message,
	}
	return command

}

type ReadMessage struct {
	message value.Message
}

func (command ReadMessage) GetMessage() value.Message {
	return command.message
}
