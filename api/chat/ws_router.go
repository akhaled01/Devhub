package chat

import (
	"RTF/types"
)

var ws_routes = map[string]func(ws *types.User, request string){
	"send_msg":  Send_Message,
	"Open_chat": Open_chat,
}
