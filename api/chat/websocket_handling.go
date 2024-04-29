package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"

	"RTF/types"
	ser "RTF/types/serializers"
	"RTF/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	ws_server = NewServer()
)

/* Handles the request to connect to chat socket */
func ChatRequestUpgrader(w http.ResponseWriter, r *http.Request) {
	// other_user_id := types.Sessions[uuid.FromStringOrNil(r.PathValue("uid"))]

	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog("Invalid session")
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("WebSocket connection established with session ID: %s\n", session_id.Value)

	// Inserting the connection into session (A hub of connected clients)
	user_session, err := types.ValidateSession(uuid.FromStringOrNil(session_id.Value))
	fmt.Println(user_session)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		conn.Close()
		return
	}
	user_session.User.Conn = conn

	ws_server.HandleWS(user_session.User, ws_routes)
}

/*
	  A function to send message
	  Inputs:

		ws: connection to send the message to
		request: the message
*/
func Send_Message(ws *websocket.Conn, request string) {
	message_contents := &ser.Message{}
	json.Unmarshal([]byte(request), message_contents)
	json_msg, _ := json.Marshal(message_contents)

	var send_to_conn *websocket.Conn
	for user := range ws_server.conns {
		if user.Username == message_contents.Recipient {
			send_to_conn = user.Conn
			break
		}
		utils.InfoConsoleLog("username wasn't found!")
		return
	}

	err := send_to_conn.WriteMessage(websocket.TextMessage, json_msg)
	if err != nil {
		utils.ErrorConsoleLog("Connection closed!")
		return
	}
	ws.WriteMessage(websocket.TextMessage, []byte("Message sent!"))

}

func Open_chat(ws *websocket.Conn, request string) {
	json_msg, _ := json.Marshal(&ser.Message{
		Msg_Content: "Open_chat",
	})

	fmt.Println(request)
	ws.WriteMessage(websocket.TextMessage, json_msg)
}
