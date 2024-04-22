package serializers

import (
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	Sender      string    `json:"sender"`
	Recipient   uuid.UUID `json:"recipient"`
	Timestamp   time.Time `json:"timestamp"`
	Msg_Content string    `json:"msg_content"`
}

type Chat_messages struct {
	Messages []Message `json:"messages"`
}

// Message struct with interface field
type WS_Request struct {
	Type    string      `json:"type"`
	Content interface{} `json:"req_Content"`
}
