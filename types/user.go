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
}
