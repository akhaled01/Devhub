package chat

import (
	"github.com/gorilla/websocket"

	ser "RTF/types/serializers"
)

var ws_routes = map[string]func(ws *websocket.Conn, request *ser.WS_Request){
	"send_msg":  Send_Message,
	"Open_chat": Open_chat,
}
