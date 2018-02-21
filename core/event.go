package core

// Event is connection event type.
type Event int

const (
	// EventConnected means server accept a new connection.
	EventConnected Event = iota
	// EventSend means conn send a packet.
	EventSend
	// EventRecv means conn recv a packet.
	EventRecv
	// EventClosed means conn is closed.
	EventClosed
)

// EventHandler is the event callback handler.
type EventHandler interface {
	OnEvent(e Event, c *Connection, m Message)
}
