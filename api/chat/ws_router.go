package chat

import (
	"RTF/types"
)

var ws_routes = map[string]func(ws *types.User, request string) error{
	"send_msg":      Send_Message,
	"Open_chat":     Open_chat,
	"get_dms":       Get_DMs,
	"load_messages": Load_Messages,
}
