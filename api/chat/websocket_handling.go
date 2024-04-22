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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Extract session ID from query parameters
	user_uuid := r.URL.Query().Get("user_uuid")
	fmt.Printf("WebSocket connection established with session ID: %s\n", user_uuid)

	// Inserting the connection into session (A hub of connected clients)
	user_session, err := types.ValidateSession(uuid.FromStringOrNil(user_uuid))
	fmt.Println(user_session)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return
	}
	user_session.Conn = conn

	ws_server.HandleWS(conn, ws_routes)
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

	session_rec := types.Sessions[message_contents.Recipient]

	err := session_rec.Conn.WriteMessage(websocket.TextMessage, json_msg)
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
