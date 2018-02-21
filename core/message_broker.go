package core

// MessageBroker ...
type MessageBroker interface {
	Publish()
	Subscribe()
}
