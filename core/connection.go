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
	closeCh chan struct{}
}

func (c *Connection) serve() {
	go c.receive()
	c.send()
	c.handler.OnEvent(EventClosed, c, nil)
}

func (c *Connection) receive() {
	defer func() {
		c.conn.Close()
		close(c.closeCh)
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		m := c.handler.Decode(msg)

		if err != nil {
			log.Println(err)
			c.conn.Close()
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
		case <-c.closeCh:
			return
		}
	}
}

// Send sends a message to the connection.
func (c *Connection) Send(m Message) {
	c.sendCh <- m
}

// NewConnection creates a websocket connection.
func NewConnection(c *websocket.Conn, h Handler) *Connection {
	return &Connection{
		conn:    c,
		handler: h,
		sendCh:  make(chan Message),
		closeCh: make(chan struct{}),
	}
}
