package posts

import (
	"encoding/json"
	"errors"
	"net/http"

	"RTF/storage/interfaces/comment"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	// parse id from URL and get the post from it
	comments, err := comment.GetPostCommentsByID(uuid.FromStringOrNil(r.PathValue("id"))) // this is post ID
	if err != nil {
		if errors.Is(err, types.ErrPostNotFound) {
			utils.WarnConsoleLog("user tried to request a non-existing comments")
			w.WriteHeader(http.StatusNotFound)
		} else {
			utils.ErrorConsoleLog("error getting comments")
			utils.PrintErrorTrace(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		utils.ErrorConsoleLog("error encoding post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
