package core

import (
	"github.com/gorilla/websocket"
	"log"
)

// Connection ...
type Connection struct {
	conn   *websocket.Conn
	proc   *Processor
	sendCh chan Message
}

func (c *Connection) serve() {
	go c.receive()
	c.send()
	c.proc.EventHandler.OnEvent(EventClosed, c, nil)
}

func (c *Connection) receive() {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		m := c.proc.MessageHandler.Decode(msg)

		if err != nil {
			log.Println(err)
			break
		}
		log.Print(msg)
		c.proc.EventHandler.OnEvent(EventRecv, c, m.(map[string]interface{}))
	}
}

func (c *Connection) send() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case _, ok := <-c.sendCh:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			//c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// NewConnection create a websocket connection.
func NewConnection(c *websocket.Conn, p *Processor) *Connection {
	return &Connection{
		conn:   c,
		proc:   p,
		sendCh: make(chan Message),
	}
}
