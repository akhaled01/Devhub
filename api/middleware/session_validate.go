package middleware

import (
	"errors"
	"net/http"

	"github.com/gofrs/uuid"

	"RTF/types"
	"RTF/utils"
)

// A middleware that validates sessions
// for any handlers that require any sort of session validation
func SessionValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.ErrorConsoleLog("ErrNoCookie found")
				w.WriteHeader(http.StatusUnauthorized)
				//http.Redirect(w, r, "/signup", http.StatusPermanentRedirect)
				return
			}
			utils.ErrorConsoleLog("error extracting session_id cookie")
			utils.PrintErrorTrace(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		is_session_valid, err := types.ValidateSession(uuid.FromStringOrNil(cookie.Value))
		if err != nil {
			utils.ErrorConsoleLog("error extracting session_id cookie")
			utils.PrintErrorTrace(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if is_session_valid == (&types.Session{}) {
			utils.WarnConsoleLog("Authentication expired for %s", is_session_valid.User.Username)
			w.WriteHeader(http.StatusUnauthorized)
			//http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
