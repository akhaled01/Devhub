package posts

import (
	"encoding/json"
	"errors"
	"net/http"

	"RTF/storage/interfaces/post"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	requester_session, ok := types.Sessions[uuid.FromStringOrNil(session_id.Value)]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// parse id from URL and get the post from it
	post, err := post.GetPostByID(requester_session.User, uuid.FromStringOrNil(r.PathValue("id")))
	if err != nil {
		if errors.Is(err, types.ErrPostNotFound) {
			utils.WarnConsoleLog("user tried to request a non-existing post")
			w.WriteHeader(http.StatusNotFound)
		} else {
			utils.ErrorConsoleLog("error getting post")
			utils.PrintErrorTrace(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		utils.ErrorConsoleLog("error encoding post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
