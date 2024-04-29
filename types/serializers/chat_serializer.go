package serializers

import (
	"fmt"
	"time"
)

type Message struct {
	Sender      string    `json:"sender"`
	Recipient   string    `json:"recipient"`
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

func (msg Message) To_String() string {
	line_1 := fmt.Sprintf("Sender: %s", msg.Sender)
	line_2 := fmt.Sprintf("Recipient: %s", msg.Recipient)
	line_3 := fmt.Sprintf("Timestamp: %s", msg.Timestamp.String())
	line_4 := fmt.Sprintf("Msg_Content: %s", msg.Msg_Content)

	return fmt.Sprintf("%s\n%s\n%s\n%s", line_1, line_2, line_3, line_4)
}
