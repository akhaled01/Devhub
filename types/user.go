package types

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Gender   string    `json:"gender"`
	Age      int       `json:"age"`
	Avatar   string    `json:"avatar"`
}

func UserHasSession(userID uuid.UUID) (string, bool) {
	for token, session := range Sessions {
		if session.User.ID == userID {
			return token, true
		}
	}
	return "", false
}
