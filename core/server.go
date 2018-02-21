package core

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Server represents a websocket server.
type Server struct {
	server   *http.Server
	path     string
	addr     string
	upgrader websocket.Upgrader
	cm       *ConnectionManager
	handler  Handler
}

// NewServer creates a websocket server but not start to accept.
func NewServer(addr, path string, h Handler) *Server {
	s := &Server{
		server: &http.Server{
			Addr:         addr,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		path: path,
		addr: addr,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		cm: &ConnectionManager{
			conns: make(map[*Connection]bool),
		},
		handler: h,
	}
	s.server.Handler = s
	return s
}

// Start runs the server
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

// Count returns the concurrency connections
func (s *Server) Count() int {
	return s.cm.count()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	go s.handleConnect(conn)
}

func (s *Server) handleConnect(wc *websocket.Conn) {

	c := NewConnection(wc, s.handler)
	s.cm.add(c)

	defer func() {
		s.cm.remove(c)
	}()

	go c.handler.OnEvent(EventConnected, c, nil)
	c.serve()
}
