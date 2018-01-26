package chat

import (
	"time"
	"log"
	"net/http"
	"context"
	"github.com/gorilla/websocket"
)

type ChatServer struct {
	Server *http.Server
	Clients map[int]*Client
	Path string
	Port string
	Num int
	AuthFunc func(id, pw string) (no int)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewChatServer(port, path string, authFunc func(id, pw string) (no int)) *ChatServer{
	s := &ChatServer{
		Server: &http.Server{
			Addr:port,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Clients: make(map[int]*Client),
		Path: path,
		Port: port,
		Num: 0,
		AuthFunc: authFunc,
	}
	
	s.Server.Handler = s

	return s
}

func (s *ChatServer)broadcast(id int, msg []byte) {
    for _, c := range s.Clients {
		if c.id != id {
			c.send <- msg
		}
    }
}

func (s *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	num := s.AuthFunc("njh0906", "aaa")
	client := &Client{id: num,conn:conn, send:make(chan []byte, 256), server:s}
	s.Clients[num] = client;
	s.Num = num
	
	log.Println("Connection accepted: ", num, conn.RemoteAddr())
	go client.read()
	go client.write()
}

func (s *ChatServer)Start(){
	log.Println("Start server")
	err := s.Server.ListenAndServe()
	log.Println(err)
}

func (s *ChatServer)Close(){
	log.Println("Stop server")
	log.Println(s.Server.Close())
}

func (s *ChatServer)Shutdown(){
	err := s.Server.Shutdown(context.Background())
	log.Println(err)
}
