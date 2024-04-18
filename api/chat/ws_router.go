package chat

import (
	"github.com/gorilla/websocket"
)

var ws_routes = map[string]func(ws *websocket.Conn, request string){
	"send_msg":  Send_Message,
	"Open_chat": Open_chat,
}
