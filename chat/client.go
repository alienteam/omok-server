package chat

import (
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
    id int
	conn *websocket.Conn
    send chan []byte
    server *ChatServer
}

func (c *Client) read() {
	defer func() {
        c.conn.Close()
        delete(c.server.Clients, c.id)
    }()

    for {
       msgType, msg, err := c.conn.ReadMessage()
       if err != nil {
           log.Println(err)
           break
       }
       log.Printf("%d: %d, %s", c.id, msgType, msg)
       c.server.broadcast(c.id, msg)
    }
}

func (c *Client) write() {
	defer func(){
        c.conn.Close()
    }()

    for{
        select{
        case message, ok := <-c.send:
            if !ok {
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            c.conn.WriteMessage(websocket.TextMessage, message)
        }
    }
}