package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"

	"RTF/storage/interfaces/chat"
	"RTF/storage/interfaces/user"
	"RTF/types"
	ser "RTF/types/serializers"
	"RTF/utils"
)

type ConnEvent struct {
	Name  string `json:"name"`
	Event string `json:"event"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var listenerChan = make(chan ConnEvent)

var (
	ws_server = NewServer()
)

/* Handles the request to connect to chat socket */
func ChatRequestUpgrader(w http.ResponseWriter, r *http.Request) {
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

	// Validating the session to extract the user
	user_session, err := types.ValidateSession(uuid.FromStringOrNil(session_id.Value))
	fmt.Println(user_session)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		conn.Close()
		return
	}

	// Give the user its connection
	user_session.User.Conn = conn

	ws_server.HandleWS(user_session.User, ws_routes)
}

/*
	  A function to send message
	  Inputs:

		ws: connection to send the message to
		request: the message
*/
func Send_Message(sender_user *types.User, request string) {
	message_contents := &ser.Message{}
	json.Unmarshal([]byte(request), message_contents)
	message_contents.Sender = sender_user.Username // Put the username in the message capsul
	message_contents.Timestamp = time.Now()

	// Encapsulate the response
	response_capusl := &ser.WS_Request{
		Type:    "message",
		Content: message_contents,
	}
	json_msg, _ := json.Marshal(response_capusl)

	var send_to_conn *websocket.Conn

	_, err := user.GetSingleUser("user_name", message_contents.Recipient)
	if err != nil {
		utils.ErrorConsoleLog("User can't be found in the database!")
		return
	}

	// Save the message in the database
	err = chat.SaveChatInDB(*message_contents)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
	}

	// Find the user within the connections
	user_idx := 0 // counter for the below loop
	for user := range ws_server.conns {
		if user.Username == message_contents.Recipient {
			send_to_conn = user.Conn
			break
		}
		// End of the loop and user wasn't found
		if user_idx == len(ws_server.conns)-1 {
			utils.InfoConsoleLog("user might not be connected!")
			return
		}
		user_idx++
	}

	// send the message to the correct user via its websocket connection
	err = send_to_conn.WriteMessage(websocket.TextMessage, json_msg)
	if err != nil {
		utils.ErrorConsoleLog("Connection closed!")
		return
	}

	sender_user.Conn.WriteMessage(websocket.TextMessage, []byte("Message sent!"))

}

func Open_chat(user *types.User, request string) {
	message_contents := &ser.Open_chat_request{}
	json.Unmarshal([]byte(request), message_contents)

	chat_messages, err := chat.Get_chat(user.Username, message_contents.User_id)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return
	}

	response_capusl := &ser.WS_Request{
		Type:    "open_chat_response",
		Content: chat_messages,
	}

	json_msg, _ := json.Marshal(&response_capusl)
	user.Conn.WriteMessage(websocket.TextMessage, json_msg)
}

/*
This function listenes for new connections and disconnections
and brodcasts them to all other users for frontend updates
*/
func OnlineListener() {
	for {
		event := <-listenerChan
		Brodcast(event)
	}
}

/*
Broadcasts an event to all connected users
an event is basically when a user connects / disconnects
*/
func Brodcast(event ConnEvent) {
	jsonMsg, _ := json.Marshal(&ser.WS_Request{
		Type:    "conn_event",
		Content: event,
	})
	for conn := range ws_server.conns {
		conn.Conn.WriteMessage(websocket.TextMessage, jsonMsg)
	}
}
