package types

import (
	"errors"
	"time"

	"RTF/log"

	"github.com/gofrs/uuid"
)

type Session struct {
	SessionID uuid.UUID
	User      User
	Expiry    time.Time
}

var Sessions = make(map[uuid.UUID]Session, 0)

// checks if a session is expired or not
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

// gets valid session based on id
func ValidateSession(session_id uuid.UUID) (*Session, error) {
	s := Sessions[session_id]

	if (Session{}) == s || s.IsExpired() {
		return &Session{}, errors.New("invalid session")
	}

	return &s, nil
}

// Returns the user ID of the user in the current session
func (s Session) GetUserID() uuid.UUID {
	return s.User.ID
}

func (s *Session) CheckExpired() {
	for !s.IsExpired() {
	}
	log.InfoConsoleLog("%s's session token has expired", s.User.Username)
	delete(Sessions, s.SessionID)
}

// Generates a new session that expires in
// 3600 seconds (one hour)
func GenSession(u User) *Session {
	session_id, err := uuid.NewV7()
	if err != nil {
		log.ErrorConsoleLog("error generating session -> %s", err)
	}

	return &Session{
		SessionID: session_id,
		User:      u,
		Expiry:    <-time.After(time.Second * 3600),
	}
}
