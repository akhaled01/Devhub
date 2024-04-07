package types

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	SessionID uuid.UUID
	User      User
	Expiry    time.Time
}

var Sessions = make(map[string]Session, 0)

// checks if a session is expired or not
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

// Returns the user ID of the user in the current session
func (s Session) GetUserID() uuid.UUID {
	return s.User.ID
}

func (s *Session) CheckExpired() {
	for !s.IsExpired() {
	}
	fmt.Printf("User %s token expired!\n", s.User.Username)
}
