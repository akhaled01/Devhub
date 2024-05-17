package types

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	ID        uuid.UUID       `json:"uuid"`
	Username  string          `json:"username"`
	FirstName string          `json:"firstname"`
	LastName  string          `json:"lastname"`
	Email     string          `json:"email"`
	Password  string          `json:"-"`
	Gender    string          `json:"gender"`
	Age       int             `json:"age"`
	Avatar    string          `json:"avatar"`
	Conn      *websocket.Conn `json:"-"` // Store the websocket connection here
}

type PartialUser struct {
	ID       uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Gender   string    `json:"gender"`
}

type Counts struct {
	Number_of_liked_comments int `json:"Number_of_liked_comments"`
	Number_of_liked_posts    int `json:"Number_of_liked_posts"`
	Number_of_comments       int `json:"Number_of_comments"`
	Number_of_posts          int `json:"Number_of_posts"`
}
