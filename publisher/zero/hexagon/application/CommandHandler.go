package application

import "context"

func NewCommandHandler(
	forBroadcastMessageToDestination ForBroadcastMessageToDestination,
	forGetGeneratedMessage ForGetGeneratedMessage,
) *CommandHandler {
	return &CommandHandler{
		forBroadcastMessageToDestination,
		forGetGeneratedMessage,
	}
}

type CommandHandler struct {
	forBroadcastMessageToDestination ForBroadcastMessageToDestination
	forGetGeneratedMessage           ForGetGeneratedMessage
}

// send to destination
func (c *CommandHandler) SendingMessageProcess(ctx context.Context) {
	// retrieve message from generator
	data := c.forGetGeneratedMessage(ctx)
	<-c.forBroadcastMessageToDestination(ctx, data)
}
