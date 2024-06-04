package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"RTF/storage/interfaces/comment"
	"RTF/storage/interfaces/user"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for a user to login to the DevHub.
It creates the session. credential can be username / email

# The Request URI for this handler

	`POST /auth/login`

# JSON Example

	{
		credential: "akhaled",
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
	req := types.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorConsoleLog("error decoding json -> %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authenticated_user, err := user.Authenticate(req.Credential, req.Password)
	// check for authentication twice
	if err != nil {
		// check if user not found or incorrect password
		// if neither, its an 500 server error
		if errors.Is(err, types.ErrUserNotFound) {
			utils.WarnConsoleLog("user Not found")
			w.WriteHeader(http.StatusNotFound)
			return
		} else if errors.Is(err, types.ErrIncorrectPassword) {
			utils.WarnConsoleLog("unauthorized -> Incorrect Password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			utils.ErrorConsoleLog("error authenticating user")
			utils.PrintErrorTrace(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if authenticated_user == (types.User{}) {
		utils.WarnConsoleLog("unauthorized Access attempted")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	encoded_avatar, err := utils.EncodeImage(authenticated_user.Avatar)
	if err != nil {
		utils.ErrorConsoleLog("error getting user's avatar")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check if a session with this user exists. logout if thats the case
	if s, has_sessions := types.UserHasSessions(authenticated_user.ID); has_sessions {
		types.LogOutBySessionToken(w, s.SessionID)
	}

	session := types.GenSession(authenticated_user)

	UserCounts, err := comment.GetUserCounts(authenticated_user.ID)
	if err != nil {
		utils.ErrorConsoleLog("error getting user's staristics")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Path:    "/",
		Value:   session.SessionID.String(),
		Expires: session.Expiry,
	})

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct {
		Session_id               string `json:"session_id"`
		Username                 string `json:"username"`
		Email                    string `json:"email"`
		Avatar                   string `json:"encoded_avatar"`
		Gender                   string `json:"gender"`
		Number_of_liked_comments int    `json:"Number_of_liked_comments"`
		Number_of_liked_posts    int    `json:"Number_of_liked_posts"`
		Number_of_comments       int    `json:"Number_of_comments"`
		Number_of_posts          int    `json:"Number_of_posts"`
	}{
		Session_id:               session.SessionID.String(),
		Username:                 authenticated_user.Username,
		Email:                    authenticated_user.Email,
		Avatar:                   encoded_avatar,
		Gender:                   authenticated_user.Gender,
		Number_of_liked_comments: UserCounts.Number_of_liked_comments,
		Number_of_liked_posts:    UserCounts.Number_of_liked_posts,
		Number_of_comments:       UserCounts.Number_of_comments,
		Number_of_posts:          UserCounts.Number_of_posts,
	}); err != nil {
		utils.ErrorConsoleLog("error encoding json")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go session.CheckExpired()
}
