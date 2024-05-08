package auth

import (
	"net/http"

	"RTF/api/chat"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

/*
 Recieves a logout request and deletes user session, and send a channel

	POST /auth/logout
*/
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog("error extracting cookie")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	delete(types.Sessions, uuid.FromStringOrNil(cookie.Value))
	chat.ListenerChan <- true
}
