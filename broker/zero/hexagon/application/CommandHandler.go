package application

import (
	"context"
)

func NewCommandHandler(
	forBroadcastMessageToDestination ForBroadcastMessageToDestination,
	forReadingMessageFromSource ForReadingMessageFromSource,
	forSavingMessage ForSavingMessage,
	forRetrievingMessage ForRetrievingMessage,
) *CommandHandler {
	return &CommandHandler{
		forBroadcastMessageToDestination,
		forReadingMessageFromSource,
		forSavingMessage,
		forRetrievingMessage,
	}
}

type CommandHandler struct {
	// out port inject here like port to db(driven adapter) adapter (ForRetrievingEventStream)
	forBroadcastMessageToDestination ForBroadcastMessageToDestination
	forReadingMessageFromSource      ForReadingMessageFromSource
	// file queue
	forSavingMessage     ForSavingMessage
	forRetrievingMessage ForRetrievingMessage
}

// get message from source
func (c *CommandHandler) GettingMessageProcess(ctx context.Context) {
	// wrapWithMsg := "BrokerCommandHandler.GettingMessage"
	data := make(chan []byte)

	go func() {
		<-c.forReadingMessageFromSource(ctx, data)
		// do somting
	}()
	for {
		select {
		case dataByte := <-data:
			// do serialize or validate with value
			// save message to broker
			if e := c.forSavingMessage(dataByte); e != nil {
				// do somting
			}
		case <-ctx.Done():
			return
		}
	}
}

// send to destination
func (c *CommandHandler) SendingMessageProcess(ctx context.Context) {
	// wrapWithMsg := "BrokerCommandHandler.SendingMessage"
	// @TODO error handling
	// @TODO can send data to domain for validate
	// @TODO can stop/start flow data manualy with domain

	// retrieve message from broker
	data := c.forRetrievingMessage(ctx)
	<-c.forBroadcastMessageToDestination(ctx, data)
	//
}
