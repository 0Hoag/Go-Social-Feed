package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/event"
)

type CommandStartedEvent = event.CommandStartedEvent
type CommandSucceededEvent = event.CommandSucceededEvent
type CommandFailedEvent = event.CommandFailedEvent

type CommandMonitor struct {
	Started   func(ctx context.Context, e *CommandStartedEvent)
	Succeeded func(ctx context.Context, e *CommandSucceededEvent)
	Failure   func(ctx context.Context, e *CommandFailedEvent)
}
