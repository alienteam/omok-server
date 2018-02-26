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
	c.conn.Close()
}

func (c *Connection) receive() {
	defer func() {
		close(c.sendCh)
		c.conn.Close()
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("ReadMessage error : %v", err)
			break
		}
		m := c.handler.Decode(msg)
		c.handler.OnEvent(EventRecv, c, m)
	}
}

func (c *Connection) send() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.sendCh:
			if !ok {
				log.Println("SendCh closed")
				return
			}
			data := c.handler.Encode(msg)
			err := c.conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("WriteMessage error: ", err)
				return
			}
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
