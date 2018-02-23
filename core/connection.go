package core

import (
	"log"

	"github.com/gorilla/websocket"
)

// The Connection type represents a websocket connection.
type Connection struct {
	conn    *websocket.Conn
	handler Handler
	sendCh  chan Message
}

func (c *Connection) serve() {
	go c.receive()
	c.send()
	c.handler.OnEvent(EventClosed, c, nil)
}

func (c *Connection) receive() {
	defer func() {
		c.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		m := c.handler.Decode(msg)

		if err != nil {
			log.Printf("ReadMessage error : %v", err)
			break
		}
		c.handler.OnEvent(EventRecv, c, m)
	}
}

func (c *Connection) send() {
	for {
		select {
		case msg, ok := <-c.sendCh:
			if !ok {
				return
			}
			data := c.handler.Encode(msg)
			c.conn.WriteMessage(websocket.TextMessage, data)
			c.handler.OnEvent(EventSend, c, msg)
		}
	}
}

// Send sends a message to the connection.
func (c *Connection) Send(m Message) {
	c.sendCh <- m
}

// Close closes a connection.
func (c *Connection) Close() {
	close(c.sendCh)
	c.conn.Close()
}

// NewConnection creates a websocket connection.
func NewConnection(c *websocket.Conn, h Handler) *Connection {
	return &Connection{
		conn:    c,
		handler: h,
		sendCh:  make(chan Message),
	}
}
