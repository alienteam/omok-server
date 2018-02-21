package core

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Server :
type Server struct {
	server   *http.Server
	path     string
	port     string
	upgrader websocket.Upgrader
	cm       *ConnectionManager
	proc     *Processor
}

// NewServer create a websocket server but not start to accept.
func NewServer(port, path string) *Server {
	s := &Server{
		server: &http.Server{
			Addr:         port,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		path: path,
		port: port,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		cm: &ConnectionManager{
			conns: make(map[*Connection]bool),
		},
	}
	s.server.Handler = s
	return s
}

// Start runs server
func (s *Server) Start() {
	s.server.ListenAndServe()
}

// Close immediately closes all listeneres and connections
func (s *Server) Close() {
	s.server.Close()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() {
	s.server.Shutdown(context.Background())
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	go s.handleConnect(conn)
}

func (s *Server) handleConnect(wc *websocket.Conn) {
	c := NewConnection(wc, s.proc)
	s.cm.add(c)
	c.proc.EventHandler.OnEvent(EventConnected, c, nil)
	c.serve()
}
