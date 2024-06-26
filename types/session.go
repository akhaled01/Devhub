package types

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"

	"RTF/utils"
)

type Session struct {
	SessionID     uuid.UUID
	User          *User
	Expiry        time.Time
	Conn          *websocket.Conn
	ChatPartnerID string
}

// Map sessionID to session
var (
	Sessions = make(map[uuid.UUID]*Session, 0)
)

// checks if a session is expired or not
func (s Session) IsExpired() bool {
	return !time.Now().Before(s.Expiry)
}

// gets valid session based on id
func ValidateSession(session_id uuid.UUID) (*Session, error) {
	s, ok := Sessions[session_id]
	if !ok {
		return &Session{}, errors.New("invalid session")
	}

	if (Session{}) == *s || s.IsExpired() {
		return &Session{}, errors.New("invalid session")
	}

	return s, nil
}

// Returns the user ID of the user in the current session
func (s Session) GetUserID() uuid.UUID {
	return s.User.ID
}

func (s *Session) CheckExpired() {
	for !s.IsExpired() {
	}
	utils.WarnConsoleLog("session token has expired", "username", s.User.Username)
	delete(Sessions, s.SessionID)
}

// Generates a new session that expires in
// 3600 seconds (one hour)
func GenSession(u User) *Session {
	session_id, err := uuid.NewV7()
	if err != nil {
		utils.ErrorConsoleLog("error generating session -> %s", err)
	}

	Sessions[session_id] = &Session{
		SessionID: session_id,
		User:      &u,
		Expiry:    time.Now().Add(time.Second * 3600),
	}

	return &Session{
		SessionID: session_id,
		User:      &u,
		Expiry:    time.Now().Add(time.Second * 3600),
	}
}

// expires current user session
func LogOutBySessionToken(w http.ResponseWriter, sessionToken uuid.UUID) {
	// Get the session from the Sessions map
	if _, ok := Sessions[sessionToken]; ok {
		delete(Sessions, sessionToken)
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		}
		http.SetCookie(w, cookie)
		w.Header().Add("Set-Cookie", "session_token=; Max-Age=0; HttpOnly")
	}
}

// checks if a user has a current session
func UserHasSessions(user_id uuid.UUID) (*Session, bool) {
	for _, s := range Sessions {
		if s.User.ID == user_id {
			return s, true
		}
	}
	return (&Session{}), false
}

func UserHasSessionByName(username string) (*Session, bool) {
	for _, s := range Sessions {
		if s.User.Username == username {
			return s, true
		}
	}
	return (&Session{}), false
}
