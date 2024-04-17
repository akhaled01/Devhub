package serializers

import "time"

// Define an interface for message content
type Req_Content interface {
	ContentType() string // Method to get the type of content
}

type Message struct {
	Sender      string    `json:"sender"`
	Recipient   string    `json:"recipient"`
	Timestamp   time.Time `json:"timestamp"`
	Msg_Content string    `json:"msg_content"`
}

type Chat_messages struct {
	Messages []Message `json:"messages"`
}

// Implement ContentType method for TextContent
func (tc Message) ContentType() string {
	return "Message"

}

func (tc Chat_messages) ContentType() string {
	return "Messages"
}

// Message struct with interface field
type WS_Request struct {
	Type    string      `json:"type"`
	Content Req_Content `json:"req_Content"`
}
