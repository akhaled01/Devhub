package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var mutex sync.Mutex

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/* A capsul that holds connections */
type Chat_Server struct {
	conns map[*websocket.Conn]bool
}

/* Creats new chat server */
func NewServer() *Chat_Server {
	return &Chat_Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

/* Reads the incoming message repeatadly */
func (s *Chat_Server) readLoop(ws *websocket.Conn) {

	msg := &Message{}
	for {
		err := ws.ReadJSON(msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		json_msg, _ := json.Marshal(&Message{
			Msg: "yes!",
		})
		fmt.Println(msg)
		ws.WriteJSON(string(json_msg))
	}
}

/* Do what you want to do with the connection */
func (s *Chat_Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new connection from client: ", ws.RemoteAddr())

	mutex.Lock()
	defer mutex.Unlock()

	s.conns[ws] = true
	s.readLoop(ws)
}

var (
	ws_server = NewServer()
)

/* Handles the request to connect to chat socket */
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ws_server.handleWS(conn)
}
