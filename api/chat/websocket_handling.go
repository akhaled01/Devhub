package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"RTF/types"
	ser "RTF/types/serializers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	ws_server = types.NewServer()
)

/* Handles the request to connect to chat socket */
func ChatRequestUpgrader(w http.ResponseWriter, r *http.Request) {
	// other_user_id := types.Sessions[uuid.FromStringOrNil(r.PathValue("uid"))]
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go ws_server.HandleWS(conn, ws_routes)
}

func Send_Message(ws *websocket.Conn, request string) {
	message_contents := &ser.Message{}
	json.Unmarshal([]byte(request), message_contents)
	json_msg, _ := json.Marshal(message_contents)

	fmt.Println(string(json_msg))
	ws.WriteMessage(websocket.TextMessage, json_msg)

}

func Open_chat(ws *websocket.Conn, request string) {
	json_msg, _ := json.Marshal(&ser.Message{
		Msg_Content: "Open_chat",
	})

	fmt.Println(request)
	ws.WriteMessage(websocket.TextMessage, json_msg)
}
