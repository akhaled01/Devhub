package types

import (
	"fmt"
	"io"
	"sync"

	"github.com/gorilla/websocket"

	ser "RTF/types/serializers"
)

var mutex sync.Mutex

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

/* Do what you want to do with the connection */
func (s *Chat_Server) HandleWS(
	ws *websocket.Conn,
	ws_routes map[string]func(ws *websocket.Conn, request *ser.WS_Request),
) {
	fmt.Println("new connection from client: ", ws.RemoteAddr())

	mutex.Lock()
	defer mutex.Unlock()

	s.conns[ws] = true
	msg := &ser.WS_Request{}
	for {
		err := ws.ReadJSON(msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)

			continue
		}

		// After reading the message, you can choose where to direct it
		ws_routes[msg.Type](ws, msg)
	}

}
