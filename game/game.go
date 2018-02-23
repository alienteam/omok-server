package game

import (
	"log"

	"github.com/alienteam/omok-server/core"
)

// Game is a main server logic.
type Game struct {
	//core.JsonMessageHandler
	core.StringMessageHandler
}

func (s *Game) OnEvent(e core.Event, c *core.Connection, m core.Message) {
	switch e {
	case core.EventConnected:
		log.Println("EVENT_CONN")
		c.Send("hello world~")
	case core.EventRecv:
		log.Printf("EVENT_RECV: %v", m)
	case core.EventSend:
		log.Printf("EVENT_SEND: %v", m)
	case core.EventClosed:
		log.Println("EVENT_CLOSED")
	}
}
