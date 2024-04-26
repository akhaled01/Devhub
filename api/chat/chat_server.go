package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/gorilla/websocket"

	ser "RTF/types/serializers"
	"RTF/utils"
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
	ws_routes map[string]func(ws *websocket.Conn, request string),
) {
	utils.InfoConsoleLog(fmt.Sprint("New connection from client: ", ws.RemoteAddr()))
	mutex.Lock()
	defer mutex.Unlock()

	s.conns[ws] = true

	go func() {
		defer func() {
			mutex.Lock()
			delete(s.conns, ws)
			mutex.Unlock()
			ws.Close()
		}()

		msg := &ser.WS_Request{}
		for {
			err := ws.ReadJSON(msg)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("read error:", err)
				break
			}

			// After reading the message, you can choose where to direct it
			passed_content_as_string, _ := json.Marshal(msg.Content)

			if handler, ok := ws_routes[msg.Type]; ok {
				handler(ws, string(passed_content_as_string))
			} else {
				utils.ErrorConsoleLog("Handler not found!")
			}
		}
	}()
}
