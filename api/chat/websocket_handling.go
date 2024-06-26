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

type CurrentStatus struct {
	User      types.User `json:"user"`
	Is_Online bool       `json:"is_online"`
}

var (
	ListenerChan = make(chan bool)
	ws_server    = NewServer()
	upgrader     = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
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
	sessionID, err := uuid.FromString(session_id.Value)
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return
	}
	if sessionID == uuid.Nil {
		types.Sessions[sessionID].ChatPartnerID = ""
	}
	// Validating the session to extract the user
	user_session, err := types.ValidateSession(uuid.FromStringOrNil(session_id.Value))
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
func Send_Message(sender_user *types.User, request string) error {
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
		return err
	}

	// Save the message in the database
	err = chat.SaveChatInDB(*message_contents)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}
	// Get the DMs of the users broadcast them to the sender
	Get_DMs(sender_user, "")

	// Find the user within the connections
	user_idx := 0 // counter for the below loop
	send_to_user := &types.User{}
	for user := range ws_server.conns {
		if user.Username == message_contents.Recipient {
			send_to_user = user
			send_to_conn = user.Conn
			break
		}
		// End of the loop and user wasn't found
		if user_idx == len(ws_server.conns)-1 {
			utils.InfoConsoleLog("user might not be connected!")
			return err
		}
		user_idx++
	}

	// Get the DMs of the users broadcast them to the receiver
	Get_DMs(send_to_user, "")

	// send the message to the correct user via its websocket connection
	err = send_to_conn.WriteMessage(websocket.TextMessage, json_msg)
	if err != nil {
		utils.ErrorConsoleLog("Connection closed!")
		return err
	}

	json_msg, _ = json.Marshal(&ser.WS_Request{
		Type:    "message_success",
		Content: "",
	})

	sender_user.Conn.WriteMessage(websocket.TextMessage, json_msg)

	return nil
}

func Open_chat(user *types.User, request string) error {
	message_contents := &ser.Open_chat_request{}
	json.Unmarshal([]byte(request), message_contents)
	var session_id uuid.UUID

	for _, s := range types.Sessions {
		if s.User.ID == user.ID {
			session_id = s.SessionID
		}
	}

	types.Sessions[session_id].ChatPartnerID = message_contents.User_id

	chat_messages, err := chat.Get_chat(user.Username, message_contents.User_id)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}

	response_capusl := &ser.WS_Request{
		Type:         "open_chat_response",
		Chating_With: types.Sessions[session_id].ChatPartnerID,
		Content:      chat_messages,
	}

	json_msg, _ := json.Marshal(&response_capusl)
	user.Conn.WriteMessage(websocket.TextMessage, json_msg)

	return nil
}

/*
This function listenes for new connections and disconnections
and brodcasts them to all other users for frontend updates
*/
func OnlineListener() {
	for {
		<-ListenerChan
		utils.InfoConsoleLog("conn change detected")
		if err := EvalOnlineUsers(); err != nil {
			utils.ErrorConsoleLog("error brodcasting online users")
			utils.PrintErrorTrace(err)
			return
		}
	}
}

func EvalOnlineUsers() error {
	for user := range ws_server.conns {
		err := Get_DMs(user, "")
		if err != nil {
			utils.ErrorConsoleLog("error getting DMs")
			return err
		}
	}
	return nil
}

func Get_DMs(req_user *types.User, request string) error {
	// get the users that have a message with the user
	DMs, err := chat.Get_Users_By_Last_Message(req_user.Username)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}

	// check if the user has a session
	for idx, u := range DMs {
		if user_session, ok := types.UserHasSessions(u.ID); ok {
			// check if the user is connected
			if _, ok := ws_server.conns[user_session.User]; ok {
				DMs[idx].Is_Online = true
			} else {
				DMs[idx].Is_Online = false
			}
		} else {
			DMs[idx].Is_Online = false
		}
	}

	// get the rest of the users
	all_users, err := user.GetAllUsers()
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}

	for _, u := range all_users {
		// construst a new DM user
		dm_user := ser.DMs_User{
			Username: u.Username,
			ID:       u.ID,
		}

		// check if the user has a session
		if _, ok := types.UserHasSessions(u.ID); ok {
			dm_user.Is_Online = true
		} else {
			dm_user.Is_Online = false
		}

		// check if its already in DMs
		in_DMs := false
		for _, dm := range DMs {
			if dm.Username == u.Username {
				in_DMs = true
				break
			}
		}

		// if its not in DMs, add it
		if !in_DMs {
			DMs = append(DMs, dm_user)
		}
	}

	response_capusl := &ser.WS_Request{
		Type:    "DMs",
		Content: DMs,
	}
	json_msg, _ := json.Marshal(&response_capusl)

	req_user.Conn.WriteMessage(websocket.TextMessage, json_msg)

	return nil
}

// END of function

// This function loads messages in between ids
func Load_Messages(user *types.User, request string) error {
	message_contents := &ser.Load_Messages_Request{}
	json.Unmarshal([]byte(request), message_contents)
	chat_messages, err := chat.Load_Messages(user.Username, message_contents.User_id, message_contents.Begin_id)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}
	response_capusl := &ser.WS_Request{
		Type:    "open_chat_response_from_load_messages",
		Content: chat_messages,
	}

	json_msg, _ := json.Marshal(&response_capusl)
	user.Conn.WriteMessage(websocket.TextMessage, json_msg)

	return nil
}

/*
This is the `typing-in-progress` engine, it sends a signal to the recipient

"start" == user typing

"stop" == user not typing / stopped typing

json sample (sender):

	{
		"type": "typing_in_progress",
		"sender_name": "user_name",
		"signal_type": "start"
	}

json sample (recipient):

	{
		"type": "typing_in_progress",
		"sender": "user_name",
		"is_typing": true
	}
*/
func TIP(user *types.User, request string) error {
	fmt.Println("RECV typing event")
	message_contents := ser.TIP_Request{}

	fmt.Println(request)

	if err := json.Unmarshal([]byte(request), &message_contents); err != nil {
		utils.ErrorConsoleLog(err.Error())
		return err
	}

	is_typing := false

	// fmt.Println(message_contents.Recipient_name)
	// fmt.Println(message_contents.Sender_name)
	// fmt.Println(message_contents.SignalType)

	if recp, ok := types.UserHasSessionByName(message_contents.Recipient_name); ok {
		if message_contents.SignalType == "start" {
			is_typing = true
		}

		fmt.Println("check_U_name")

		recp.User.Conn.WriteJSON(struct {
			Type     string `json:"type"`
			Sender   string `json:"sender"`
			IsTyping bool   `json:"is_typing"`
		}{
			Type:     "typing_in_progress",
			Sender:   user.Username,
			IsTyping: is_typing,
		})
	} else {
		return types.ErrNoConn
	}

	return nil
}
