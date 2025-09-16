package mongo

import (
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClientOptions is a wrapper of options.ClientOptions.
type ClientOptions struct {
	clo *options.ClientOptions
}

// NewClientOptions creates a new ClientOptions instance.
func NewClientOptions() ClientOptions {
	return ClientOptions{
		clo: options.Client(),
	}
}

// ApplyURI adds an option to specify the URI of the MongoDB deployment to connect to.
func (co ClientOptions) ApplyURI(uri string) ClientOptions {
	co.clo.ApplyURI(uri)

	return co
}

// SetMonitor adds an option to specify the monitor to receive command monitoring events.
func (co ClientOptions) SetMonitor(m CommandMonitor) ClientOptions {
	co.clo.SetMonitor(&event.CommandMonitor{
		Started:   m.Started,
		Succeeded: m.Succeeded,
		Failed:    m.Failure,
	})

	return co
}
