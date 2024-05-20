package likes

import (
	"RTF/storage"
	"RTF/types"
	"RTF/utils"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
)

const query_all_user_stats = `SELECT
(SELECT COUNT(*) FROM comment_likes WHERE user_id =  ?) AS number_of_liked_comments,
(SELECT COUNT(*) FROM post_likes WHERE user_id =  ?) AS number_of_liked_posts,
(SELECT COUNT(*) FROM posts WHERE user_id =  ?) AS number_of_posts,
(SELECT COUNT(*) FROM comments WHERE user_id =  ?) AS number_of_comments LIMIT 100`

func GetUserCounts(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	user_id := types.Sessions[uuid.FromStringOrNil(session_id.Value)].GetUserID()
	user := types.Sessions[uuid.FromStringOrNil(session_id.Value)].User

	UserCounts := types.Counts{}
	stmt, err := storage.DB_Conn.Prepare(query_all_user_stats)
	if err != nil {
		utils.ErrorConsoleLog("post id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	if err := stmt.QueryRow(user_id, user_id, user_id, user_id).Scan(&UserCounts.Number_of_liked_comments, &UserCounts.Number_of_liked_posts, &UserCounts.Number_of_posts, &UserCounts.Number_of_comments); err != nil {
		utils.ErrorConsoleLog("post id not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting user's avatar
	encoded_avatar, err := utils.EncodeImage(user.Avatar)
	if err != nil {
		utils.ErrorConsoleLog("error getting user's avatar")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Filling user details
	UserCounts.SessionID = uuid.FromStringOrNil(session_id.Value)
	UserCounts.Username = user.Username
	UserCounts.Email = user.Email
	UserCounts.Avatar = encoded_avatar
	UserCounts.Gender = user.Gender

	// Returning user counts
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(UserCounts); err != nil {
		utils.ErrorConsoleLog("error encoding all posts")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
