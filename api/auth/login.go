package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"RTF/log"
	"RTF/storage/interfaces/user"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for a user to login to the DevHub.
It creates the session

# The Request URI for this handler

	`POST /auth/login`

# JSON Examples

	{
		credential: "akhaled",
		password: "azt@345"
	}

	OR

	{
		credential: "akhaledlarus@gmail.com",
		password: "azt@345"
	}
	
EXAMPLE SUCCESSFUL RESPONSE (200 OK)

	{
		session_id : "xxxxxxxxxxxxxxxxx-xxxxxxxx-xxxxxx",
		username : "akhaled",
		email : "akhaledlarus@gmail.com"
		Avatar : "ENCODED_AVATAR_JIBBERISH"
	}
*/
func Login(w http.ResponseWriter, r *http.Request) {
	req := &types.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.ErrorConsoleLog("error decoding json -> %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authenticated_user, err := user.Authenticate(req.Credential, req.Password)
	// check for authentication twice
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			log.WarnConsoleLog("unauthorized Access attempted")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			log.ErrorConsoleLog("error authenticating user -> %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if authenticated_user == nil {
		log.WarnConsoleLog("unauthorized Access attempted")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session := types.GenSession(*authenticated_user)
	go session.CheckExpired()

	encoded_avatar, err := utils.EncodeImage(authenticated_user.Avatar)
	if err != nil {
		log.ErrorConsoleLog("error getting user's avatar")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(struct {
		Session_id string `json:"session_id"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Avatar     string `json:"encoded_avatar"`
	}{
		Session_id: session.SessionID.String(),
		Username:   authenticated_user.Username,
		Email:      authenticated_user.Email,
		Avatar:     encoded_avatar,
	}); err != nil {
		log.WarnConsoleLog("error encoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
