package serializers

import (
	"fmt"
	"time"
)

type Message struct {
	Id          string    `json:"id"`
	Sender      string    `json:"sender"`
	Recipient   string    `json:"recipient"`
	Timestamp   time.Time `json:"timestamp"`
	Msg_Content string    `json:"msg_content"`
	Msg_Status  bool      `json:"msg_status"`
}

// Message struct with interface field
type WS_Request struct {
	Type         string      `json:"type"`
	Content      interface{} `json:"req_Content"`
	Chating_With string      `json:"string"`
}

func (msg Message) To_String() string {
	line_1 := fmt.Sprintf("Sender: %s", msg.Sender)
	line_2 := fmt.Sprintf("Recipient: %s", msg.Recipient)
	line_3 := fmt.Sprintf("Timestamp: %s", msg.Timestamp.String())
	line_4 := fmt.Sprintf("Msg_Content: %s", msg.Msg_Content)

	return fmt.Sprintf("%s\n%s\n%s\n%s", line_1, line_2, line_3, line_4)
}

type Open_chat_request struct {
	User_id string `json:"user_id"`
}

type Load_Messages_Request struct {
	User_id  string `json:"user_id"`
	Begin_id int    `json:"begin_id"`
}

type TIP_Request struct {
	Sender_name    string `json:"sender_name"`
	Recipient_name string `json:"recipient_name"`
	SignalType     string `json:"signal_type"` // can be "start" or "stop"
}
