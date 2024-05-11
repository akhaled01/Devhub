package likes

import (
	"RTF/storage/interfaces/likes"
	"RTF/storage/interfaces/comment"

	"RTF/types"
	"RTF/utils"
	"net/http"

	"github.com/gofrs/uuid"
	"encoding/json" // Add this line to import the "encoding/json" package
)

func LikeComment(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user_id := types.Sessions[uuid.FromStringOrNil(session_id.Value)].GetUserID()
	comment_id := r.PathValue("id")
	if comment_id == "" {
		utils.ErrorConsoleLog("post id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	like_bool, err := likes.CheckUserCommentLike(uuid.FromStringOrNil(comment_id), user_id)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var updatedComment *types.Comment
	if !like_bool {
		err = likes.CreateLikeRecordComment(uuid.FromStringOrNil(comment_id), user_id)
		if err != nil {
			utils.ErrorConsoleLog(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		updatedComment, err = comment.GetCommentByID(uuid.FromStringOrNil(comment_id))
        if err != nil {
            utils.ErrorConsoleLog(err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
	} else {
		err = likes.DeleteLikeRecordComment(uuid.FromStringOrNil(comment_id), user_id)
		if err != nil {
			utils.ErrorConsoleLog(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		updatedComment, err = comment.GetCommentByID(uuid.FromStringOrNil(comment_id))
        if err != nil {
            utils.ErrorConsoleLog(err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedComment)
}
