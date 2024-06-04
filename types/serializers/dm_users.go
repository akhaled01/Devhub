package serializers

import "github.com/gofrs/uuid"

type DMs_User struct {
	ID              uuid.UUID `json:"id"`
	Username        string    `json:"username"`
	LastMessageTime string    `json:"last_message_time"`
	Is_Online       bool      `json:"is_online"`
	Msg_sender      string    `json:"msg_sender"`
	Msg_Status      bool      `json:"msg_status"`
}
