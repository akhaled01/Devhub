package utils

import (
	"net/http"

	"RTF/types"
)

// Validates Session Based on cookie from client
func GetValidSession(w http.ResponseWriter, r *http.Request) (*types.Session, bool) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return nil, false
		}

		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	session_token := cookie.Value
	userSession, exists := types.Sessions[session_token]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, false
	}

	if userSession.IsExpired() {
		delete(types.Sessions, session_token)
		w.WriteHeader(http.StatusUnauthorized)
		return nil, false
	}

	return &userSession, true
}
