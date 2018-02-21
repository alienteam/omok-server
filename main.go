package main

import (
	"log"

	"github.com/alienteam/omok-server/core"
)

var server *core.Server

type chat struct {
	//core.JsonMessageHandler
	core.StringMessageHandler
}

func (s *chat) OnEvent(e core.Event, c *core.Connection, m core.Message) {
	switch e {
	case core.EventConnected:
		log.Println("EVENT_CONN", server.Count())
		c.Send("hello world~")
	case core.EventRecv:
		log.Printf("EVENT_RECV: %v", m)
	case core.EventSend:
		log.Printf("EVENT_SEND: %v", m)
	case core.EventClosed:
		log.Println("EVENT_CLOSED")
	}
}

func main() {
	h := &chat{}
	server = core.NewServer("localhost:8080", "/", h)
	server.Start()
}
