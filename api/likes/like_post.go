package likes

import (
	"RTF/storage/interfaces/likes"
	"RTF/types"
	"RTF/utils"
	"net/http"

	"github.com/gofrs/uuid"
)

func LikePost(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user_id := types.Sessions[uuid.FromStringOrNil(session_id.Value)].GetUserID()
	post_id := r.PathValue("id")
	if post_id == "" {
		utils.ErrorConsoleLog("post id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	like_bool, err := likes.CheckUserPostLike(uuid.FromStringOrNil(post_id), user_id)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !like_bool {
		err = likes.CreateLikeRecordPost(uuid.FromStringOrNil(post_id), user_id)
		if err != nil {
			utils.ErrorConsoleLog(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err = likes.DeleteLikeRecordPost(uuid.FromStringOrNil(post_id), user_id)
		if err != nil {
			utils.ErrorConsoleLog(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
