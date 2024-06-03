package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"RTF/types"
	ser "RTF/types/serializers"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

var mutex sync.Mutex

/* A capsul that holds connections */
type Chat_Server struct {
	conns map[*types.User]bool
}

/* Creats new chat server */
func NewServer() *Chat_Server {
	return &Chat_Server{
		conns: make(map[*types.User]bool),
	}
}

/* Do what you want to do with the connection */
func (s *Chat_Server) HandleWS(
	user *types.User,
	ws_routes map[string]func(user *types.User, request string) error,
) {
	utils.InfoConsoleLog(fmt.Sprint("New connection from client: ", user.Conn.RemoteAddr()))
	ListenerChan <- true
	mutex.Lock()
	defer mutex.Unlock()

	s.conns[user] = true

	go func() {
		defer func() {
			mutex.Lock()
			ListenerChan <- true
			delete(s.conns, user)
			mutex.Unlock()
			user.Conn.Close()
		}()

		msg := &ser.WS_Request{}
		for {
			// Listen for messages (blocking)
			err := user.Conn.ReadJSON(msg)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("read error:", err)
				var session_id uuid.UUID

				for _, s := range types.Sessions {
					if s.User.ID == user.ID {
						session_id = s.SessionID
					}
				}
				types.Sessions[session_id].ChatPartnerID = ""
				break
			}

			// After reading the message, you can choose where to direct it
			passed_content_as_string, _ := json.Marshal(msg.Content)

			if handler, ok := ws_routes[msg.Type]; ok {
				handler(user, string(passed_content_as_string))
			} else {
				utils.ErrorConsoleLog("Handler not found!")
			}
		}
	}()
}
