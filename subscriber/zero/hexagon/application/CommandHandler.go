package application

import "context"

func NewCommandHandler(
	forReadingMessageFromSource ForReadingMessageFromSource,
	forSavingMessage ForSavingMessage,
) *CommandHandler {
	return &CommandHandler{
		forReadingMessageFromSource,
		forSavingMessage,
	}
}

type CommandHandler struct {
	forReadingMessageFromSource ForReadingMessageFromSource
	forSavingMessage            ForSavingMessage
}

// get message from broker
func (c *CommandHandler) GettingMessageProcess(ctx context.Context) {
	data := make(chan []byte)

	go func() {
		<-c.forReadingMessageFromSource(ctx, data)
		// do somting
	}()
	for {
		select {
		case dataByte := <-data:
			if e := c.forSavingMessage(dataByte); e != nil {
				// do somting
			}
		case <-ctx.Done():
			return
		}
	}
}
